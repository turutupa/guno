package guno

import "math/rand"

type Deck []Card

func Build() Deck {
	cards := Deck{}

	// add numbered cards
	for _, c := range COLORS {
		// add numbers 0, 1, 1, 2, 2, 3, 3, 4, 4 ...
		for i := 0; i < 10; i++ {
			cardOne := Card{kind: NUMBER, number: i, color: c}

			cards = append(cards, cardOne)
			if i > 0 {
				cardTwo := Card{kind: NUMBER, number: i, color: c}
				cards = append(cards, cardTwo)
			}
		}
	}

	// add special cards
	for _, c := range COLORS {
		for _, t := range SPECIAL_CARDS {
			card := Card{kind: t}

			if t != WILD && t != WILD_DRAW_FOUR {
				card.color = c
			}

			cards = append(cards, card)
			if (t == WILD_DRAW_FOUR) || (t == WILD) {
				continue
			}

			cards = append(cards, Card{kind: t, color: c})
		}
	}

	return cards
}

// I'm not sure this is working the expected way
func (deck *Deck) shuffle() {
	d := *deck
	for i := 1; i < len(d); i++ {
		rand.Shuffle(len(d), func(i, j int) {
			d[i], d[j] = d[j], d[i]
		})
	}

	*deck = d
}

func (deck *Deck) pop() *Card {
	d := *deck

	if len(d) == 0 {
		return nil
	}

	card, d := d[len(d)-1], d[:len(d)-1]
	*deck = d
	return &card
}

func (deck *Deck) push(card *Card) {
	d := *deck
	d = append(d, *card)
	*deck = d
}

func shift(deck *Deck) *Card {
	d := *deck

	if len(d) == 0 {
		return nil
	}

	card, d := d[0], d[1:]
	return &card
}

func (deck *Deck) peak() *Card {
	d := *deck

	if len(d) == 0 {
		return nil
	}

	return &d[len(d)-1]
}

func (deck *Deck) sortByColor() {
	copy := make(Deck, 0)

	for _, color := range COLORS {
		for _, card := range *deck {
			if card.color == color {
				copy = append(copy, card)
			}
		}
	}

	for _, card := range *deck {
		if card.color == "" {
			copy = append(copy, card)
		}
	}

	*deck = copy
}
