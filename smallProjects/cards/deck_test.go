package main

import "testing"

// make sure the deck is created with X amount of cards
// every function starts with capital T
// t is the test value() test handler
// T is the type of value that being passed into the function
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}
}

// first card is Ace of spades
// last card is four of clubs
// test save to deck and new deck from file
