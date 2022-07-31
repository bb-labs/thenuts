package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestCreateDeck(t *testing.T) {
	deck := poker.NewDeck()

	fmt.Println(deck)
}
