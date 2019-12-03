//Package deckofcards provides utilities for creating and working with decks of cards for any number of games
package deckofcards

import (
	"crypto/rand"
	"errors"
)

const (
	shuffleNum int = 7

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
	for pos, s := range Suits {
		for v := 1; v <= MaxValue; v++ {
			temp[(pos*13)+v] = Card{s, v, true}
		}
	}
	d.Cards = temp
}

//Shuffle Given a deck of cards, shuffle it seven times
func (d *Deck) Shuffle() {
	for i := 0; i < shuffleNum; i++ {
		for pos := range d.Cards {
			temp, _ := rand.Int(rand.Reader, len(d))
			d.Cards[pos], d.Cards[temp] = d.Cards[temp], d.Cards[pos]
		}
	}
}

//Deal Deal the "top" card from the deck
func (d *Deck) Deal() (Card, error) {
	if !d.IsEmpty {
		dealt := d.Cards[0]
		if len(d.Cards == 1) {
			d.IsEmpty = true
		} else {
			d.Cards = d.Cards[1:]
		}
		return dealt, nil
	} else {
		return nil, errors.New("Tried to deal from an empty deck")
	}

}
