package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()

	// remainingCards.print()
	// fmt.Println(cards.toString())

}
