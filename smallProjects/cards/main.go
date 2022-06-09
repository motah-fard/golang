package main

func main() {
	// 	card := newCard()
	// []string is equal to deck
	cards := deck{newCard(), "ace of diamonds"}
	cards = append(cards, "six six six")
	cards.print()
	// fmt.Println(cards, card)
}

func newCard() string {
	return "yooohooo"
}
