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
	players         []player
	inPlay          []Card
	currentPlayerId uuid.UUID
}

func (g game) run() {
	for x := range g.inboundEvents {
		switch e := x.(type) {
		case *playCardInboundEvent:
			if e.player().id != g.currentPlayerId {
				e.player().writeClientViolation("cannot play out of turn")
				continue
			}

			if !validCardPlayed(g.inPlay, *e) {
				e.player().writeClientViolation("cannot play off suit")
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
				e.player().writeCloseMessageError(errors.New("game is full"))
				continue
			}

			g.players = append(g.players, player(e))

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
			e.player().writeCloseMessageError(errors.New("inboundEvent handler not found"))
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

func validCardPlayed(inPlay []Card, event playCardInboundEvent) bool {
	if len(inPlay) == 0 {
		return true // It can be anything your heart desires.
	}

	suitMatches := true
	for _, c := range inPlay {
		if c.Suit != event.Card.Suit {
			suitMatches = false
			break
		}
	}

	var playerHasMatchingSuit, playerHasCard bool
	for _, c := range event.player().hand {
		if c.Suit == inPlay[0].Suit {
			playerHasMatchingSuit = true
		}
		if reflect.DeepEqual(c, event.Card) {
			playerHasCard = true
		}
	}

	return playerHasCard && (suitMatches || !playerHasMatchingSuit)
}
