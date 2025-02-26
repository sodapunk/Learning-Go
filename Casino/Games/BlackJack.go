package Games

import (
	"GoLangProjects/Cards"
	"fmt"
	"strings"
)

//https://cli.urfave.org

type Hand []Cards.Card

// Define a nicer print for readability
func (hand Hand) String() string {
	cards := make([]string, len(hand))
	for i := range hand {
		cards[i] = hand[i].String()
	}
	return strings.Join(cards, ", ")
}

func CheckCards() {
	deck := Cards.Shuffle(Cards.NewDeck())
	var card Cards.Card
	for i := 0; i < 10; i++ {
		card, deck = deck[0], deck[1:]
		fmt.Println(card)
	}
	var hand Hand = []Cards.Card{deck[0], deck[1], deck[2]}
	fmt.Println(hand)
	fmt.Println("DEALING CARDS:")
	player, dealer := DealCards(deck)
	fmt.Printf("Player : %s\nDealer: %s\n", player, dealer.DealerString())
}

// DealCards gives both the player and dealer two cards, in order. Returns Player hand, and Dealer hand
func DealCards(deck []Cards.Card) (Hand, Hand) {
	var player, dealer Hand
	var card Cards.Card
	for i := 0; i < 2; i++ {
		// Use pointers to directly update the hands
		// This logic allows us to update both hands using a single loop
		for _, hand := range []*Hand{&player, &dealer} {
			card, deck = Deal(deck)
			*hand = append(*hand, card)
		}
	}
	return player, dealer
}

// Deal deals one card from the deck, and returns the card, and the rest of the deck
func Deal(deck []Cards.Card) (Cards.Card, []Cards.Card) {
	return deck[0], deck[1:]
}

func (hand Hand) DealerString() string {
	return hand[0].String() + ", *HIDDEN*"
}

// PlayBlackJack defines the basic gameloop
func PlayBlackJack() {
	gameloop := true
	deck := Cards.Shuffle(Cards.NewDeck())
	player, dealer := DealCards(deck)
	var input string
	var card Cards.Card

	for gameloop {
		fmt.Println("Player: ", player)
		fmt.Println("Dealer: ", dealer.DealerString())
		fmt.Println("Would you like to: (h)it, (s)tand,")
		fmt.Scanf("%s", &input)
		switch input {
		case "h":
			card, deck = Deal(deck)
			player = append(player, card)
		case "s":
			gameloop = false
		}
	}
	fmt.Println("===FINAL HANDS===")
	fmt.Println("Player: ", player)
	fmt.Println("Dealer: ", dealer)

}
