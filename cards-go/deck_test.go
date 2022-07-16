package main

import (
	"os"
	"testing"
)

func TestNewCard(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length to be 16 but instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be \"Ace of Spades\" but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be \"Four of Clubs\" but got %v", d[len(d)-1])
	}
}

func TestSavetoDeckAndNewDeckFromFile (t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length to be 16 but instead got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}