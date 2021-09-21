package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	//d.shuffle()
	if len(d) != 52 {
		t.Errorf("Error, expected length of deck is 52 but found %v ", len(d))
	} else if d[0] != "Spades of Ace" {
		t.Errorf("First card should be Ace of spades but found %v", d[0])
	} else if d[len(d)-1] != "Clubs of Jack" {
		t.Errorf("Last Card should be Clubs of Jack but found %v", d[len(d)-1])
	}
}
func TestToSaveFileReadFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()

	fmt.Print(d.tosaveFile("_decktesting.txt"))

	loadedDeck := readFromFile("_decktesting.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("The data wasn't inserted properly %v", len(loadedDeck))
	}
	os.Remove("_decktesting.txt")

}
