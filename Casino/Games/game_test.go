package Games

import (
	"GoLangProjects/Cards"
	"testing"
)

func TestScore_Ace(t *testing.T) {
	card1 := Cards.Card{Rank: Cards.Ace, Suit: Cards.Club}
	card2 := Cards.Card{Rank: Cards.Eight, Suit: Cards.Club}
	hand := Hand{card1, card2}
	score := hand.Score()
	if score != 19 {
		t.Error("Expected 19 but got ", score)
	}
}
