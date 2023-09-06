package player

import (
	"fmt"

	"github.com/gabeleonn/gojack/deck"
)

type Player struct {
	Hand deck.Hand
}

func (p *Player) GetAnotherCard(d *deck.Deck) {
	c := d.Pop()
	(*p).Hand = append((*p).Hand, c)
	fmt.Printf("Got: %v\n", c)
}

func (p *Player) GetInitialHand(d *deck.Deck) {
	(*p).Hand = append((*p).Hand, d.Pop())
	(*p).Hand = append((*p).Hand, d.Pop())
}

func (p Player) GetHandCount() int {
	return p.Hand.GetCount()
}
