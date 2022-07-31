package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestThings(t *testing.T) {
	hand1 := poker.NewPokerHand(
		poker.Hand{
			{
				poker.Ace,
				poker.Club,
			},
			{
				poker.Queen,
				poker.Spade,
			},
			{
				poker.King,
				poker.Spade,
			},
			{
				poker.Jack,
				poker.Spade,
			},
			{
				poker.Nine,
				poker.Heart,
			},
		},
	)

	hand2 := poker.NewPokerHand(
		poker.Hand{
			{
				poker.Two,
				poker.Club,
			},
			{
				poker.Two,
				poker.Spade,
			},
			{
				poker.Three,
				poker.Spade,
			},
			{
				poker.Four,
				poker.Spade,
			},
			{
				poker.Five,
				poker.Heart,
			},
		},
	)

	score1 := hand1.Evaluate()
	score2 := hand2.Evaluate()

	fmt.Println(hand1.Cards, score1)
	fmt.Println(hand2.Cards, score2)
}
