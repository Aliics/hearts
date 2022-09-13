package web

import (
	"fmt"
	. "github.com/aliics/hearts/web/dom"
	"syscall/js"
)

func Run() {
	keepAlive := make(chan any)
	Body.AppendChildren(menuScreen())

	<-keepAlive // Gross hack to keep the script alive and block forever.
}

func menuScreen() Element {
	gameIDInput := Input(TypeText, Placeholder("Game ID"))()

	return Div(DisplayFlex, FlexDirectionColumn, Margin("auto"), Width("fit-content"))(
		Div(DisplayFlex)(
			gameIDInput,
			Input(TypeButton, Value("Join Game"))().
				AddEventListener(EventTypeClick, func(_ Element) {
					beginGame(gameIDInput.Get("value").String())
				}),
		),

		P()(StringLiteral("or")),

		Input(TypeButton, Value("Create Game"))().
			AddEventListener(EventTypeClick, func(_ Element) {
				req := js.Global().Get("XMLHttpRequest").New()
				req.Call("open", "POST", "http://localhost:8080/game/")
				req.Call("send")
				req.Set("onload", func() {})
			}),
	)
}

func beginGame(gameID string) {
	Element{Value: js.Global().
		Get("WebSocket").
		New(fmt.Sprintf("ws://localhost:8080/game/%s/", gameID))}.
		AddEventListener(EventTypeOpen, func(_ Element) {
			fmt.Println("Websocket connected!")
		})
}
