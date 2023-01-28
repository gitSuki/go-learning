package main

import (
	"fmt"
)

func main() {
	// Arrays are always fixed length, a Slice can grow or shrink.
	// Every element in the slice must be of the same data type
	// the data type is inferred when using :=
	cards := newDeck()
	fmt.Println(cards.toString())
}
