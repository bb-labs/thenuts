package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestChooseHands(t *testing.T) {
	deck := poker.NewDeck()

	hands := poker.ChooseN(deck, 5)

	fmt.Println(len(hands))
}

func TestChooseStraightFlush(t *testing.T) {
	deck := poker.NewDeck()

	straightFlushHands := []poker.Hand{}

	hands := poker.ChooseN(deck, 5)

	for _, hand := range hands {
		pokerHand := poker.NewPokerHand(hand)

		if pokerHand.Type == poker.StraightFlush {
			fmt.Println(pokerHand.Type)
			straightFlushHands = append(straightFlushHands, pokerHand.Cards)
		}
	}

	fmt.Println(straightFlushHands, len(straightFlushHands))
}
