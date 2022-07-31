package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestThings(t *testing.T) {
	scores := poker.PokerHandScore{poker.Ace, poker.Eight, poker.Queen, poker.King}

	scores.Sort()

	fmt.Println(scores)
}
