package main

import "math/rand"

type Deck = []Card

// TODO: would be nice to have a method to sort a deck by
// colors. This way players could have cards sorted out nicely

func buildDeck() Deck {
	cards := make(Deck, 0)

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
			// add drawTwo
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
func shuffleDeck(d *Deck) {
	for i := 1; i < len(*d); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			(*d)[r], (*d)[i] = (*d)[i], (*d)[r]
		}
	}
}

func pop(deck *Deck) *Card {
	d := *deck

	if len(d) == 0 {
		return nil
	}

	card, d := d[len(d)-1], d[:len(d)-1]
	*deck = d
	return &card
}

func peak(deck *Deck) *Card {
	d := *deck

	if len(d) == 0 {
		return nil
	}

	return &d[len(d)-1]
}
