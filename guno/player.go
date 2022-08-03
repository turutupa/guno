package guno

import (
	"fmt"
	"strconv"
)

// PLAYER
func (player *Player) init(name string) Player {
	player.name = name
	player.cards = make([]Card, 0)
	player.hasWon = false
	return *player
}

func (player *Player) removeCardAt(i int) Card {
	removedCard := player.cards[i]
	player.cards = append(player.cards[0:i], player.cards[i+1:]...)
	return removedCard
}

func (player *Player) getCardIndex(card *Card) int {
	for i, c := range player.cards {
		if card == &c {
			return i
		}
	}
	panic("No matching card was found at getCardIndex()")
}

func (player *Player) drawCard(card *Card) {
	player.cards = append(player.cards, *card)
}

func (player *Player) hasValidCard(topDiscardPileCard *Card) bool {
	for _, card := range player.cards {
		if isValidCard(topDiscardPileCard, &card) {
			return true
		}
	}

	return false
}

func (player *Player) printCards() {
	for i, card := range player.cards {
		index := i + 1
		prefix := "[" + strconv.Itoa(index) + "] "
		fmt.Println(prefix + getCardText(&card))
	}
}

func getCardText(card *Card) string {
	switch card.kind {
	case NUMBER:
		return "[" + strconv.Itoa(card.number) + "][" + card.color + "]"
	case WILD:
		return "[x] Wild: Change color"
	case WILD_DRAW_FOUR:
		return "[x] Wild + draw four!"
	case SKIP:
		return "[x][" + card.color + "] Skip player"
	case DRAW_TWO:
		return "[x][" + card.color + "] Draw two!"
	case REVERSE:
		return "[x][" + card.color + "] Reverse!"
	default:
		return ""
	}
}
