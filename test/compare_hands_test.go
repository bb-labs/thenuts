package test

import (
	"testing"

	"github.com/bb-labs/poker"
)

func TestCompareHands(t *testing.T) {
	handA := poker.Hand{
		poker.Card{Rank: poker.Ace, Suit: poker.Spade},
		poker.Card{Rank: poker.Ace, Suit: poker.Club},
	}

	handB := poker.Hand{
		poker.Card{Rank: poker.Ace, Suit: poker.Diamond},
		poker.Card{Rank: poker.Ace, Suit: poker.Heart},
	}

	poker.Compare(handA, handB)
}
