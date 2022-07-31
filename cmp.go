package poker

import (
	"math"
	"sort"
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
)

func GetPokerHandType(ranks RankCounter, suits SuitCounter) PokerHandType {
	isFlush := len(suits) == 1
	isStraight := IsStraight(ranks)
	IsPair := IsNOfAKind(ranks, 2)
	IsThreeOfKind := IsNOfAKind(ranks, 3)

	if isStraight && isFlush {
		return StraightFlush
	}
	if isStraight {
		return Straight
	}
	if isFlush {
		return Flush
	}
	if IsNOfAKind(ranks, 4) {
		return FourOfAKind
	}
	if IsThreeOfKind && IsPair {
		return FullHouse
	}
	if IsThreeOfKind {
		return ThreeOfAKind
	}
	if IsTwoPair(ranks) {
		return TwoPair
	}
	if IsPair {
		return Pair
	}

	return HighCard
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

	return pokerHand
}

func (hand PokerHand) Sorted() {
	sort.Slice(hand.Cards, func(i, j int) bool {
		// If the two cards have different rank counts, the one with the higher count is greater
		if hand.RankCounts[hand.Cards[i].Rank] != hand.RankCounts[hand.Cards[j].Rank] {
			return hand.RankCounts[hand.Cards[i].Rank] < hand.RankCounts[hand.Cards[j].Rank]
		}

		// Same rank count, then lower rank means what you think it means
		return hand.Cards[i].Rank < hand.Cards[j].Rank
	})
}

func (hand PokerHand) Evaluate() int {
	value := 0

	for i, card := range hand.Cards {
		value += int(math.Pow(13, float64(i)) * float64(card.Rank))
	}

	return value * int(math.Pow(13, float64(hand.Type)))
}
