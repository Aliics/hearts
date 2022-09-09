package main

import (
	"errors"
	"github.com/google/uuid"
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
		p := g.playerFromId(x.playerId())

		switch e := x.(type) {
		case *playCardInboundEvent:
			if len(g.players) < fullGameCount {
				p.writeClientViolation("cannot play without a full game")
				continue
			}

			if e.playerId() != g.currentPlayerId {
				p.writeClientViolation("cannot play out of turn")
				continue
			}

			if !validCardPlayed(g.inPlay, p.hand, e.Card) {
				p.writeClientViolation("invalid card")
				continue
			}

			g.inPlay = append(g.inPlay, e.Card)

			g.broadcastUpdate(map[string]any{"inPlay": g.inPlay})

			var currentIndex int
			for i, p := range g.players {
				if p.id == g.currentPlayerId {
					currentIndex = i
					break
				}
			}
			if currentIndex < fullGameCount-1 {
				g.currentPlayerId = g.players[currentIndex+1].id
			} else {
				g.currentPlayerId = g.players[0].id
			}

			g.broadcastUpdate(map[string]any{"currentPlayerNum": g.currentPlayerId})
		case connectPlayerInboundEvent:
			if len(g.players) == fullGameCount {
				p.writeCloseMessageError(errors.New("game is full"))
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
				g.broadcastUpdate(map[string]any{"currentPlayerNum": g.currentPlayerId})
			}
		default:
			p.writeCloseMessageError(errors.New("inboundEvent handler not found"))
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

func (g game) playerFromId(id uuid.UUID) *player {
	var found *player
	for _, p := range g.players {
		if p.id == id {
			found = p
		}
	}
	return found
}

func validCardPlayed(inPlay []Card, hand []Card, played Card) bool {
	suitMatches := true
	for _, c := range inPlay {
		if c.Suit != played.Suit {
			suitMatches = false
			break
		}
	}

	var playerHasMatchingSuit, playerHasCard bool
	for _, c := range hand {
		if len(inPlay) > 0 && c.Suit == inPlay[0].Suit {
			playerHasMatchingSuit = true
		}
		if reflect.DeepEqual(c, played) {
			playerHasCard = true
		}
	}

	return playerHasCard && (suitMatches || !playerHasMatchingSuit)
}
