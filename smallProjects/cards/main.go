package main

import "fmt"

func main() {
	card := newCard()
	cards := []string{newCard(), "ace of diamonds"}
	cards = append(cards, "six six six")
	for i, myCard := range cards {
		fmt.Println(i, myCard)
	}
	fmt.Println(cards, card)
}

func newCard() string {
	return "yooohooo"
}
