package deck

import (
	"fmt"

	"github.com/gabeleonn/gojack/utils"
)

type Card struct {
	Label string
	Value int
}

type Hand []Card

func (h Hand) GetCount() int {
	count := 0
	for _, c := range h {
		count += c.Value
	}

	ac := h.getAcesCount()
acecheckloop:
	for ac > 0 {
		if count > 21 {
			count -= 10
			ac--
		} else {
			break acecheckloop
		}
	}

	return count
}

func (h Hand) getAcesCount() int {
	count := 0
	for _, c := range h {
		if c.Value == 1 {
			count += c.Value
		}
	}
	return count
}

func (h Hand) Busted() bool {
	return h.GetCount() > 21
}

func (h Hand) Blackjack() bool {
	return h.GetCount() == 21
}

type Deck []Card

func (d *Deck) GenerateCards() {
	for v, c := range cards {
		for _, s := range suits {
			fv := v + 1
			if fv > 10 {
				fv = 10
			} else if fv == 1 {
				fv = 11
			}
			*d = append(*d, Card{Label: fmt.Sprintf("%v of %v", c, s), Value: fv})
		}
	}
}

func (d *Deck) Shuffle() {
	r := utils.GetSpecialRand()
	r.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

func (d *Deck) Pop() Card {
	popped, remaning := (*d)[0], (*d)[1:]
	*d = remaning
	return popped
}
