package test

import (
	"fmt"
	"testing"

	"github.com/bb-labs/poker"
)

func TestFullHouse(t *testing.T) {
	tests := []struct {
		Hand    poker.Hand
		IsFlush bool
	}{
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Ace, Suit: poker.Diamond},
				poker.Card{Rank: poker.King, Suit: poker.Heart},
				poker.Card{Rank: poker.King, Suit: poker.Spade},
				poker.Card{Rank: poker.King, Suit: poker.Diamond},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Heart},
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Ace, Suit: poker.Club},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
	}

	for _, test := range tests {
		fmt.Println(test.Hand, poker.IsFullHouse(poker.NewCardCounter(test.Hand)))
	}
}
func TestTwoPair(t *testing.T) {
	tests := []struct {
		Hand    poker.Hand
		IsFlush bool
	}{
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Ace, Suit: poker.Diamond},
				poker.Card{Rank: poker.King, Suit: poker.Heart},
				poker.Card{Rank: poker.King, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Heart},
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
	}

	for _, test := range tests {
		fmt.Println(test.Hand, poker.IsTwoPair(poker.NewCardCounter(test.Hand)))
	}
}

func Test3OfKind(t *testing.T) {
	tests := []struct {
		Hand    poker.Hand
		IsFlush bool
	}{
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Ace, Suit: poker.Diamond},
				poker.Card{Rank: poker.Ace, Suit: poker.Heart},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Heart},
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
	}

	for _, test := range tests {
		fmt.Println(test.Hand, poker.IsThreeOfKind(poker.NewCardCounter(test.Hand)))
	}
}
func TestFlush(t *testing.T) {
	tests := []struct {
		Hand    poker.Hand
		IsFlush bool
	}{
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Heart},
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
	}

	for _, test := range tests {
		fmt.Println(test.Hand, poker.IsFlush(poker.NewCardCounter(test.Hand)))
	}
}

func TestStraight(t *testing.T) {
	tests := []struct {
		Hand       poker.Hand
		IsStraight bool
	}{
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.King, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Six, Suit: poker.Spade},
				poker.Card{Rank: poker.Six, Suit: poker.Heart},
				poker.Card{Rank: poker.King, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Six, Suit: poker.Heart},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			false,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Three, Suit: poker.Spade},
				poker.Card{Rank: poker.Four, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Two, Suit: poker.Spade},
				poker.Card{Rank: poker.Three, Suit: poker.Spade},
				poker.Card{Rank: poker.Four, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
				poker.Card{Rank: poker.Six, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Three, Suit: poker.Spade},
				poker.Card{Rank: poker.Four, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
				poker.Card{Rank: poker.Six, Suit: poker.Spade},
				poker.Card{Rank: poker.Seven, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Four, Suit: poker.Spade},
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
				poker.Card{Rank: poker.Six, Suit: poker.Spade},
				poker.Card{Rank: poker.Seven, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Five, Suit: poker.Spade},
				poker.Card{Rank: poker.Six, Suit: poker.Spade},
				poker.Card{Rank: poker.Seven, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Nine, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Six, Suit: poker.Spade},
				poker.Card{Rank: poker.Seven, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Nine, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Seven, Suit: poker.Spade},
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Nine, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Jack, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Eight, Suit: poker.Spade},
				poker.Card{Rank: poker.Nine, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Jack, Suit: poker.Spade},
				poker.Card{Rank: poker.Queen, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Nine, Suit: poker.Spade},
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Jack, Suit: poker.Spade},
				poker.Card{Rank: poker.Queen, Suit: poker.Spade},
				poker.Card{Rank: poker.King, Suit: poker.Spade},
			},
			true,
		},
		{
			poker.Hand{
				poker.Card{Rank: poker.Ten, Suit: poker.Spade},
				poker.Card{Rank: poker.Jack, Suit: poker.Spade},
				poker.Card{Rank: poker.Queen, Suit: poker.Spade},
				poker.Card{Rank: poker.King, Suit: poker.Spade},
				poker.Card{Rank: poker.Ace, Suit: poker.Spade},
			},
			true,
		},
	}

	for _, test := range tests {
		fmt.Println(test.Hand, poker.IsStraight(poker.NewCardCounter(test.Hand)))
	}

}
