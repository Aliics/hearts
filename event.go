package main

import (
	"encoding/json"
	"github.com/google/uuid"
)

type websocketMessage struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type inboundEventType string

const (
	inboundEventPlayCard inboundEventType = "playCard"
)

type outboundEventType string

const (
	outboundEventClientViolation outboundEventType = "clientViolation"
	outboundEventGameUpdate      outboundEventType = "gameUpdate"
)

type inboundEvent interface {
	playerId() uuid.UUID
}

type connectPlayerInboundEvent player

func (c connectPlayerInboundEvent) playerId() uuid.UUID { return c.id }

type playCardInboundEvent struct {
	playedBy uuid.UUID
	card     card
}

func (p playCardInboundEvent) playerId() uuid.UUID { return p.playedBy }

func handleWebsocketMessages(p player, g game) {
	for {
		var wm websocketMessage
		err := p.ReadJSON(&wm)
		if err != nil {
			p.writeClientViolation(err.Error())
			continue
		}

		switch inboundEventType(wm.Type) {
		case inboundEventPlayCard:
			cardMap, hasCard := wm.Data["card"]
			if !hasCard {
				p.writeClientViolation("card object is missing")
				continue
			}

			cardJson, err := json.Marshal(cardMap)
			if err != nil {
				p.writeClientViolation(err.Error())
				continue
			}

			var card card
			logNonFatal(json.Unmarshal(cardJson, &card))

			g.inboundEvents <- playCardInboundEvent{p.id, card}
		}
	}
}
