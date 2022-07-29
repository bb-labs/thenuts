package test

import (
	"testing"

	"github.com/bb-labs/pokerface"
)

func TestCompareHands(t *testing.T) {
	handA := pokerface.Hand{
		pokerface.Card{Rank: pokerface.Ace, Suit: pokerface.Spade},
		pokerface.Card{Rank: pokerface.Ace, Suit: pokerface.Club},
	}

	handB := pokerface.Hand{
		pokerface.Card{Rank: pokerface.Ace, Suit: pokerface.Diamond},
		pokerface.Card{Rank: pokerface.Ace, Suit: pokerface.Heart},
	}

	pokerface.Compare(handA, handB)
}
