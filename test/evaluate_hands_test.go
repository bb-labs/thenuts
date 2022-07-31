package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestEvaluateHands(t *testing.T) {
	hand1 := poker.NewPokerHand(
		poker.Hand{
			{
				poker.Five,
				poker.Club,
			},
			{
				poker.Four,
				poker.Spade,
			},
			{
				poker.Three,
				poker.Spade,
			},
			{
				poker.Two,
				poker.Spade,
			},
			{
				poker.Ace,
				poker.Heart,
			},
		},
	)

	hand2 := poker.NewPokerHand(
		poker.Hand{
			{
				poker.Three,
				poker.Club,
			},
			{
				poker.Four,
				poker.Spade,
			},
			{
				poker.Five,
				poker.Spade,
			},
			{
				poker.Six,
				poker.Spade,
			},
			{
				poker.Seven,
				poker.Heart,
			},
		},
	)

	score1 := hand1.Evaluate()
	score2 := hand2.Evaluate()

	fmt.Println(hand1.Cards, score1)
	fmt.Println(hand2.Cards, score2)
}
