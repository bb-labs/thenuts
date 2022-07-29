package pokerface

import "fmt"

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Heart
	Club
	Diamond
)

const (
	One Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Suit Suit
	Rank Rank
}

var suitMap = map[Suit]string{
	Spade:   "♠️",
	Heart:   "♥️",
	Club:    "♣️",
	Diamond: "♦️",
}

var rankMap = map[Rank]string{
	One:   "1",
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
	Ace:   "A",
}

func (card Card) String() string {
	return fmt.Sprintf("%v %v", rankMap[card.Rank], suitMap[card.Suit])
}

type Deck struct {
	Cards [52]Card
}

func New() *Deck {
	deck := &Deck{}

	for suit := Suit(0); suit < 4; suit++ {
		for rank := Rank(0); rank < 13; rank++ {
			index := uint8(suit)*13 + uint8(rank)
			deck.Cards[index] = Card{suit, rank}
		}
	}

	return deck
}
