package poker

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

func IsStraight(counts CardCounter) bool {
	if len(counts.Ranks) == 1 {
		return true
	}

	numCards := 0

	// Check low ace
	_, hasLowAce := counts.Ranks[Ace]

	if hasLowAce {
		numCards += 1
	}

	for rank := Two; rank <= Ace; rank++ {
		_, hasRank := counts.Ranks[rank]

		// We have a straight
		if hasRank && (numCards+1) == len(counts.Ranks) {
			return true
		}

		// Keep going, we haven't missed yet
		if hasRank {
			numCards += 1
			continue
		}

		// We missed, but we don't have any cards yet
		if numCards == 0 {
			continue
		}

		// We missed, but we have a low ace. Maybe the straight is royal
		if hasLowAce && rank == Two {
			numCards = 0
		}

		return false
	}

	return true
}

type CardCounter struct {
	Ranks map[Rank]uint8
	Suits map[Suit]uint8
}

func NewCardCounter(hand Hand) CardCounter {
	counter := CardCounter{
		map[Rank]uint8{},
		map[Suit]uint8{},
	}

	for _, card := range hand {
		counter.Ranks[card.Rank] += 1
		counter.Suits[card.Suit] += 1
	}

	return counter
}

func Compare(a, b Hand) Hand {
	countsA := NewCardCounter(a)
	countsB := NewCardCounter(b)

	fmt.Println(countsA)
	fmt.Println(countsB)

	return nil
}
