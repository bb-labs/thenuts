package test

import (
	"testing"
)

func TestHandPercentile(t *testing.T) {
	// deck := poker.NewDeck()

	// hand := poker.Hand{
	// 	poker.Card{poker.King, poker.Club},
	// 	poker.Card{poker.King, poker.Heart},
	// }

	// handOpp := poker.Hand{
	// 	poker.Card{poker.Ace, poker.Club},
	// 	poker.Card{poker.Ace, poker.Diamond},
	// }

	// deck = deck.Remove(hand[0])
	// deck = deck.Remove(hand[1])
	// deck = deck.Remove(handOpp[0])
	// deck = deck.Remove(handOpp[1])

	// totalHands := 0
	// winningHands := 0
	// winningHandsOpp := 0

	// // Burn a card
	// for _, burn1 := range deck {
	// 	deck = deck.Remove(burn1)
	// 	// Check the flops
	// 	allFlops := poker.ChooseN(deck, 3, nil)
	// 	for _, flop := range allFlops {
	// 		deck = deck.Remove(flop[0])
	// 		deck = deck.Remove(flop[1])
	// 		deck = deck.Remove(flop[2])
	// 		// // Burn another card
	// 		// for _, burn2 := range deck {
	// 		// 	deck = deck.Remove(burn2)
	// 		// 	// Check the turns
	// 		// 	allTurns := poker.ChooseN(deck, 1, nil)
	// 		// 	for j, turn := range allTurns {
	// 		// 		deck = deck.Remove(turn[0])
	// 		// 		// Burn another card
	// 		// 		for _, burn3 := range deck {
	// 		// 			deck = deck.Remove(burn3)
	// 		// 			// Check the river
	// 		// 			allRivers := poker.ChooseN(deck, 1, nil)
	// 		// 			for k, river := range allRivers {
	// 		// 				iter := i*len(allTurns)*len(allRivers) + j*len(allRivers) + k
	// 		// 				if iter%10000 == 0 {
	// 		// 					fmt.Println(iter, "/", len(allFlops)*len(allTurns)*len(allRivers))
	// 		// 				}

	// 		possibleHand := append(hand, flop...)
	// 		// possibleHand = append(possibleHand, turn...)
	// 		// possibleHand = append(possibleHand, river...)
	// 		// possibleHands := poker.ChooseN(poker.Deck(possibleHand), 5, nil)

	// 		possibleHandOpp := append(handOpp, flop...)
	// 		// possibleHandOpp = append(possibleHandOpp, turn...)
	// 		// possibleHandOpp = append(possibleHandOpp, river...)
	// 		// possibleOppHands := poker.ChooseN(poker.Deck(possibleHandOpp), 5, nil)

	// 		// sort.Slice(possibleHands, func(i, j int) bool {
	// 		// 	return poker.NewPokerHand(possibleHands[i]).Evaluate() < poker.NewPokerHand(possibleHands[j]).Evaluate()
	// 		// })
	// 		// sort.Slice(possibleOppHands, func(i, j int) bool {
	// 		// 	return poker.NewPokerHand(possibleOppHands[i]).Evaluate() < poker.NewPokerHand(possibleOppHands[j]).Evaluate()
	// 		// })

	// 		bestPossibleHand := poker.NewPokerHand(possibleHand)
	// 		bestPossibleHandOpp := poker.NewPokerHand(possibleHandOpp)

	// 		if bestPossibleHand.Evaluate() == bestPossibleHandOpp.Evaluate() {
	// 			continue
	// 		}

	// 		totalHands++
	// 		if bestPossibleHand.Evaluate() > bestPossibleHandOpp.Evaluate() {
	// 			winningHands++
	// 		} else {
	// 			winningHandsOpp++
	// 		}

	// 		// 			}
	// 		// 			deck = append(deck, burn3)
	// 		// 		}
	// 		// 		deck = append(deck, turn...)
	// 		// 	}
	// 		// 	deck = append(deck, burn2)
	// 		// }
	// 		deck = append(deck, flop...)
	// 	}
	// 	deck = append(deck, burn1)
	// }

	// fmt.Println(hand, float32(winningHands)/float32(totalHands))
	// fmt.Println(handOpp, float32(winningHandsOpp)/float32(totalHands))

}
