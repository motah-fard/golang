package main

import "fmt"

// create a new type of deck
//  which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	// as slice of strings
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	// as slice of strings
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Sarbaz", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// reciever on a function
// d is like this or self in python or JS/ but we call it the actual thing that it is
// d is the abbreviation of deck cause cards are part of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
