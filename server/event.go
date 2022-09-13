package server

import (
	"encoding/json"
	"github.com/aliics/hearts/data"
	"github.com/google/uuid"
)

type inboundEvent interface {
	playerID() uuid.UUID
}

type connectPlayerInboundEvent player

func (c connectPlayerInboundEvent) playerID() uuid.UUID { return c.id }

type disconnectPlayerInboundEvent player

func (d disconnectPlayerInboundEvent) playerID() uuid.UUID { return d.id }

type playCardInboundEvent struct {
	playedBy uuid.UUID
	card     data.Card
}

func (p playCardInboundEvent) playerID() uuid.UUID { return p.playedBy }

func handleWebsocketMessages(p *player, g game) {
	for !p.isClosed {
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

			cardJSON, err := json.Marshal(cardMap)
			if err != nil {
				p.writeClientViolation(err.Error())
				continue
			}

			var card data.Card
			logNonFatal(json.Unmarshal(cardJSON, &card))

			g.inboundEvents <- playCardInboundEvent{p.id, card}
		}
	}
	g.disconnectPlayer(p)
}
