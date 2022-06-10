package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
//  which is a slice of strings
type deck []string

func newDeck() deck {
	// an empty deck
	cards := deck{}
	// as slice of strings
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	// as slice of strings
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Sarbaz", "Queen", "King"}
	for _, suit := range cardSuits {
		// _ means no i no j cause we do not use it in the for loop
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

// function is gonna return to types of deck and deck
func deal(d deck, handSize int) (deck, deck) {
	// fmt.Println(d)
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// take a slice of string and join them automatically
	//  we have to import string package
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) error {
	// 0666 anyone can read and write this file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// read file return byte slice and error
	// if nothing goes wrong error value will be nil which is no value
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 - log the error and etirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//ace of spades, two of spades, three of spades, .... is bs
	// string slice is s
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
