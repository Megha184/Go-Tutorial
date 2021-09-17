package main

import "fmt"

// Create a new type of deck which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King",
		"Queen", "Jack"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func toString(d deck) string {
	card := ""
	for _, c := range d {
		card += "\n" + c
	}
	return card

}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}
