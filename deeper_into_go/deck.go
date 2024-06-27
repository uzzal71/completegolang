package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{"Ace of DiamondS", "Six of Spades"}
	cards = append(cards, "Six of Spades")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
