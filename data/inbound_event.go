package data

import (
	"encoding/json"
	"errors"
	"fmt"
)

type InboundEventType string

const (
	InboundPlayCard InboundEventType = "playCard"
)

type InboundPayload struct {
	Type InboundEventType `json:"type"`
	Data InboundEvent     `json:"data"`
}

func setInboundData[E InboundEvent](i *InboundPayload, dataBytes []byte) error {
	var e E
	err := json.Unmarshal(dataBytes, &e)
	if err != nil {
		return err
	}
	i.Data = e
	return nil
}

func (i *InboundPayload) UnmarshalJSON(bytes []byte) (err error) {
	typeString, dataBytes, err := unmarshalEvent(bytes)
	if err != nil {
		return
	}

	i.Type = InboundEventType(typeString)
	switch InboundEventType(typeString) {
	case InboundPlayCard:
		err = setInboundData[PlayCardEvent](i, dataBytes)
	default:
		return errors.New(fmt.Sprintf("%s is not a valid type", typeString))
	}

	return nil
}

type InboundEvent interface {
	IsInboundEvent()
}

type PlayCardEvent struct {
	Card Card `json:"card"`
}

func (PlayCardEvent) IsInboundEvent() {}
