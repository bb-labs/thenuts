package poker

type Game struct {
	Deck    Deck
	Players []PokerHand
}

func NewGame(numPlayers int) Game {
	return Game{
		Deck:    NewDeck(),
		Players: make([]PokerHand, numPlayers),
	}
}

func (game *Game) Burn() {
	// _, newDeck := game.Deck.Draw()
	// game.Deck = newDeck
}

func Deal() {

}
