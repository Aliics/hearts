package server

import (
	"errors"
	"github.com/aliics/hearts/data"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"log"
	"reflect"
)

const (
	fullGameCount = 3
)

var (
	games = make(map[uuid.UUID]*game)
)

type game struct {
	inboundEvents   chan gameEvent
	players         []*player
	inPlay          []data.PlayerCard
	currentPlayerID uuid.UUID
}

func (g *game) run() {
	for x := range g.inboundEvents {
		switch e := x.event.(type) {
		case data.PlayCardEvent:
			p := g.playerByID(x.playerID)

			if len(g.players) < fullGameCount {
				p.writeClientViolation("cannot play without a full game")
				continue
			}

			if x.playerID != g.currentPlayerID {
				p.writeClientViolation("cannot play out of turn")
				continue
			}

			cardIndex := indexOfValidPlayedCard(g.inPlay, p.hand, e.Card)
			if cardIndex == -1 {
				p.writeClientViolation("invalid Card")
				continue
			}

			// Current cards in play for the round.
			g.inPlay = append(g.inPlay, data.PlayerCard{PlayerID: p.id, Card: e.Card})
			g.broadcastOutboundPayload(data.OutboundPayload{
				Type: data.OutboundInPlay,
				Data: data.InPlayEvent{InPlay: g.inPlay},
			})

			// Player loses Card played from their hand.
			p.hand = slices.Delete(p.hand, cardIndex, cardIndex+1)
			p.writeOutboundPayload(data.OutboundPayload{
				Type: data.OutboundNewHand,
				Data: data.NewHandEvent{Hand: p.hand},
			})

			if len(g.inPlay) < fullGameCount {
				// Simply rotate players.
				currentPlayerIndex := g.indexOfPlayerById(g.currentPlayerID)
				if currentPlayerIndex < fullGameCount-1 {
					g.currentPlayerID = g.players[currentPlayerIndex+1].id
				} else {
					g.currentPlayerID = g.players[0].id
				}
			} else {
				highestPlayerCard, addPoints := getHighestInPlay(g.inPlay)

				// Apply points to the winner and make them the next to play.
				g.currentPlayerID = highestPlayerCard.PlayerID
				g.playerByID(g.currentPlayerID).points += addPoints

				points := make(map[uuid.UUID]int)
				for _, p := range g.players {
					points[p.id] = p.points
				}
				g.broadcastOutboundPayload(data.OutboundPayload{
					Type: data.OutboundPoints,
					Data: data.PointsEvent{Points: points},
				})

				g.inPlay = nil // Reset the game for the next round.
			}
			g.broadcastOutboundPayload(data.OutboundPayload{
				Type: data.OutboundCurrentPlayer,
				Data: data.CurrentPlayerEvent{PlayerID: g.currentPlayerID},
			})
		default:
			log.Println("unhandled event")
		}
	}
}

func (g *game) connectPlayer(p *player) {
	if len(g.players) == fullGameCount {
		p.WriteCloseMessageError(errors.New("game is full"))
		return
	}

	g.players = append(g.players, p)

	g.broadcastCurrentPlayers()

	// Deal some cards. :)
	if len(g.players) == fullGameCount {
		deck := newShuffledDeck()
		handSize := len(deck) / fullGameCount
		for i, p := range g.players {
			p.hand = deck[handSize*i : handSize*(i+1)]
			p.writeOutboundPayload(data.OutboundPayload{
				Type: data.OutboundNewHand,
				Data: data.NewHandEvent{Hand: p.hand},
			})
		}

		g.currentPlayerID = g.players[0].id
		g.broadcastOutboundPayload(data.OutboundPayload{
			Type: data.OutboundCurrentPlayer,
			Data: data.CurrentPlayerEvent{PlayerID: g.currentPlayerID},
		})
	}
}

func (g *game) disconnectPlayer(p *player) {
	playerIndex := g.indexOfPlayerById(p.id)
	g.players = slices.Delete(g.players, playerIndex, playerIndex+1)
	g.broadcastCurrentPlayers()
}

func (g *game) broadcastCurrentPlayers() {
	var playerIDs []uuid.UUID
	for _, p := range g.players {
		playerIDs = append(playerIDs, p.id)
	}
	g.broadcastOutboundPayload(data.OutboundPayload{
		Type: data.OutboundCurrentPlayers,
		Data: data.CurrentPlayersEvent{
			PlayerIDs: playerIDs,
			GameReady: len(g.players) == fullGameCount,
		},
	})
}

func (g *game) broadcastOutboundPayload(payload data.OutboundPayload) {
	for _, p := range g.players {
		p.writeOutboundPayload(payload)
	}
}

func (g *game) playerByID(id uuid.UUID) *player {
	return g.players[g.indexOfPlayerById(id)]
}

func (g *game) indexOfPlayerById(id uuid.UUID) int {
	return slices.IndexFunc[*player](g.players, func(p *player) bool { return p.id == id })
}

// Tally points for the winner.
func getHighestInPlay(inPlay []data.PlayerCard) (highest data.PlayerCard, points int) {
	highest = inPlay[0]
	for _, c := range inPlay {
		points += c.Worth()
		if c.Beats(highest.Card) {
			highest = c
		}
	}

	return
}

// indexOfValidPlayedCard finds the index of the played Card
// with the following conditions:
//
//		The played Card must be in "hand" and...
//		   1. The played Card must be in "inPlay"
//	    or
//		   2. The played Card's suit must not be in "inPlay"
func indexOfValidPlayedCard(inPlay []data.PlayerCard, hand []data.Card, played data.Card) int {
	cardIndex := -1
	inPlaySuitInHand := false
	for i, c := range hand {
		if reflect.DeepEqual(c, played) {
			cardIndex = i
		}
		if len(inPlay) > 0 && c.Suit == inPlay[0].Suit {
			inPlaySuitInHand = true
		}
	}

	if len(inPlay) == 0 || played.Suit == inPlay[0].Suit || !inPlaySuitInHand {
		return cardIndex
	} else {
		return -1
	}
}
