package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestHandCounts(t *testing.T) {
	deck := poker.NewDeck()

	tests := []struct {
		PokerHandType poker.PokerHandType
		ExpectedCount int
	}{
		{
			poker.HighCard,
			1_302_540,
		},
		{
			poker.Pair,
			1_098_240,
		},
		{
			poker.TwoPair,
			123_552,
		},
		{
			poker.ThreeOfAKind,
			54_912,
		},
		{
			poker.Straight,
			10_200,
		},
		{
			poker.Flush,
			5_108,
		},
		{
			poker.FullHouse,
			3_744,
		},
		{
			poker.FourOfAKind,
			624,
		},
		{
			poker.StraightFlush,
			40,
		},
	}

	for _, test := range tests {
		targetHands := []poker.Hand{}
		allHands := poker.ChooseN(deck, 5)

		for _, hand := range allHands {
			pokerHand := poker.NewPokerHand(hand)

			if pokerHand.Type == test.PokerHandType {
				targetHands = append(targetHands, pokerHand.Cards)
			}
		}

		fmt.Println(test.PokerHandType, len(targetHands), test.ExpectedCount)
	}
}
