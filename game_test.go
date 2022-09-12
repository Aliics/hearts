package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_indexOfValidPlayedCard(t *testing.T) {
	type args struct {
		inPlay []playerCard
		hand   []card
		played card
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
				[]card{{SuitDiamonds, ValueAce}},
				card{SuitDiamonds, ValueAce},
			},
			0,
		},
		{
			"last card and different suit played",
			args{
				[]playerCard{{card: card{SuitClubs, ValueKing}}},
				[]card{{SuitDiamonds, ValueAce}},
				card{SuitDiamonds, ValueAce},
			},
			0,
		},
		{
			"many cards of suit and nothing played",
			args{
				nil,
				[]card{{SuitDiamonds, ValueKing}, {SuitDiamonds, ValueAce}},
				card{SuitDiamonds, ValueAce},
			},
			1,
		},
		{
			"many cards of suit and cards played",
			args{
				[]playerCard{{card: card{SuitDiamonds, ValueEight}}},
				[]card{{SuitDiamonds, ValueKing}, {SuitDiamonds, ValueAce}},
				card{SuitDiamonds, ValueAce},
			},
			1,
		},
		{
			"not matching suits but suit is in hand",
			args{
				[]playerCard{{card: card{SuitSpades, ValueAce}}},
				[]card{{SuitDiamonds, ValueKing}, {SuitSpades, ValueSeven}},
				card{SuitDiamonds, ValueKing},
			},
			-1,
		},
		{
			"not matching suits but no suit in hand",
			args{
				[]playerCard{{card: card{SuitSpades, ValueAce}}},
				[]card{{SuitDiamonds, ValueKing}},
				card{SuitDiamonds, ValueKing},
			},
			0,
		},
		{
			"playerId does not have card",
			args{
				nil,
				[]card{{SuitDiamonds, ValueKing}},
				card{SuitSpades, ValueKing},
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

func Test_getHighestInPlay(t *testing.T) {
	type args struct {
		inPlay []playerCard
	}
	tests := []struct {
		name        string
		args        args
		wantHighest playerCard
		wantPoints  int
	}{
		{
			"all same suit, no points cards",
			args{
				[]playerCard{
					{card: card{SuitSpades, ValueThree}},
					{card: card{SuitSpades, ValueFour}},
					{card: card{SuitSpades, ValueFive}},
				},
			},
			playerCard{card: card{SuitSpades, ValueFive}},
			0,
		},
		{
			"all same suit, some hearts",
			args{
				[]playerCard{
					{card: card{SuitSpades, ValueThree}},
					{card: card{SuitHearts, ValueQueen}},
					{card: card{SuitSpades, ValueFive}},
				},
			},
			playerCard{card: card{SuitSpades, ValueFive}},
			1,
		},
		{
			"hearts round",
			args{
				[]playerCard{
					{card: card{SuitHearts, ValueThree}},
					{card: card{SuitHearts, ValueQueen}},
					{card: card{SuitHearts, ValueFive}},
				},
			},
			playerCard{card: card{SuitHearts, ValueQueen}},
			3,
		},
		{
			"queen of spades against clubs",
			args{
				[]playerCard{
					{card: card{SuitClubs, ValueTwo}},
					{card: card{SuitClubs, ValueThree}},
					{card: card{SuitSpades, ValueQueen}},
				},
			},
			playerCard{card: card{SuitClubs, ValueThree}},
			13,
		},
		{
			"queen of spades on a hearts round",
			args{
				[]playerCard{
					{card: card{SuitHearts, ValueKing}},
					{card: card{SuitHearts, ValueAce}},
					{card: card{SuitSpades, ValueQueen}},
				},
			},
			playerCard{card: card{SuitHearts, ValueAce}},
			15,
		},
		{
			"queen of spades, self inflicted",
			args{
				[]playerCard{
					{card: card{SuitSpades, ValueJack}},
					{card: card{SuitSpades, ValueTen}},
					{card: card{SuitSpades, ValueQueen}},
				},
			},
			playerCard{card: card{SuitSpades, ValueQueen}},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHighest, gotPoints := getHighestInPlay(tt.args.inPlay)
			assert.Equalf(t, tt.wantHighest, gotHighest, "getHighestInPlay(%v)", tt.args.inPlay)
			assert.Equalf(t, tt.wantPoints, gotPoints, "getHighestInPlay(%v)", tt.args.inPlay)
		})
	}
}
