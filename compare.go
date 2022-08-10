package poker

import (
	"math"
	"sort"
)

const (
	RankNumberBase = 15
	SmallestHand   = 15 * 15 * 15 * 15
)

type PokerHandType int

const (
	HighCard PokerHandType = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	UnknownHand
)

func (pokerHandType PokerHandType) String() string {
	switch pokerHandType {
	case HighCard:
		return "HighCard"
	case Pair:
		return "Pair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case StraightFlush:
		return "StraightFlush"
	}
	return ""
}

type RankCounter map[Rank]int
type SuitCounter map[Suit]int

type PokerHand struct {
	Cards      Hand
	Type       PokerHandType
	RankCounts RankCounter
	SuitCounts SuitCounter
}

func NewPokerHand(hand Hand) PokerHand {
	rankCounts := RankCounter{}
	suitCounts := SuitCounter{}

	for _, card := range hand {
		rankCounts[card.Rank] += 1
		suitCounts[card.Suit] += 1
	}

	pokerHand := PokerHand{
		hand,
		GetPokerHandType(rankCounts, suitCounts),
		rankCounts,
		suitCounts,
	}

	pokerHand.Sorted()

	// Ace is counted as 1 in low straight
	if IsLowStraight(pokerHand.RankCounts) {
		pokerHand.Cards = Hand{
			pokerHand.Cards[4],
			pokerHand.Cards[0],
			pokerHand.Cards[1],
			pokerHand.Cards[2],
			pokerHand.Cards[3],
		}
	}

	return pokerHand
}

func (hand PokerHand) Sorted() {
	sort.Slice(hand.Cards, func(i, j int) bool {
		handIRank := hand.Cards[i].Rank
		handJRank := hand.Cards[j].Rank

		// If the two cards have different rank counts, the one with the higher count is greater
		if hand.RankCounts[handIRank] != hand.RankCounts[handJRank] {
			return hand.RankCounts[handIRank] < hand.RankCounts[handJRank]
		}

		// Same rank count, then lower rank means what you think it means
		return handIRank < handJRank
	})
}

func (hand PokerHand) Evaluate() int {
	value := 0

	for i, card := range hand.Cards {
		base := math.Pow(RankNumberBase, float64(i))
		digit := float64(card.Rank)
		value += int(base * digit)
	}

	return value * int(math.Pow(RankNumberBase, float64(hand.Type)))
}
