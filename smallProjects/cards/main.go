package main

func main() {
	// 	card := newCard()
	// []string is equal to deck
	// cards := deck{newCard(), "ace of diamonds"}
	// cards = append(cards, "six six six")
	// cards.print()
	// fmt.Println(cards, card)
	// cards := newDeck()
	// cards.print()
	// hand, reaminingCards := deal(cards, 5)
	// hand.print()
	// reaminingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "yooohooo"
// }
