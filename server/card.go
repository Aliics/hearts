package server

import (
	"hearts/data"
	"math/rand"
	"time"
)

func newShuffledDeck() []data.Card {
	rand.Seed(time.Now().UnixNano())

	var cards []data.Card
	for suit := data.SuitHearts; suit <= data.SuitSpades; suit++ {
		for value := data.ValueTwo; value <= data.ValueAce; value++ {
			if suit == data.SuitDiamonds && value == data.ValueTwo {
				continue
			}
			cards = append(cards, data.Card{Suit: suit, Value: value})
		}
	}

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}
