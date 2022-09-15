package server

import (
	"github.com/aliics/hearts/data"
	"github.com/google/uuid"
)

type gameEvent struct {
	playerID uuid.UUID
	event    data.InboundEvent
}

func handleWebsocketMessages(p *player, g *game) {
	for !p.isClosed {
		var ip data.InboundPayload
		err := p.ReadJSON(&ip)
		if err != nil {
			p.writeClientViolation(err.Error())
			continue
		}

		g.inboundEvents <- gameEvent{p.id, ip.Data}
	}
	g.disconnectPlayer(p)
}
