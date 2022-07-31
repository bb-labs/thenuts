package poker

func IsStraight(ranks RankCounter) bool {
	if len(ranks) != 5 {
		return false
	}

	numCards := 0

	_, hasLowAce := ranks[Ace]
	if hasLowAce {
		numCards = 1
	}

	for rank := Two; rank <= Ace; rank++ {
		_, hasRank := ranks[rank]

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

// // func IsStraightFlush(counts CardCounter) bool {
// // 	return IsFlush(counts) && IsStraight(counts)
// // }

// func CheckTwoPair(counts CardCounter) PokerHandScore {
// 	var lowPair Rank = -1
// 	var highPair Rank = -1
// 	var tieBreak Rank = -1

// 	for rank, count := range ranks {
// 		if count == 1 {
// 			tieBreak = rank
// 		} else if count == 2 {
// 			if lowPair == -1 {
// 				lowPair = rank
// 			} else {
// 				lowPair = Rank(math.Min(float64(lowPair), float64(rank)))
// 				highPair = Rank(math.Max(float64(lowPair), float64(rank)))
// 			}
// 		}
// 	}

// 	if lowPair == -1 || highPair == -1 {
// 		return nil
// 	}

// 	return PokerHandScore{TwoPair, highPair, lowPair, tieBreak}
// }

// func CheckFullHouse(counts CardCounter) PokerHandScore {
// 	var pair Rank = -1
// 	var threeOfAKind Rank = -1

// 	for rank, count := range ranks {
// 		if count == 3 {
// 			threeOfAKind = rank
// 		} else if count == 2 {
// 			pair = rank
// 		}
// 	}

// 	if pair == -1 || threeOfAKind == -1 {
// 		return nil
// 	}

// 	return PokerHandScore{Pair, threeOfAKind, pair}

// }

// func CheckPair(counts CardCounter) PokerHandScore {
// 	return CheckGroup(counts, 2)
// }

// func CheckThreeOfKind(counts CardCounter) PokerHandScore {
// 	return CheckGroup(counts, 3)
// }

// func CheckFourOfKind(counts CardCounter) PokerHandScore {
// 	return CheckGroup(counts, 4)
// }

// func CheckGroup(counts CardCounter, n int) PokerHandScore {
// 	var groupRank Rank = -1
// 	tiebreak := PokerHandScore{}

// 	for rank, count := range ranks {
// 		if count == n {
// 			groupRank = rank
// 		} else {
// 			tiebreak = append(tiebreak, rank)
// 		}
// 	}

// 	if groupRank == -1 {
// 		return nil
// 	}

// 	tiebreak.Sort()

// 	return append(PokerHandScore{Pair, groupRank}, tiebreak...)
// }

// func CheckHighCard(counts CardCounter) PokerHandScore {
// 	tiebreak := PokerHandScore{}

// 	for rank, _ := range ranks {
// 		tiebreak = append(tiebreak, rank)
// 	}

// 	tiebreak.Sort()

// 	return append(PokerHandScore{HighCard}, tiebreak...)
// }
