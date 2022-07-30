package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestCreateDeck(t *testing.T) {
	deck := poker.NewDeck()

	for i, card := range deck.Cards {
		if i%13 == 0 {
			fmt.Println()
		}
		fmt.Print(card, "  ")
	}
	fmt.Println()
}
