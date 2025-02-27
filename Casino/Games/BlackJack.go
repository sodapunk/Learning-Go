package Games

import (
	"GoLangProjects/Cards"
	"fmt"
	"strings"
)

//https://cli.urfave.org
/*
TODO:
	- add split, double down, surrender
	- add graphical version of the board (ASCII)
	- add dealer AI (std dealer rules)
*/
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
	player, dealer, deck := DealCards(deck)
	fmt.Printf("Player : %s\nDealer: %s\n", player, dealer.DealerString())
}

// DealCards gives both the player and dealer two cards, in order. Returns Player hand, and Dealer hand
func DealCards(deck []Cards.Card) (Hand, Hand, []Cards.Card) {
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
	return player, dealer, deck
}

// Does not allow ACE == 11
func (hand Hand) MinScore() int {
	score := 0
	for _, card := range hand {
		score += min(10, int(card.Rank))
	}
	return score
}
func (hand Hand) isBlackJack() bool {
	if len(hand) < 2 {
		return false
	} else if hand.Score() != 21 {
		return false
	}
	return true
}
func (hand Hand) Score() int {
	minScore := hand.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, card := range hand {
		if card.Rank == Cards.Ace {
			return minScore + 10
		}
	}
	return minScore
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
	player, dealer, deck := DealCards(deck)
	var input string
	var card Cards.Card
	var busted bool = false

	for gameloop {

		fmt.Println("Player: ", player, "\nScore: ", player.Score())
		fmt.Println("Dealer: ", dealer.DealerString())
		if player.Score() > 21 {
			fmt.Println("You Bust!")
			busted = true
			break
		}
		if player.isBlackJack() {
			fmt.Println("Black Jack!")
			break
		}
		fmt.Println("Would you like to: (h)it, (s)tand, s(p)lit, (d)ouble down, su(r)render")
		fmt.Scanf("%s", &input)
		switch input {
		case "h":
			card, deck = Deal(deck)
			player = append(player, card)
		case "s":
			gameloop = false
		}
	}
	playerScore := player.Score()
	dealerScore := dealer.Score()
	fmt.Println("===FINAL HANDS===")
	fmt.Println("Player: ", player, "\n Score : ", playerScore)
	fmt.Println("Dealer: ", dealer, "\n Score : ", dealerScore)
	switch {

	}

	if player.Score() > dealer.Score() {
		if !busted {
			fmt.Println("You Win!")
		}
	} else {
		fmt.Println("You Lose!")
	}
}
