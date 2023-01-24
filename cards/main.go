package main

import "fmt"

func main() {
	// Arrays are always fixed length, a Slice can grow or shrink.
	// Every element in the slice must be of the same data type
	// the data type is inferred when using :=
	cards := []string{newCard(), newCard()}
	// append will return a new slice, it does not alter the original slice
	cards = append(cards, "Old Card")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "New Card"
}
