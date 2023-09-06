package main

import (
	"fmt"

	"github.com/gabeleonn/gojack/deck"
	"github.com/gabeleonn/gojack/player"
)

func main() {
	dk := deck.Deck{}
	dk.GenerateCards()
	dk.Shuffle()

	p := player.Player{}
	p.GetInitialHand(&dk)

	d := player.Player{}
	d.GetInitialHand(&dk)

	pp := true
	for pp {
		printScoreAndHand(p)
		rpi := getPlayerInput()
		if rpi != "y" && rpi != "yes" {
			pp = false
		} else {
			p.GetAnotherCard(&dk)
		}

		if p.Hand.Busted() {
			fmt.Println("You lost!")
			fmt.Printf("Your hand score is: %v\n", p.GetHandCount())
			pp = false
		}
	}

	if !p.Hand.Busted() {
		for d.Hand.GetCount() < 17 {
			d.GetAnotherCard(&dk)
		}

		if d.Hand.Busted() {
			fmt.Println("You won!")
			fmt.Printf("Your hand score is: %v\n", p.GetHandCount())
			fmt.Printf("Dealer hand score is: %v\n", d.GetHandCount())
		} else if d.Hand.Blackjack() && p.Hand.Blackjack() || d.Hand.GetCount() == p.Hand.GetCount() {
			fmt.Println("It is a draw!")
			fmt.Printf("Your hand score is: %v\n", p.GetHandCount())
			fmt.Printf("Dealer hand score is: %v\n", d.GetHandCount())
		} else if d.Hand.GetCount() > p.Hand.GetCount() {
			fmt.Println("You lost!")
			fmt.Printf("Your hand score is: %v\n", p.GetHandCount())
			fmt.Printf("Dealer hand score is: %v\n", d.GetHandCount())
		} else {
			fmt.Println("You won!")
			fmt.Printf("Your hand score is: %v\n", p.GetHandCount())
			fmt.Printf("Dealer hand score is: %v\n", d.GetHandCount())
		}
	}
}

func getPlayerInput() string {
	ri := ""
	fmt.Println("Would you like another card? (y)es / n(o)")
	fmt.Scanf("%v", &ri)
	return ri
}

func printScoreAndHand(p player.Player) {
	fmt.Printf("Your hand score is: %v\n", p.GetHandCount())
	fmt.Printf("Your Hand: %v\n", p.Hand)
}
