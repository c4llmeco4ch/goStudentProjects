//Package deckofcards provides utilities for creating and working with decks of cards for any number of games
package deckofcards

import "fmt"

//Card A generic playing card
type Card struct {
	Suit     byte
	Value    int
	IsFaceUp bool
}

//Flip Flip a given card object c
func (c *Card) Flip() {
	c.IsFaceUp = !c.IsFaceUp
}

//String Print the given card
func (c Card) String() string {
	if IsFaceUp {
		return fmt.Sprintf("%d%s", c.Value, c.Suit)
	}
	return fmt.Sprintf("??")
}
