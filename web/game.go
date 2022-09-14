package web

import (
	"fmt"
	. "github.com/aliics/hearts/web/dom"
	"syscall/js"
)

func createGameScreen(gameID string) Element {
	ws := NewWebSocket(fmt.Sprintf("ws://localhost:8080/game/%s/", gameID)).
		AddEventListener(EventTypeOpen, func(js.Value, []js.Value) {
			fmt.Println("Websocket connected!")
		}).
		AddEventListener(EventTypeMessage, func(_ js.Value, messages []js.Value) {
			for _, message := range messages {
				fmt.Println(message.Get("data"))
			}
		})

	return Div()(
		P()(StringLiteral(fmt.Sprintf("Game ID: %s", gameID))),
		Button()(StringLiteral("<< Exit")).
			AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
				ws.Call("close")
				screenChangeCh <- screenChange{screenType: screenMenu}
			}),
	)
}
