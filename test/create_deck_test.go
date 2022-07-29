package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/pokerface"
)

func TestCreateDeck(t *testing.T) {
	deck := pokerface.NewDeck()

	for _, card := range deck.Cards {
		fmt.Println(card)
	}

}
