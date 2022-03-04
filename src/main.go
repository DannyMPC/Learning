package main

func main() {
	// var card string = "Ace of Spades"

	cards := newDeck()
	deal(cards, 5)
	cards.shuffle()
	//	cards.print()

	//cards2 := newDeckFromFile("MyCard")
	//	cards2.print()

}
