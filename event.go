package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"hearts/data"
)

type inboundEvent interface {
	playerId() uuid.UUID
}

type connectPlayerInboundEvent player

func (c connectPlayerInboundEvent) playerId() uuid.UUID { return c.id }

type playCardInboundEvent struct {
	playedBy uuid.UUID
	card     data.Card
}

func (p playCardInboundEvent) playerId() uuid.UUID { return p.playedBy }

func handleWebsocketMessages(p player, g game) {
	for {
		var wm data.WebsocketMessage
		err := p.ReadJSON(&wm)
		if err != nil {
			p.writeClientViolation(err.Error())
			continue
		}

		switch data.InboundEventType(wm.Type) {
		case data.InboundEventPlayCard:
			cardMap, hasCard := wm.Data["card"]
			if !hasCard {
				p.writeClientViolation("Card object is missing")
				continue
			}

			cardJson, err := json.Marshal(cardMap)
			if err != nil {
				p.writeClientViolation(err.Error())
				continue
			}

			var card data.Card
			logNonFatal(json.Unmarshal(cardJson, &card))

			g.inboundEvents <- playCardInboundEvent{p.id, card}
		}
	}
}
