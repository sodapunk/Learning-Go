package Cards

import (
	"reflect"
	"testing"
)

func TestCard_String(t *testing.T) {
	card := Card{Suit: Joker}
	if card.String() != "Joker" {
		t.Error("Expected Joker, got ", card.String())
	}
	card = Card{Rank: King, Suit: Heart}
	if card.String() != "King of Hearts" {
		t.Error("Expected King of Hearts, got ", card.String())
	}
	card = Card{Rank: Queen, Suit: Club}
	if card.String() != "Queen of Clubs" {
		t.Error("Expected Queen of Clubs, got ", card.String())
	}
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 52 {
		t.Error("Expected 52 cards, got ", len(deck))
	}
	// Add test for ordering?
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	shuffledDeck := Shuffle(NewDeck()) // NOTE: shuffle happens in-place?
	if reflect.DeepEqual(deck, shuffledDeck) {
		t.Error("Deck is not shuffled", deck, shuffledDeck)
	}
}
