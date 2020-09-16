package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	d := newDeck()
	testNewDeck(d, t)
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	testNewDeck(loadedDeck, t)

	os.Remove("_decktesting")

}

func testNewDeck(d deck, t *testing.T) {
	if len(d) == 160 {
		t.Error("Expected deck length to not be 160 but it is")
	}

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0].toString() != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %s", d[0].toString())
	}

	if d[len(d)-1].toString() != "Four of Clubs" {
		t.Errorf("Expected first card to be Four of Clubs but got %s", d[len(d)-1].toString())
	}
}
