package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validCardPlayed(t *testing.T) {
	type args struct {
		inPlay []Card
		hand   []Card
		played Card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"last card and nothing played",
			args{
				nil,
				[]Card{{SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			0,
		},
		{
			"last card and different suit played",
			args{
				[]Card{{SuitClubs, ValueKing}},
				[]Card{{SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			0,
		},
		{
			"many cards of suit and nothing played",
			args{
				nil,
				[]Card{{SuitDiamonds, ValueKing}, {SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			1,
		},
		{
			"many cards of suit and cards played",
			args{
				[]Card{{SuitDiamonds, ValueEight}},
				[]Card{{SuitDiamonds, ValueKing}, {SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			1,
		},
		{
			"not matching suits but suit is in hand",
			args{
				[]Card{{SuitSpades, ValueAce}},
				[]Card{{SuitDiamonds, ValueKing}, {SuitSpades, ValueSeven}},
				Card{SuitDiamonds, ValueKing},
			},
			-1,
		},
		{
			"not matching suits but no suit in hand",
			args{
				[]Card{{SuitSpades, ValueAce}},
				[]Card{{SuitDiamonds, ValueKing}},
				Card{SuitDiamonds, ValueKing},
			},
			0,
		},
		{
			"playerId does not have card",
			args{
				nil,
				[]Card{{SuitDiamonds, ValueKing}},
				Card{SuitSpades, ValueKing},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, indexOfValidPlayedCard(tt.args.inPlay, tt.args.hand, tt.args.played))
		})
	}
}
