package data

import (
	"github.com/google/uuid"
	"reflect"
)

type PlayerCard struct {
	PlayerID uuid.UUID `json:"player"`
	Card     `json:"card"`
}

type Card struct {
	Suit  `json:"suit"`
	Value `json:"value"`
}

func (c Card) Worth() int {
	if c.Suit == SuitHearts {
		return 1
	} else if reflect.DeepEqual(c, Card{SuitSpades, ValueQueen}) {
		return 13
	}
	return 0
}

func (c Card) Beats(other Card) bool {
	return c.Suit == other.Suit && c.Value > other.Value
}

type Suit uint8

const (
	SuitHearts Suit = iota
	SuitDiamonds
	SuitClubs
	SuitSpades
)

type Value uint8

const (
	ValueTwo Value = iota
	ValueThree
	ValueFour
	ValueFive
	ValueSix
	ValueSeven
	ValueEight
	ValueNine
	ValueTen
	ValueJack
	ValueQueen
	ValueKing
	ValueAce
)
