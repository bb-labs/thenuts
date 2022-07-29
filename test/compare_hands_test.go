package test

import (
	"testing"

	"github.com/bb-labs/pokerface"
)

func TestCompareHands(t *testing.T) {
	deck := pokerface.NewDeck()

	handA := pokerface.Hand{deck.Cards[2], deck.Cards[44]}
	handB := pokerface.Hand{deck.Cards[32], deck.Cards[26]}

	pokerface.Compare(handA, handB)
}
