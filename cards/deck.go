package main

import (
	"fmt"
	"strings"
)

// create a new type of deck
// a slice of strings.
// this syntax is like saying the deck extends or inherits all the behavior of a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen"}

	// we can use underscores in cases like this where we want to use a positional argument but don't actually need to access it's value
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append will return a new slice, it does not alter the original slice
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// the (d deck) is a reciever.
// the d is the actual copy of the deck we're working with within the function code itself. that makes it very similar to this / self keywords, but in go we always want to refer to the actual thing itself rather an abstract keyword. by convention we normally use a one or two letter abbreviation of the actual data type
// the deck says any variables of type deck have access to print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}
