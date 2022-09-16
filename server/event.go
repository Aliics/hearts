package server

import (
	"github.com/aliics/hearts/data"
	"github.com/google/uuid"
)

type playerGameEvent struct {
	playerID uuid.UUID
	event    data.InboundEvent
}

type connectionEventType uint8

const (
	connectionEventConnect connectionEventType = iota
	connectionEventDisconnect
)

type connectionEvent struct {
	eventType connectionEventType
	player    *player
}

func handleWebsocketMessages(p *player, g game) {
	for !p.isClosed {
		var ip data.InboundPayload
		err := p.ReadJSON(&ip)
		if err != nil {
			p.writeClientViolation(err.Error())
			continue
		}

		g.playerGameEventsCh <- playerGameEvent{p.id, ip.Data}
	}
	g.connectionEventsCh <- connectionEvent{connectionEventDisconnect, p}
}
