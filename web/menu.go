package web

import (
	. "github.com/aliics/hearts/web/dom"
	"net/http"
	"strings"
	"syscall/js"
)

func createMenuScreen() Element {
	var gameIDInput, joinGameButton Element
	gameIDInput = Input(TypeText, PlaceholderAttribute("Game ID"))().
		AddEventListener(EventTypeInput, func(js.Value, []js.Value) {
			inputFilled := strings.TrimSpace(gameIDInput.Get("value").String()) == ""
			joinGameButton.Set("disabled", inputFilled)
		})
	joinGameButton = Button(TypeButton, DisabledAttribute(true))(StringLiteral("Join Game")).
		AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
			beginGame(gameIDInput.Get("value").String())
		})

	return Div(DisplayFlex, FlexDirectionColumn, MarginAttribute("auto"), WidthAttribute("fit-content"))(
		Div(DisplayFlex)(gameIDInput, joinGameButton),

		P(StyleAttribute("text-align: center"))(StringLiteral("or")),

		Input(TypeButton, ValueAttribute("Create Game"))().
			AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
				req := NewXHR(http.MethodPost, "http://localhost:8080/game/")
				req.AddEventListener(EventTypeLoad, func(js.Value, []js.Value) {
					if req.Get("status").Int() == 200 {
						beginGame(req.Get("response").String())
					}
				})
				req.Call("send")
			}),
	)
}
