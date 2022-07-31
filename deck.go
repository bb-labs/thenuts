package poker

import (
	"fmt"
	"strings"
)

const (
	NumSuits     = 4
	NumRanks     = 13
	MaxRankValue = 14
)

type Suit int
type Rank int

const (
	Spade Suit = iota
	Heart
	Club
	Diamond
)

func (suit Suit) String() string {
	switch suit {
	case Spade:
		return "♠️"
	case Heart:
		return "♥️"
	case Club:
		return "♣️"
	case Diamond:
		return "♦️"
	default:
		return ""
	}
}

const (
	LowAce Rank = iota + 1
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

func (rank Rank) String() string {
	switch rank {
	case LowAce:
		return "A"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return ""
	}
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	return fmt.Sprint(card.Rank, card.Suit)
}

type Hand []Card

func (hand Hand) String() string {
	cards := []string{}
	for _, card := range hand {
		cards = append(cards, card.String())
	}
	return fmt.Sprint(strings.Join(cards, " , "))
}

type Deck [52]Card

func NewDeck() Deck {
	deck := Deck{}

	for suit := Suit(0); suit < NumSuits; suit++ {
		for rank := Rank(0); rank < NumRanks; rank++ {
			index := int(suit)*NumRanks + int(rank)
			deck[index] = Card{rank, suit}
		}
	}

	return deck
}
