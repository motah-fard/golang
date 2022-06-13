package main

import "testing"

// make sure the deck is created with X amount of cards
// every function starts with capital T
// t is the test value() test handler
// T is the type of value that being passed into the function
func TestNewDeck(t *testing.T) {
	d := newDeck()
	x := 52
	if len(d) != x {
		t.Errorf("Expected deck length of %x, but got %d", x, len(d))
	}
	c := "Spades of Ace"
	if d[0] != c {
		t.Errorf("Expected first card of Spades of Ace but got %v", d[0])
	}
	if d[len(d)-1] != "Clubs of King" {
		t.Errorf("Expected last card of Four of Spades but got %v", d[len(d)-1])
	}
}

// first card is Ace of spades
// last card is four of clubs
// test save to deck and new deck from file
