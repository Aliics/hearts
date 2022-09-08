package main

type Card struct {
	Suite `json:"suite"`
	Value `json:"value"`
}

type Suite uint8

const (
	SuiteHearts = iota
	SuiteDiamonds
	SuiteClubs
	SuiteSpades
)

type Value uint8

const (
	ValueTwo = iota
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
