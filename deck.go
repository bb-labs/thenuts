package pokerface

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
