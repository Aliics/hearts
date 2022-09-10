package main

import (
	"math/rand"
	"reflect"
	"time"
)

func newShuffledDeck() []Card {
	rand.Seed(time.Now().UnixNano())

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

func (c Card) points() int {
	if c.Suit == SuitHearts {
		return 1
	} else if reflect.DeepEqual(c, Card{SuitSpades, ValueQueen}) {
		return 13
	}
	return 0
}

func (c Card) beats(other Card) bool {
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
