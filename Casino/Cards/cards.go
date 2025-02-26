//go:generate stringer -type=Suit,Rank
package Cards

import (
	"fmt"
	"sort"
)

type Suit uint8
type Rank uint8

// sort suit-rank by Bridge rules:
const (
	Club Suit = iota
	Diamond
	Heart
	Spade
	Joker
)
const (
	Ace Rank = iota + 1 // give rank corresponding val
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var suits = [...]Suit{Club, Diamond, Heart, Spade}

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit Suit
	Rank Rank
}

// Generate readable text given a card
func (card Card) String() string {
	if card.Suit == Joker {
		return "Joker"
	} else {
		return fmt.Sprintf("%s of %ss", card.Rank, card.Suit)
	}
}

// absoluteRank lets us compare the relative value based on suit:
func absoluteRank(card Card) int {
	return int(card.Suit)*int(maxRank) + int(card.Rank)
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absoluteRank(cards[i]) < absoluteRank(cards[j])
	}
}

func Sort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}
