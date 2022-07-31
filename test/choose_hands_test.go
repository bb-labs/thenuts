package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
	"golang.org/x/exp/slices"
)

// func TestHandCounts(t *testing.T) {
// 	deck := poker.NewDeck()
// 	allHands := poker.ChooseN(deck, 5)

// 	tests := []struct {
// 		PokerHandType poker.PokerHandType
// 		ExpectedCount int
// 	}{
// 		{
// 			poker.HighCard,
// 			1_302_540,
// 		},
// 		{
// 			poker.Pair,
// 			1_098_240,
// 		},
// 		{
// 			poker.TwoPair,
// 			123_552,
// 		},
// 		{
// 			poker.ThreeOfAKind,
// 			54_912,
// 		},
// 		{
// 			poker.Straight,
// 			10_200,
// 		},
// 		{
// 			poker.Flush,
// 			5_108,
// 		},
// 		{
// 			poker.FullHouse,
// 			3_744,
// 		},
// 		{
// 			poker.FourOfAKind,
// 			624,
// 		},
// 		{
// 			poker.StraightFlush,
// 			40,
// 		},
// 	}

// 	for _, test := range tests {
// 		targetHands := []poker.Hand{}

// 		for _, hand := range allHands {
// 			pokerHand := poker.NewPokerHand(hand)

// 			if pokerHand.Type == test.PokerHandType {
// 				targetHands = append(targetHands, pokerHand.Cards)
// 			}
// 		}

// 		fmt.Println(test.PokerHandType, len(targetHands), test.ExpectedCount)
// 	}
// }

func TestSimpleOrdering(t *testing.T) {
	deck := poker.NewDeck()
	allHands := poker.ChooseN(deck, 5)

	tests := []struct {
		PokerHandType        poker.PokerHandType
		LesserPokerHandTypes []poker.PokerHandType
	}{
		{
			poker.Pair,
			[]poker.PokerHandType{
				poker.HighCard,
			},
		},
		{
			poker.TwoPair,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
			},
		},
		{
			poker.ThreeOfAKind,
			[]poker.PokerHandType{poker.HighCard,
				poker.Pair,
				poker.TwoPair,
			},
		},
		{
			poker.Straight,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
			},
		},
		{
			poker.Flush,
			[]poker.PokerHandType{poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
			},
		},
		{
			poker.FullHouse,
			[]poker.PokerHandType{poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
				poker.Flush,
			},
		},
		{
			poker.FourOfAKind,
			[]poker.PokerHandType{poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
				poker.Flush,
				poker.FullHouse,
			},
		},
		{
			poker.StraightFlush,
			[]poker.PokerHandType{poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
				poker.Flush,
				poker.FullHouse,
				poker.FourOfAKind,
			},
		},
	}

	for _, test := range tests[:1] {
		greaterHands := []poker.PokerHand{}
		lesserHands := []poker.PokerHand{}

		for _, hand := range allHands {
			pokerHand := poker.NewPokerHand(hand)

			if pokerHand.Type == test.PokerHandType {
				greaterHands = append(greaterHands, pokerHand)
			} else if slices.Contains(test.LesserPokerHandTypes, pokerHand.Type) {
				lesserHands = append(lesserHands, pokerHand)
			}
		}

		for i, greaterHand := range greaterHands {
			for _, lesserHand := range lesserHands {
				if greaterHand.Evaluate() < lesserHand.Evaluate() {
					fmt.Println("ERROR", greaterHand, lesserHand)
					return
				}
			}
			fmt.Println(i)
		}

	}
}
