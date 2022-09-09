package main

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
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

			suitMatches := true
			for _, card := range g.inPlay {
				if card.Suit != e.Card.Suit {
					suitMatches = false
					break
				}
			}

			var playerHasSuit bool
			for _, card := range e.player().hand {
				if card.Suit == e.Card.Suit {
					playerHasSuit = true
					break
				}
			}

			if !suitMatches && playerHasSuit {
				e.player().writeOutboundEventMessage(outboundEventClientViolation, "cannot play off suit")
				continue
			}

			g.inPlay = append(g.inPlay, e.Card)

			for _, p := range g.players {
				inPlay, err := json.Marshal(g.inPlay)
				logNonFatal(err) // TODO: Handle more gracefully?
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
		default:
			e.player().writeCloseMessageError(errors.New("inboundEvent handler not found"))
		}
	}
}

func (g game) connectPlayer(p player) {
	g.inboundEvents <- connectPlayerInboundEvent(p)
}

type player struct {
	*websocket.Conn
	id   uuid.UUID
	hand []Card
}

func (p player) writeOutboundEventMessage(eventType outboundEventType, msg string) {
	p.writeOutboundEvent(eventType, map[string]any{"msg": msg})
}

func (p player) writeOutboundEvent(eventType outboundEventType, data map[string]any) {
	we, err := json.Marshal(websocketEvent{
		Type: string(eventType),
		Data: data,
	})
	logNonFatal(err)
	logNonFatal(p.WriteMessage(websocket.TextMessage, we))
}

func (p player) writeCloseMessageError(err error) {
	logNonFatal(p.WriteControl(websocket.CloseMessage, []byte(err.Error()), time.Now()))
}
