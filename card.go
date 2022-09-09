package main

type Card struct {
	Suit  `json:"suit"`
	Value `json:"value"`
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
