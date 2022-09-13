package server

import (
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
	games = make(map[uuid.UUID]game)
)

type playerCard struct {
	player
	data.Card
}

type game struct {
	inboundEvents   chan inboundEvent
	players         []*player
	inPlay          []playerCard
	currentPlayerID uuid.UUID
}

func (g game) run() {
	for x := range g.inboundEvents {
		switch e := x.(type) {
		case playCardInboundEvent:
			p := g.playerByID(x.playerID())

			if len(g.players) < fullGameCount {
				p.writeClientViolation("cannot play without a full game")
				continue
			}

			if e.playerID() != g.currentPlayerID {
				p.writeClientViolation("cannot play out of turn")
				continue
			}

			cardIndex := indexOfValidPlayedCard(g.inPlay, p.hand, e.card)
			if cardIndex == -1 {
				p.writeClientViolation("invalid Card")
				continue
			}

			// Current cards in play for the round.
			g.inPlay = append(g.inPlay, playerCard{*p, e.card})
			g.broadcastUpdate(map[string]any{"inPlay": g.inPlay})

			// Player loses Card played from their hand.
			p.hand = slices.Delete(p.hand, cardIndex, cardIndex+1)
			p.writeOutboundEvent(data.OutboundEventGameUpdate, map[string]any{"hand": p.hand})

			if len(g.inPlay) < fullGameCount {
				// Simply rotate players.
				currentPlayerIndex := g.indexOfPlayerById(g.currentPlayerID)
				if currentPlayerIndex < fullGameCount-1 {
					g.currentPlayerID = g.players[currentPlayerIndex+1].id
				} else {
					g.currentPlayerID = g.players[0].id
				}
				g.broadcastUpdate(map[string]any{"currentPlayerID": g.currentPlayerID})
			} else {
				highestPlayerCard, addPoints := getHighestInPlay(g.inPlay)

				// Apply points to the winner and make them the next to play.
				g.currentPlayerID = highestPlayerCard.id
				g.playerByID(g.currentPlayerID).points += addPoints

				points := make(map[uuid.UUID]int)
				for _, p := range g.players {
					points[p.id] = p.points
				}
				g.broadcastUpdate(map[string]any{
					"currentPlayerID": g.currentPlayerID,
					"points":          points,
				})

				g.inPlay = nil // Reset the game for the next round.
			}
		case connectPlayerInboundEvent:
			if len(g.players) == fullGameCount {
				// p.writeCloseMessageError(errors.New("game is full"))
				continue
			}

			newPlayer := player(e)
			g.players = append(g.players, &newPlayer)

			g.broadcastUpdate(map[string]any{
				"playerCount": len(g.players),
				"gameReady":   len(g.players) == fullGameCount,
			})

			// Deal some cards. :)
			if len(g.players) == fullGameCount {
				deck := newShuffledDeck()
				handSize := len(deck) / fullGameCount
				for i, p := range g.players {
					p.hand = deck[handSize*i : handSize*(i+1)]
					p.writeOutboundEvent(data.OutboundEventGameUpdate, map[string]any{"hand": p.hand})
				}

				g.currentPlayerID = g.players[0].id
				g.broadcastUpdate(map[string]any{"currentPlayerID": g.currentPlayerID})
			}
		case disconnectPlayerInboundEvent:
			playerIndex := g.indexOfPlayerById(e.playerID())
			g.players = slices.Delete(g.players, playerIndex, playerIndex+1)
			g.broadcastUpdate(map[string]any{
				"playerCount": len(g.players),
				"gameReady":   len(g.players) == fullGameCount,
			})
		default:
			log.Println("unhandled event")
		}
	}
}

func (g game) connectPlayer(p *player) {
	g.inboundEvents <- connectPlayerInboundEvent(*p)
}

func (g game) disconnectPlayer(p *player) {
	g.inboundEvents <- disconnectPlayerInboundEvent(*p)
}

func (g game) broadcastUpdate(msg map[string]any) {
	for _, p := range g.players {
		p.writeOutboundEvent(data.OutboundEventGameUpdate, msg)
	}
}

func (g game) playerByID(id uuid.UUID) *player {
	return g.players[g.indexOfPlayerById(id)]
}

func (g game) indexOfPlayerById(id uuid.UUID) int {
	return slices.IndexFunc[*player](g.players, func(p *player) bool { return p.id == id })
}

// Tally points for the winner.
func getHighestInPlay(inPlay []playerCard) (highest playerCard, points int) {
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
func indexOfValidPlayedCard(inPlay []playerCard, hand []data.Card, played data.Card) int {
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
