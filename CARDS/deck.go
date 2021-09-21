package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of string
type deck []string

//type byte []byte

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

func (d deck) toString() string {
	return strings.Join(d, "\n")

}
func (d deck) tosaveFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	cards := strings.Split(string(bs), "\n")
	//deck(cards).print()
	return deck(cards)
}
func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(s)
	for index := range d {
		newindex := seed.Intn(len(d))

		d[index], d[newindex] = d[newindex], d[index]
	}
}
func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}
