package main

import (
	"os"
	"testing"
)

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
	// first card is Ace of spades
	c := "Spades of Ace"
	if d[0] != c {
		t.Errorf("Expected first card of Spades of Ace but got %v", d[0])
	}
	// last card is four of clubs
	if d[len(d)-1] != "Clubs of King" {
		t.Errorf("Expected last card of Four of Spades but got %v", d[len(d)-1])
	}
}

// test save to deck and new deck from file
//  create deck
//  save to file
//  file created
// attempt to load load file
// crash

// delete any files with the name decktesting
// create deck
// save to file
// load from file
// assert deck length
// delete any files in current working directory with name

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	x := 52
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != x {
		t.Errorf("Expected deck length of %x, but got %v", x, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
