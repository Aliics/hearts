package main

import (
	"encoding/json"
)

var (
	eventsByName = map[string]func(player) event{
		"playCard": func(p player) event { return &playCardEvent{playedBy: p} },
	}
)

type event interface {
	player() player
}

type connectPlayerEvent player

func (c connectPlayerEvent) player() player { return player(c) }

type playCardEvent struct {
	playedBy player
	Card     `json:"card"`
}

func (p playCardEvent) player() player { return p.playedBy }

type websocketEvent struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

func handleIncomingEvents(p player, g game) {
	for {
		var we websocketEvent
		err := p.ReadJSON(&we)
		if err != nil {
			p.writeError(err)
			return
		}

		data, err := json.Marshal(we.Data)
		if err != nil {
			p.writeError(err)
			return
		}

		e := eventsByName[we.Type](p)
		err = json.Unmarshal(data, e)
		if err != nil {
			p.writeError(err)
			return
		}

		g.events <- e
	}
}
