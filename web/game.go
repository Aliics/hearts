package web

import (
	"encoding/json"
	"fmt"
	"github.com/aliics/hearts/data"
	"github.com/aliics/hearts/util"
	. "github.com/aliics/hearts/web/dom"
	"github.com/google/uuid"
	"log"
	"syscall/js"
)

var ws Element

func createGameScreen(gameID string) Element {
	var (
		playerId           uuid.UUID
		currentHandElement = Div()()
		notReadyElement    = Div()()
		inPlayElement      = Div()()
	)

	gameElement := Div()(
		P()(StringLiteral(fmt.Sprintf("Game ID: %s", gameID))),
		Button()(StringLiteral("<< Exit")).
			AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
				ws.Call("close")
				screenChangeCh <- screenChange{screenType: screenMenu}
			}),
		Div(DisplayFlex)(
			notReadyElement,
			currentHandElement,
			inPlayElement,
		),
	)

	ws = NewWebSocket(fmt.Sprintf("ws://localhost:8080/game/%s/", gameID)).
		AddEventListener(EventTypeOpen, func(js.Value, []js.Value) {
			fmt.Println("Websocket connected!")
		}).
		AddEventListener(EventTypeMessage, func(_ js.Value, messages []js.Value) {
			for _, message := range messages {
				log.Println(message.Get("data"))

				var payload data.OutboundPayload
				util.Try0(json.Unmarshal([]byte(message.Get("data").String()), &payload))

				switch event := payload.Data.(type) {
				case data.CurrentPlayersEvent:
					if playerId == [16]byte{} {
						playerId = event.PlayerIDs[len(event.PlayerIDs)-1]
						fmt.Println("Your playerID is:", playerId)
					}

					if event.GameReady {
						notReadyElement = notReadyElement.ReplacedWithEmpty()
					} else {
						inPlayElement = inPlayElement.ReplacedWithEmpty()
						currentHandElement = currentHandElement.ReplacedWithEmpty()
						notReadyElement = notReadyElement.Replaced(
							Div()(
								P()(StringLiteral("Waiting for players")),
								P()(StringLiteral(fmt.Sprintf("Current Players: %d/3", len(event.PlayerIDs)))),
							),
						)
					}
				case data.NewHandEvent:
					var innerElements []Element
					innerElements = append(
						innerElements,
						P()(StringLiteral(fmt.Sprintf("Your Hand (%s)", playerId.String()))),
					)

					if len(event.Hand) > 0 {
						for _, card := range event.Hand {
							innerElements = append(innerElements, cardElement(card))
						}
					} else {
						innerElements = append(innerElements, StringLiteral("Empty Hand."))
					}

					currentHandElement = currentHandElement.Replaced(
						Div(DisplayFlex, FlexDirectionColumn)(innerElements...),
					)
				case data.InPlayEvent:
					var cardElements []Element
					for _, card := range event.InPlay {
						cardElements = append(cardElements, cardElement(card.Card))
					}

					inPlayElement = inPlayElement.Replaced(
						Div()(cardElements...),
					)
				}
			}
		})

	return gameElement
}

func cardElement(card data.Card) Element {
	var color string
	if card.Suit == data.SuitHearts || card.Suit == data.SuitDiamonds {
		color = "#f00"
	} else {
		color = "#000"
	}
	return Div(ColorAttribute(color))(StringLiteral(fmt.Sprintf("%d %d", card.Suit, card.Value))).
		AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
			playCard, err := json.Marshal(data.InboundPayload{
				Type: data.InboundPlayCard,
				Data: data.PlayCardEvent{Card: card},
			})
			util.Try0(err)
			ws.Call("send", string(playCard))
		})
}
