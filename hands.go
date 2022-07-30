package poker

type PokerHand uint8

const (
	HighCard PokerHand = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func (hand PokerHand) String() string {
	switch hand {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of a Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case StraightFlush:
		return "Straight Flush"
	}

	return ""
}

func Evaluate(hand Hand) PokerHand {
	counts := NewCardCounter(hand)

	if IsStraightFlush(counts) {
		return StraightFlush
	}
	if IsFourOfKind(counts) {
		return FourOfAKind
	}
	if IsFullHouse(counts) {
		return FullHouse
	}
	if IsFlush(counts) {
		return Flush
	}
	if IsStraight(counts) {
		return Straight
	}
	if IsThreeOfKind(counts) {
		return ThreeOfAKind
	}
	if IsTwoPair(counts) {
		return TwoPair
	}
	if IsPair(counts) {
		return OnePair
	}

	return HighCard
}

func IsFlush(counts CardCounter) bool {
	return len(counts.Suits) == 1
}

func IsStraight(counts CardCounter) bool {
	if len(counts.Ranks) != 5 {
		return false
	}

	numCards := 0

	_, hasLowAce := counts.Ranks[Ace]
	if hasLowAce {
		numCards = 1
	}

	for rank := Two; rank <= Ace; rank++ {
		_, hasRank := counts.Ranks[rank]

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

func IsStraightFlush(counts CardCounter) bool {
	return IsFlush(counts) && IsStraight(counts)
}

func IsTwoPair(counts CardCounter) bool {
	pairs := 0

	for _, count := range counts.Ranks {
		if count == 2 {
			pairs += 1
		}
	}

	return pairs == 2
}

func IsRepeat(counts CardCounter, n uint8) bool {
	for _, count := range counts.Ranks {
		if count == n {
			return true
		}
	}

	return false
}

func IsPair(counts CardCounter) bool {
	return IsRepeat(counts, 2)
}

func IsThreeOfKind(counts CardCounter) bool {
	return IsRepeat(counts, 3)
}

func IsFourOfKind(counts CardCounter) bool {
	return IsRepeat(counts, 4)
}

func IsFullHouse(counts CardCounter) bool {
	return IsPair(counts) && IsThreeOfKind(counts)
}
