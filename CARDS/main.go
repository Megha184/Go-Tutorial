package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard() //figure out the type
	//card:="five"
	cards := []string{newCard(), "Ace"}
	fmt.Print(cards) //slice of type string
	cards = append(cards, "Six")
	fmt.Println(card)
	for i, c := range cards {
		fmt.Println(i, c)
	}

}
func newCard() string {
	return "five"
}
