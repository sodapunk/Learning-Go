package Cards

import "math/rand"

// NewDeck generates standard dec
// TODO : add the option to play with multiple decks
func NewDeck() []Card {
	var deck []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// Shuffle we use Fisher-Yates shuffle
func Shuffle(deck []Card) []Card {
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i] // swap positions
	}
	return deck
}
