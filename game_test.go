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
		want bool
	}{
		{
			"last card and nothing played",
			args{
				nil,
				[]Card{{SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			true,
		},
		{
			"last card and different suit played",
			args{
				[]Card{{SuitClubs, ValueKing}},
				[]Card{{SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			true,
		},
		{
			"many cards of suit and nothing played",
			args{
				nil,
				[]Card{{SuitDiamonds, ValueKing}, {SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			true,
		},
		{
			"many cards of suit and cards played",
			args{
				[]Card{{SuitDiamonds, ValueEight}},
				[]Card{{SuitDiamonds, ValueKing}, {SuitDiamonds, ValueAce}},
				Card{SuitDiamonds, ValueAce},
			},
			true,
		},
		{
			"not matching suits but suit is in hand",
			args{
				[]Card{{SuitSpades, ValueAce}},
				[]Card{{SuitDiamonds, ValueKing}, {SuitSpades, ValueSeven}},
				Card{SuitDiamonds, ValueKing},
			},
			false,
		},
		{
			"not matching suits but no suit in hand",
			args{
				[]Card{{SuitSpades, ValueAce}},
				[]Card{{SuitDiamonds, ValueKing}},
				Card{SuitDiamonds, ValueKing},
			},
			true,
		},
		{
			"playerId does not have card",
			args{
				nil,
				[]Card{{SuitDiamonds, ValueKing}},
				Card{SuitSpades, ValueKing},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validCardPlayed(tt.args.inPlay, tt.args.hand, tt.args.played))
		})
	}
}
