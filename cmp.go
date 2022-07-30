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

func IsFlush(counts CardCounter) bool {
	return len(counts.Suits) == 1
}

func IsStraight(counts CardCounter) bool {
	numCards := 0

	_, hasLowAce := counts.Ranks[Ace]
	if hasLowAce {
		numCards = 1
	}

	for rank := Two; rank <= Ace; rank++ {
		count, hasRank := counts.Ranks[rank]

		// We can't have a straight if we're doubled up
		if count > 1 {
			return false
		}

		// We have a straight
		if hasRank && (numCards+1) == 5 {
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
			continue
		}

		return false
	}

	return false
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
