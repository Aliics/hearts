package server

import (
	"github.com/aliics/hearts/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_indexOfValidPlayedCard(t *testing.T) {
	type args struct {
		inPlay []data.PlayerCard
		hand   []data.Card
		played data.Card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"last Card and nothing played",
			args{
				nil,
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueAce}},
				data.Card{Suit: data.SuitDiamonds, Value: data.ValueAce},
			},
			0,
		},
		{
			"last Card and different suit played",
			args{
				[]data.PlayerCard{{Card: data.Card{Suit: data.SuitClubs, Value: data.ValueKing}}},
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueAce}},
				data.Card{Suit: data.SuitDiamonds, Value: data.ValueAce},
			},
			0,
		},
		{
			"many cards of suit and nothing played",
			args{
				nil,
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueKing}, {Suit: data.SuitDiamonds, Value: data.ValueAce}},
				data.Card{Suit: data.SuitDiamonds, Value: data.ValueAce},
			},
			1,
		},
		{
			"many cards of suit and cards played",
			args{
				[]data.PlayerCard{{Card: data.Card{Suit: data.SuitDiamonds, Value: data.ValueEight}}},
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueKing}, {Suit: data.SuitDiamonds, Value: data.ValueAce}},
				data.Card{Suit: data.SuitDiamonds, Value: data.ValueAce},
			},
			1,
		},
		{
			"not matching suits but suit is in hand",
			args{
				[]data.PlayerCard{{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueAce}}},
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueKing}, {Suit: data.SuitSpades, Value: data.ValueSeven}},
				data.Card{Suit: data.SuitDiamonds, Value: data.ValueKing},
			},
			-1,
		},
		{
			"not matching suits but no suit in hand",
			args{
				[]data.PlayerCard{{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueAce}}},
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueKing}},
				data.Card{Suit: data.SuitDiamonds, Value: data.ValueKing},
			},
			0,
		},
		{
			"playerID does not have Card",
			args{
				nil,
				[]data.Card{{Suit: data.SuitDiamonds, Value: data.ValueKing}},
				data.Card{Suit: data.SuitSpades, Value: data.ValueKing},
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
		inPlay []data.PlayerCard
	}
	tests := []struct {
		name        string
		args        args
		wantHighest data.PlayerCard
		wantPoints  int
	}{
		{
			"all same suit, no points cards",
			args{
				[]data.PlayerCard{
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueThree}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueFour}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueFive}},
				},
			},
			data.PlayerCard{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueFive}},
			0,
		},
		{
			"all same suit, some hearts",
			args{
				[]data.PlayerCard{
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueThree}},
					{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueQueen}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueFive}},
				},
			},
			data.PlayerCard{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueFive}},
			1,
		},
		{
			"hearts round",
			args{
				[]data.PlayerCard{
					{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueThree}},
					{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueQueen}},
					{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueFive}},
				},
			},
			data.PlayerCard{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueQueen}},
			3,
		},
		{
			"queen of spades against clubs",
			args{
				[]data.PlayerCard{
					{Card: data.Card{Suit: data.SuitClubs, Value: data.ValueTwo}},
					{Card: data.Card{Suit: data.SuitClubs, Value: data.ValueThree}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueQueen}},
				},
			},
			data.PlayerCard{Card: data.Card{Suit: data.SuitClubs, Value: data.ValueThree}},
			13,
		},
		{
			"queen of spades on a hearts round",
			args{
				[]data.PlayerCard{
					{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueKing}},
					{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueAce}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueQueen}},
				},
			},
			data.PlayerCard{Card: data.Card{Suit: data.SuitHearts, Value: data.ValueAce}},
			15,
		},
		{
			"queen of spades, self inflicted",
			args{
				[]data.PlayerCard{
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueJack}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueTen}},
					{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueQueen}},
				},
			},
			data.PlayerCard{Card: data.Card{Suit: data.SuitSpades, Value: data.ValueQueen}},
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
