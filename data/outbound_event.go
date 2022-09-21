package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type OutboundEventType string

const (
	OutboundClientViolation OutboundEventType = "clientViolation"
	OutboundNewHand         OutboundEventType = "newHand"
	OutboundCurrentPlayer   OutboundEventType = "currentPlayer"
	OutboundInPlay          OutboundEventType = "inPlay"
	OutboundCurrentPlayers  OutboundEventType = "currentPlayers"
	OutboundPoints          OutboundEventType = "points"
)

type OutboundPayload struct {
	Type OutboundEventType `json:"type"`
	Data OutboundEvent     `json:"data"`
}

func setOutboundData[E OutboundEvent](o *OutboundPayload, dataBytes []byte) error {
	var e E
	err := json.Unmarshal(dataBytes, &e)
	if err != nil {
		return err
	}
	o.Data = e
	return nil
}

func (o *OutboundPayload) UnmarshalJSON(bytes []byte) (err error) {
	typeString, dataBytes, err := unmarshalEvent(bytes)
	if err != nil {
		return err
	}

	o.Type = OutboundEventType(typeString)
	switch OutboundEventType(typeString) {
	case OutboundClientViolation:
		err = setOutboundData[ClientViolationEvent](o, dataBytes)
	case OutboundNewHand:
		err = setOutboundData[NewHandEvent](o, dataBytes)
	case OutboundCurrentPlayer:
		err = setOutboundData[CurrentPlayerEvent](o, dataBytes)
	case OutboundInPlay:
		err = setOutboundData[InPlayEvent](o, dataBytes)
	case OutboundCurrentPlayers:
		err = setOutboundData[CurrentPlayersEvent](o, dataBytes)
	case OutboundPoints:
		err = setOutboundData[PointsEvent](o, dataBytes)
	default:
		return errors.New(fmt.Sprintf("%s is not a valid type", typeString))
	}

	return err
}

type OutboundEvent interface {
	IsOutboundEvent()
}

type ClientViolationEvent struct {
	Message string `json:"message"`
}

func (ClientViolationEvent) IsOutboundEvent() {}

type NewHandEvent struct {
	Hand []Card `json:"hand"`
}

func (NewHandEvent) IsOutboundEvent() {}

type CurrentPlayerEvent struct {
	PlayerID uuid.UUID `json:"playerID"`
}

func (CurrentPlayerEvent) IsOutboundEvent() {}

type InPlayEvent struct {
	InPlay []PlayerCard `json:"inPlay"`
}

func (InPlayEvent) IsOutboundEvent() {}

type CurrentPlayersEvent struct {
	PlayerIDs []uuid.UUID `json:"players"`
	GameReady bool        `json:"gameReady"`
}

func (CurrentPlayersEvent) IsOutboundEvent() {}

type PointsEvent struct {
	Points map[uuid.UUID]int `json:"points"`
}

func (PointsEvent) IsOutboundEvent() {}
