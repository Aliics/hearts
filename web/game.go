package web

import (
	"encoding/json"
	"fmt"
	"github.com/aliics/hearts/data"
	"github.com/aliics/hearts/util"
	. "github.com/aliics/hearts/web/dom"
	"github.com/google/uuid"
	"strconv"
	"syscall/js"
)

var (
	ws              Element
	playerID        uuid.UUID
	currentPlayerID uuid.UUID
	playerIDs       []uuid.UUID
	points          = make(map[uuid.UUID]int)
)

func createGameScreen(gameID string) Element {
	var (
		notReadyElement     = Empty()
		playersHandsElement = Empty()
		currentHandElement  = Empty()
		inPlayElement       = Empty()
	)

	gameElement := Div()(
		Div()(StringLiteral(fmt.Sprintf("Game ID: %s", gameID))),
		notReadyElement,
		Div(DisplayGrid)(
			playersHandsElement,
			inPlayElement,
			currentHandElement,
		),
	)

	ws = NewWebSocket(fmt.Sprintf("ws://localhost:8080/game/%s/", gameID)).
		AddEventListener(EventTypeOpen, func(js.Value, []js.Value) {
			fmt.Println("Websocket connected!")
		}).
		AddEventListener(EventTypeMessage, func(_ js.Value, messages []js.Value) {
			for _, message := range messages {
				var payload data.OutboundPayload
				util.Try0(json.Unmarshal([]byte(message.Get("data").String()), &payload))

				switch event := payload.Data.(type) {
				case data.CurrentPlayersEvent:
					playerIDs = event.PlayerIDs
					if playerID == [16]byte{} {
						playerID = event.PlayerIDs[len(event.PlayerIDs)-1]
						fmt.Println("Your Player ID is:", playerID)
					}

					if event.GameReady {
						notReadyElement = notReadyElement.ReplacedWithEmpty()
						playersHandsElement = playersHandsElement.Replaced(createPlayersHandElement())
					} else {
						playersHandsElement = playersHandsElement.ReplacedWithEmpty()
						inPlayElement = inPlayElement.ReplacedWithEmpty()
						currentHandElement = currentHandElement.ReplacedWithEmpty()
						notReadyElement = notReadyElement.Replaced(
							centeredModal(
								P()(StringLiteral("Waiting for players...")),
								P()(StringLiteral(fmt.Sprintf("Current Players: %d/3", len(event.PlayerIDs)))),
							),
						)
					}
				case data.CurrentPlayerEvent:
					currentPlayerID = event.PlayerID
					playersHandsElement = playersHandsElement.Replaced(createPlayersHandElement())
				case data.NewHandEvent:
					var innerElements []Element
					if len(event.Hand) > 0 {
						for _, card := range event.Hand {
							playableCard := createCardElement(card, true)
							innerElements = append(innerElements, playableCard)
						}
					} else {
						innerElements = append(innerElements, StringLiteral("Empty Hand."))
					}

					currentHandElement = currentHandElement.Replaced(
						Div(
							DisplayFlex,
							PositionAttribute("absolute"),
							WidthAttribute("100%"),
							StyleAttribute("bottom: 0; flex-wrap: wrap; justify-content: center;"),
						)(innerElements...),
					)
				case data.InPlayEvent:
					var cardElements []Element
					for _, card := range event.InPlay {
						cardElements = append(cardElements, createCardElement(card.Card, false))
					}

					inPlayElement = inPlayElement.Replaced(
						Div(DisplayFlex, StyleAttribute("justify-content: center;"))(cardElements...),
					)
				case data.PointsEvent:
					points = event.Points
					inPlayElement = inPlayElement.ReplacedWithEmpty()
					playersHandsElement = playersHandsElement.Replaced(createPlayersHandElement())
				}
			}
		})

	return gameElement
}

func createPlayersHandElement() Element {
	var playerHandElements []Element
	for _, pID := range playerIDs {
		if pID != playerID {
			var styling BackgroundColorAttribute
			if pID == currentPlayerID {
				styling = "#f00"
			}

			var playerPoints int
			if p, ok := points[pID]; ok {
				playerPoints = p
			}

			playerHandElements = append(playerHandElements, Div(DisplayGrid, styling)(
				Div(WidthAttribute("8em"), HeightAttribute("12em"))(
					StringLiteral(fmt.Sprintf("PlayerID %s", pID)),
				),
				StringLiteral(strconv.Itoa(playerPoints)),
			))
		}
	}
	return Div(
		DisplayFlex,
		MarginAttribute("2em"),
		StyleAttribute("justify-content: space-between;"),
	)(playerHandElements...)
}

func createCardElement(card data.Card, playable bool) Element {
	var color string
	if card.Suit == data.SuitHearts || card.Suit == data.SuitDiamonds {
		color = "#f00"
	} else {
		color = "#000"
	}

	element := Div(WidthAttribute("5em"), HeightAttribute("8em"), ColorAttribute(color))(
		StringLiteral(fmt.Sprintf("%d %d", card.Suit, card.Value)),
	)

	if playable {
		element.AddEventListener(EventTypeClick, func(js.Value, []js.Value) {
			playCard, err := json.Marshal(data.InboundPayload{
				Type: data.InboundPlayCard,
				Data: data.PlayCardEvent{Card: card},
			})
			util.Try0(err)
			ws.Call("send", string(playCard))
		})
		StyleAttribute("cursor: pointer;").Apply(&element)
	}

	return element
}
