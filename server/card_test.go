package server

import (
	"github.com/aliics/hearts/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newShuffledDeck(t *testing.T) {
	deck := newShuffledDeck()

	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueTwo})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueThree})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueFour})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueFive})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueSix})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueSeven})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueEight})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueNine})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueTen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueJack})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueQueen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueKing})
	assert.Contains(t, deck, data.Card{Suit: data.SuitHearts, Value: data.ValueAce})

	assert.NotContains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueTwo})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueThree})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueFour})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueFive})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueSix})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueSeven})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueEight})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueNine})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueTen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueJack})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueQueen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueKing})
	assert.Contains(t, deck, data.Card{Suit: data.SuitDiamonds, Value: data.ValueAce})

	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueTwo})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueThree})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueFour})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueFive})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueSix})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueSeven})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueEight})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueNine})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueTen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueJack})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueQueen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueKing})
	assert.Contains(t, deck, data.Card{Suit: data.SuitClubs, Value: data.ValueAce})

	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueTwo})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueThree})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueFour})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueFive})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueSix})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueSeven})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueEight})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueNine})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueTen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueJack})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueQueen})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueKing})
	assert.Contains(t, deck, data.Card{Suit: data.SuitSpades, Value: data.ValueAce})
}
