package main

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

var (
	games = make(map[uuid.UUID]game)
)

type game struct {
	inboundEvents chan inboundEvent
	players       []player
	inPlay        []Card
	currentPlayer player
}

func (g game) run() {
	for x := range g.inboundEvents {
		switch e := x.(type) {
		case *playCardInboundEvent:
			if e.player().id != g.currentPlayer.id {
				e.player().writeOutboundEventMessage(outboundEventClientViolation, "cannot play out of turn")
				return
			}

			if validCardPlayed(g.inPlay, *e) {
				e.player().writeOutboundEventMessage(outboundEventClientViolation, "cannot play off suit")
				continue
			}

			g.inPlay = append(g.inPlay, e.Card)

			for _, p := range g.players {
				inPlay, err := json.Marshal(g.inPlay)
				logNonFatal(err)
				p.writeOutboundEvent(outboundEventGameUpdate, map[string]any{"inPlay": inPlay})
			}
		case connectPlayerInboundEvent:
			if len(g.players) == 3 {
				e.player().writeCloseMessageError(errors.New("game is full"))
				continue
			}
			g.players = append(g.players, player(e))

			for _, p := range g.players {
				p.writeOutboundEvent(outboundEventGameUpdate, map[string]any{
					"playerCount": len(g.players),
					"gameReady":   len(g.players) == 3,
				})
			}

			if len(g.players) == 3 {
				deck := newShuffledDeck()
				handSize := len(deck) / 3
				for i, p := range g.players {
					p.hand = deck[handSize*i : handSize*(i+1)]
					p.writeOutboundEvent(outboundEventGameUpdate, map[string]any{"hand": p.hand})
				}
			}
		default:
			e.player().writeCloseMessageError(errors.New("inboundEvent handler not found"))
		}
	}
}

func (g game) connectPlayer(p player) {
	g.inboundEvents <- connectPlayerInboundEvent(p)
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

	var playerHasSuit bool
	for _, c := range event.player().hand {
		if c.Suit == inPlay[0].Suit {
			playerHasSuit = true
			break
		}
	}

	return suitMatches || !playerHasSuit
}
