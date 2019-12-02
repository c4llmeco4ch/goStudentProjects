//Package deckofcards provides utilities for creating and working with decks of cards for any number of games
package deckofcards

const (
	//Suits A byte array representing the four suits
	Suits = [4]byte{'H', 'S', 'D', 'C'}

	//MaxValue The max value a playing card (namely, the King) can have
	MaxValue int = 13
)

//Card A generic playing card
type Card struct {
	Suit     byte
	Value    int
	IsFaceUp bool
}

//Deck A generic deck of cards
type Deck struct {
	Cards   []Card
	IsEmpty bool
}

//Flip Flip a given card object c
func (c *Card) Flip() {
	c.IsFaceUp = !c.IsFaceUp
}

//FillDeck Give deck d a new set of 52 playing cards
func (d *Deck) FillDeck() {
	temp := make([]Card, 52)
}
