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
	events  chan event
	players []player
	inPlay  []Card
}

func (g game) run() {
	for x := range g.events {
		switch e := x.(type) {
		case *playCardEvent:
			suiteMatches := true
			for _, card := range g.inPlay {
				if card.Suite != e.Card.Suite {
					suiteMatches = false
					break
				}
			}

			var playerHasSuite bool
			for _, card := range e.playedBy.hand {
				if card.Suite == e.Card.Suite {
					playerHasSuite = true
					break
				}
			}

			if !suiteMatches && playerHasSuite {
				e.playedBy.writeMessage("cannot play off suite")
				continue
			}

			g.inPlay = append(g.inPlay, e.Card)

			for _, p := range g.players {
				inPlay, err := json.Marshal(g.inPlay)
				logNonFatal(err) // TODO: Handle more gracefully?
				p.writeMessage(string(inPlay))
			}
		case connectPlayerEvent:
			if len(g.players) == 3 {
				e.player().writeError(errors.New("game is full"))
				continue
			}
			g.players = append(g.players, player(e))
		default:
			e.player().writeMessage("event handler not found")
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

func (p player) writeMessage(msg string) {
	logNonFatal(p.WriteMessage(websocket.TextMessage, []byte(msg)))
}

func (p player) writeError(err error) {
	logNonFatal(p.WriteControl(websocket.CloseMessage, []byte(err.Error()), time.Now()))
}
