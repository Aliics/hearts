package main

import (
	"encoding/json"
	"github.com/google/uuid"
)

var (
	inboundEventsByName = map[string]func(uuid.UUID) inboundEvent{
		"playCard": func(id uuid.UUID) inboundEvent { return &playCardInboundEvent{playedBy: id} },
	}
)

type inboundEvent interface {
	playerId() uuid.UUID
}

type connectPlayerInboundEvent player

func (c connectPlayerInboundEvent) playerId() uuid.UUID { return c.id }

type playCardInboundEvent struct {
	playedBy uuid.UUID
	Card     `json:"card"`
}

func (p playCardInboundEvent) playerId() uuid.UUID { return p.playedBy }

type websocketEvent struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type outboundEventType string

const (
	outboundEventClientViolation outboundEventType = "clientViolation"
	outboundEventGameUpdate      outboundEventType = "gameUpdate"
)

func handleIncomingEvents(p player, g game) {
	for {
		var we websocketEvent
		err := p.ReadJSON(&we)
		if err != nil {
			p.writeCloseMessageError(err)
			return
		}

		data, err := json.Marshal(we.Data)
		if err != nil {
			p.writeCloseMessageError(err)
			return
		}

		e := inboundEventsByName[we.Type](p.id)
		err = json.Unmarshal(data, e)
		if err != nil {
			p.writeCloseMessageError(err)
			return
		}

		g.inboundEvents <- e
	}
}
