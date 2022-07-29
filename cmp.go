package pokerface

import "fmt"

// ------- Poker hands ---------

// High Card
// One Pair
// Two Pair
// Three of a Kind
// Straight
// Flush
// Full House
// Four of a Kind
// Straight Flush

type RankCounter map[Rank]uint8
type SuitCounter map[Suit]uint8

func Compare(a []Card, b []Card) []Card {
	rankCountsA := RankCounter{}
	suitCountsA := SuitCounter{}
	rankCountsB := RankCounter{}
	suitCountsB := SuitCounter{}

	for _, card := range a {
		rankCountsA[card.Rank] += 1
		suitCountsA[card.Suit] += 1
	}

	for _, card := range b {
		rankCountsB[card.Rank] += 1
		suitCountsB[card.Suit] += 1
	}

	fmt.Println(rankCountsA, suitCountsA)
	fmt.Println(rankCountsB, suitCountsB)

	return nil
}
