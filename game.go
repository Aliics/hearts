package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	games = make(map[uuid.UUID]*game)
)

type game struct {
	events  chan event
	players []player
	inPlay  []Card
}

func (g game) run() {
	for x := range g.events {
		switch e := x.(type) {
		case *playCardEvent:
			fmt.Println(e.Card)
		case connectPlayerEvent:
			if len(g.players) == 3 {
				logNonFatal(player(e).WriteMessage(websocket.TextMessage, []byte("game is full")))
				continue
			}
			g.players = append(g.players, player(e))
		default:
			logNonFatal(e.player().WriteMessage(websocket.TextMessage, []byte("event handler not found")))
		}
	}
}

func (g game) connectPlayer(p player) {
	g.events <- connectPlayerEvent(p)
}

type player struct {
	*websocket.Conn
	hand []Card
}
