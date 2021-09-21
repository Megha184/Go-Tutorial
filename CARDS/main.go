package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard() //figure out the type
	//card:="five"

	// cards := deck{newCard(), "Ace"}
	// fmt.Print(cards) //slice of type string
	// cards = append(cards, "Six")
	// cards.print()
	//fmt.Println(card)
	// for i, c := range cards {
	// 	fmt.Println(i, c)
	// }
	//value := newDeck()
	//value.print()
	//d, value := deal(value, 5)

	//d.print()
	//value.print()

	//value.tosaveFile("cards_left.txt")

	cards := readFromFile("cards_left.txt")
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "five"
// }
