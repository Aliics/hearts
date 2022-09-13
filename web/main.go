package web

import (
	"fmt"
	. "github.com/aliics/hearts/dom"
)

func Run() {
	keepAlive := make(chan any)
	Body.AppendChildren(
		Div(Style("text-align: center"))(
			Div(DisplayFlex)(
				Input(TypeText)().
					AddEventListener(EventTypeInput, func(element Element) {
						fmt.Printf("Input is %s\n", element.Get("value").String())
					}),
				Input(TypeButton, Value("Join game"))().
					AddEventListener(EventTypeClick, func(element Element) {
						fmt.Printf("Clicked %s\n", element.Get("value").String())
					}),
			),
			Input(TypeButton, Value("Create game"))(),
		),
	)

	<-keepAlive // Gross hack to keep the script alive and block forever.
}
