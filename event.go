package main

import (
	"encoding/json"
)

var (
	inboundEventsByName = map[string]func(player) inboundEvent{
		"playCard": func(p player) inboundEvent { return &playCardInboundEvent{playedBy: p} },
	}
)

type inboundEvent interface {
	player() player
}

type connectPlayerInboundEvent player

func (c connectPlayerInboundEvent) player() player { return player(c) }

type playCardInboundEvent struct {
	playedBy player
	Card     `json:"card"`
}

func (p playCardInboundEvent) player() player { return p.playedBy }

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

		e := inboundEventsByName[we.Type](p)
		err = json.Unmarshal(data, e)
		if err != nil {
			p.writeCloseMessageError(err)
			return
		}

		g.inboundEvents <- e
	}
}
