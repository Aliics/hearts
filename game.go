package main

import (
	"github.com/google/uuid"
	"log"
	"reflect"
)

const (
	fullGameCount = 3
)

var (
	games = make(map[uuid.UUID]game)
)

type game struct {
	inboundEvents   chan inboundEvent
	players         []*player
	inPlay          []Card
	currentPlayerId uuid.UUID
}

func (g game) run() {
	for x := range g.inboundEvents {
		switch e := x.(type) {
		case playCardInboundEvent:
			p := g.players[g.indexOfPlayerById(x.playerId())]

			if len(g.players) < fullGameCount {
				p.writeClientViolation("cannot play without a full game")
				continue
			}

			if e.playerId() != g.currentPlayerId {
				p.writeClientViolation("cannot play out of turn")
				continue
			}

			cardIndex := indexOfValidPlayedCard(g.inPlay, p.hand, e.card)
			if cardIndex == -1 {
				p.writeClientViolation("invalid card")
				continue
			}

			// Current cards in play for the round.
			g.inPlay = append(g.inPlay, e.card)
			g.broadcastUpdate(map[string]any{"inPlay": g.inPlay})

			// Player loses card played from their hand.
			p.hand = append(p.hand[:cardIndex], p.hand[cardIndex+1:]...)
			p.writeOutboundEvent(outboundEventGameUpdate, map[string]any{"hand": p.hand})

			if len(g.inPlay) < fullGameCount {
				// Simply rotate players.
				currentPlayerIndex := g.indexOfPlayerById(g.currentPlayerId)
				if currentPlayerIndex < fullGameCount-1 {
					g.currentPlayerId = g.players[currentPlayerIndex+1].id
				} else {
					g.currentPlayerId = g.players[0].id
				}
				g.broadcastUpdate(map[string]any{"currentPlayerId": g.currentPlayerId})
			} else {
				// Tally points for the "winner".
				var (
					pointsCount      int
					highestCardIndex int
					highestCard      = g.inPlay[0]
				)
				for i, c := range g.inPlay {
					pointsCount += c.points()
					if c.beats(highestCard) {
						highestCardIndex = i
						highestCard = c
					}
				}

				g.inPlay = nil
				winner := g.players[highestCardIndex]
				winner.points += pointsCount

				g.currentPlayerId = winner.id
				points := make(map[uuid.UUID]int)
				for _, p := range g.players {
					points[p.id] = p.points
				}
				g.broadcastUpdate(map[string]any{
					"currentPlayerId": g.currentPlayerId,
					"points":          points,
				})
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

			if len(g.players) == fullGameCount {
				deck := newShuffledDeck()
				handSize := len(deck) / fullGameCount
				for i, p := range g.players {
					p.hand = deck[handSize*i : handSize*(i+1)]
					p.writeOutboundEvent(outboundEventGameUpdate, map[string]any{"hand": p.hand})
				}

				g.currentPlayerId = g.players[0].id
				g.broadcastUpdate(map[string]any{"currentPlayerId": g.currentPlayerId})
			}
		default:
			log.Println("unhandled event")
		}
	}
}

func (g game) connectPlayer(p player) {
	g.inboundEvents <- connectPlayerInboundEvent(p)
}

func (g game) broadcastUpdate(msg map[string]any) {
	for _, p := range g.players {
		p.writeOutboundEvent(outboundEventGameUpdate, msg)
	}
}

func (g game) indexOfPlayerById(id uuid.UUID) int {
	for i, p := range g.players {
		if p.id == id {
			return i
		}
	}
	return -1
}

// indexOfValidPlayedCard finds the index of the played card
// with the following conditions:
//
//		The played card must be in "hand" and...
//		   1. The played card must be in "inPlay"
//	    or
//		   2. The played card's suit must not be in "inPlay"
func indexOfValidPlayedCard(inPlay []Card, hand []Card, played Card) int {
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
