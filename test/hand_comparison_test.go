package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/bb-labs/poker"
	"golang.org/x/exp/slices"
)

func TestSimpleHandComparisons(t *testing.T) {
	deck := poker.NewDeck()
	allHands := poker.ChooseN(deck, 5)

	tests := []struct {
		PokerHandType        poker.PokerHandType
		ExpectedCount        int
		LesserPokerHandTypes []poker.PokerHandType
		LowestHandRanks      []poker.Rank
		HighestHandRanks     []poker.Rank
	}{
		{
			poker.HighCard,
			1_302_540,
			[]poker.PokerHandType{},
			[]poker.Rank{poker.Two, poker.Three, poker.Four, poker.Five, poker.Seven},
			[]poker.Rank{poker.Ace, poker.King, poker.Queen, poker.Jack, poker.Nine},
		},
		{
			poker.Pair,
			1_098_240,
			[]poker.PokerHandType{
				poker.HighCard,
			},
			[]poker.Rank{poker.Two, poker.Two, poker.Three, poker.Four, poker.Five},
			[]poker.Rank{poker.Ace, poker.Ace, poker.King, poker.Queen, poker.Jack},
		},
		{
			poker.TwoPair,
			123_552,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
			},
			[]poker.Rank{poker.Two, poker.Two, poker.Three, poker.Three, poker.Four},
			[]poker.Rank{poker.Ace, poker.Ace, poker.King, poker.King, poker.Queen},
		},
		{
			poker.ThreeOfAKind,
			54_912,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
			},
			[]poker.Rank{poker.Two, poker.Two, poker.Two, poker.Three, poker.Four},
			[]poker.Rank{poker.Ace, poker.Ace, poker.Ace, poker.King, poker.Queen},
		},
		{
			poker.Straight,
			10_200,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
			},
			[]poker.Rank{poker.LowAce, poker.Two, poker.Three, poker.Four, poker.Five},
			[]poker.Rank{poker.Ace, poker.King, poker.Queen, poker.Jack, poker.Ten},
		},
		{
			poker.Flush,
			5_108,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
			},
			[]poker.Rank{poker.Two, poker.Three, poker.Four, poker.Five, poker.Seven},
			[]poker.Rank{poker.Ace, poker.King, poker.Queen, poker.Jack, poker.Nine},
		},
		{
			poker.FullHouse,
			3_744,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
				poker.Flush,
			},
			[]poker.Rank{poker.Two, poker.Two, poker.Two, poker.Three, poker.Three},
			[]poker.Rank{poker.Ace, poker.Ace, poker.Ace, poker.King, poker.King},
		},
		{
			poker.FourOfAKind,
			624,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
				poker.Flush,
				poker.FullHouse,
			},
			[]poker.Rank{poker.Two, poker.Two, poker.Two, poker.Two, poker.Three},
			[]poker.Rank{poker.Ace, poker.Ace, poker.Ace, poker.Ace, poker.King},
		},
		{
			poker.StraightFlush,
			40,
			[]poker.PokerHandType{
				poker.HighCard,
				poker.Pair,
				poker.TwoPair,
				poker.ThreeOfAKind,
				poker.Straight,
				poker.Flush,
				poker.FullHouse,
				poker.FourOfAKind,
			},
			[]poker.Rank{poker.LowAce, poker.Two, poker.Three, poker.Four, poker.Five},
			[]poker.Rank{poker.Ace, poker.King, poker.Queen, poker.Jack, poker.Ten},
		},
	}

	for _, test := range tests {
		hands := []poker.PokerHand{}
		lesserHands := []poker.PokerHand{}

		for _, hand := range allHands {
			pokerHand := poker.NewPokerHand(hand)

			if pokerHand.Type == test.PokerHandType {
				hands = append(hands, pokerHand)
			} else if slices.Contains(test.LesserPokerHandTypes, pokerHand.Type) {
				lesserHands = append(lesserHands, pokerHand)
			}
		}

		sort.Slice(hands, func(i, j int) bool {
			return hands[i].Evaluate() < hands[j].Evaluate()
		})

		// Each type should have the proper count
		if test.ExpectedCount != len(hands) {
			t.Errorf("%v Got %v; want %v", test.PokerHandType, len(hands), test.ExpectedCount)
		}

		// Lowest and highest hand for the given type
		lowestHand := hands[0]
		highestHand := hands[len(hands)-1]

		fmt.Println(lowestHand.Cards, lowestHand.Evaluate())
		fmt.Println(highestHand.Cards, highestHand.Evaluate())

		// Check that these are the proper hands
		for i := 0; i < 5; i++ {
			lowHandCard := lowestHand.Cards[i]
			highHandCard := highestHand.Cards[i]

			if contains := slices.Contains(test.LowestHandRanks, lowHandCard.Rank); !contains {
				t.Errorf("%v Got %v; want %v", test.PokerHandType, lowestHand.Cards, test.LowestHandRanks)
			}

			if contains := slices.Contains(test.HighestHandRanks, highHandCard.Rank); !contains {
				t.Errorf("%v Got %v; want %v", test.PokerHandType, highestHand.Cards, test.HighestHandRanks)
			}
		}

		// All cards of lesser rank should be lower than the lowest hand of the current type
		for _, hand := range lesserHands {
			if hand.Evaluate() >= lowestHand.Evaluate() {
				t.Errorf("(%v) Hand %v; should not be greater than %v", test.PokerHandType, hand.Cards, lowestHand.Cards)
			}
		}
	}
}
