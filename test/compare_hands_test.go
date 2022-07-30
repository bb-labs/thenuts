package test

import (
	"testing"

	"github.com/bb-labs/poker"
)

func TestCompareHands(t *testing.T) {
	handA := poker.Hand{
		poker.Card{Rank: poker.Ace, Suit: poker.Spade},
		poker.Card{Rank: poker.Ace, Suit: poker.Club},
		poker.Card{Rank: poker.King, Suit: poker.Club},
		poker.Card{Rank: poker.King, Suit: poker.Diamond},
		poker.Card{Rank: poker.King, Suit: poker.Spade},
	}

	handB := poker.Hand{
		poker.Card{Rank: poker.Ace, Suit: poker.Diamond},
		poker.Card{Rank: poker.Ace, Suit: poker.Heart},
		poker.Card{Rank: poker.Ace, Suit: poker.Spade},
		poker.Card{Rank: poker.Queen, Suit: poker.Club},
		poker.Card{Rank: poker.Queen, Suit: poker.Heart},
	}

	poker.Compare(handA, handB)
}
