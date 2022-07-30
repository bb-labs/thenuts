package poker

import "fmt"

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

func Compare(handA, handB Hand) Hand {
	if len(handA) != 5 {
		return nil
	}

	if len(handA) != len(handB) {
		return nil
	}

	scoreA := Evaluate(handA)
	scoreB := Evaluate(handB)
	fmt.Println(scoreA)
	fmt.Println(scoreB)

	return nil
}
