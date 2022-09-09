package main

import "math/rand"

func newShuffledDeck() []Card {
	var cards []Card
	for suit := SuitHearts; suit <= SuitSpades; suit++ {
		for value := ValueTwo; value <= ValueAce; value++ {
			if suit == SuitDiamonds && value == ValueTwo {
				continue
			}
			cards = append(cards, Card{suit, value})
		}
	}

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}

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
