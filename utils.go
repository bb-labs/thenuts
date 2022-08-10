package poker

func ChooseN(cards []*Card, n int) []Hand {
	choices := []Hand{}
	choice := map[*Card]bool{}

	var helper func(int)

	helper = func(nextIndex int) {
		chosenHand := Hand{}
		if len(choice) == n {
			for card := range choice {
				chosenHand = append(chosenHand, card)
			}
			choices = append(choices, chosenHand)
			return
		}

		for i := nextIndex; i < len(cards); i++ {
			card := cards[i]

			choice[card] = true
			helper(i + 1)
			delete(choice, card)
		}
	}

	helper(0)

	return choices
}

func GetPokerHandType(ranks RankCounter, suits SuitCounter) PokerHandType {
	isFlush := len(suits) == 1
	isStraight := IsStraight(ranks) || IsLowStraight(ranks)

	if isStraight && isFlush {
		return StraightFlush
	}
	if isFlush {
		return Flush
	}
	if isStraight {
		return Straight
	}
	if IsNOfAKind(ranks, 4) {
		return FourOfAKind
	}

	IsPair := IsNOfAKind(ranks, 2)
	IsThreeOfKind := IsNOfAKind(ranks, 3)

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

func IsLowStraight(ranks RankCounter) bool {
	_, hasAce := ranks[Ace]
	_, hasTwo := ranks[Two]
	_, hasThree := ranks[Three]
	_, hasFour := ranks[Four]
	_, hasFive := ranks[Five]

	return hasAce && hasTwo && hasThree && hasFour && hasFive
}

func IsStraight(ranks RankCounter) bool {
	if len(ranks) != 5 {
		return false
	}

	numCards := 0

	for rank := Two; rank <= Ace; rank++ {
		_, hasRank := ranks[rank]

		if hasRank {
			numCards += 1
			if numCards == 5 {
				return true
			}
			continue
		}

		if numCards == 0 {
			continue
		}

		break
	}

	return false
}

func IsNOfAKind(ranks RankCounter, n int) bool {
	for _, count := range ranks {
		if count == n {
			return true
		}
	}

	return false
}

func IsTwoPair(ranks RankCounter) bool {
	numPairs := 0

	for _, count := range ranks {
		if count == 2 {
			numPairs += 1
		}
	}

	return numPairs == 2
}
