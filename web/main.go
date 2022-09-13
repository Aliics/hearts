//go:build wasm && js

package web

import (
	"fmt"
	. "github.com/aliics/hearts/web/dom"
	"net/http"
	"syscall/js"
)

func Run() {
	keepAlive := make(chan any)
	DocumentBody.AppendChildren(menuScreen())

	<-keepAlive // Gross hack to keep the script alive and block forever.
}

func menuScreen() Element {
	gameIDInput := Input(TypeText, PlaceholderAttribute("Game ID"))()

	return Div(DisplayFlex, FlexDirectionColumn, MarginAttribute("auto"), WidthAttribute("fit-content"))(
		Div(DisplayFlex)(
			gameIDInput,
			Input(TypeButton, ValueAttribute("Join Game"))().
				AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
					beginGame(gameIDInput.Get("value").String())
				}),
		),

		P()(StringLiteral("or")),

		Input(TypeButton, ValueAttribute("Create Game"))().
			AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
				req := NewXHR(http.MethodPost, "http://localhost:8080/game/")
				req.Call("send")
				req.Set("onload", func() {})
			}),
	)
}

func beginGame(gameID string) {
	NewWebSocket(fmt.Sprintf("ws://localhost:8080/game/%s/", gameID)).
		AddEventListener(EventTypeOpen, func(js.Value, []js.Value) {
			fmt.Println("Websocket connected!")
		}).
		AddEventListener(EventTypeMessage, func(_ js.Value, messages []js.Value) {
			for _, message := range messages {
				fmt.Println(message.Get("data"))
			}
		})
}
