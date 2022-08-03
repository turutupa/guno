package guno

// UNO helpers
func isValidCard(topDiscardPileCard *Card, card *Card) bool {
	if card.kind == WILD || card.kind == WILD_DRAW_FOUR {
		return true
	}

	if card.kind != NUMBER && topDiscardPileCard.kind == card.kind {
		return true
	}

	// check when top discard pile card is number and might match by color with
	// - skip
	// - reverse
	// - draw two
	if topDiscardPileCard.kind == NUMBER &&
		(card.kind == SKIP || card.kind == REVERSE || card.kind == DRAW_TWO) &&
		topDiscardPileCard.color == card.color {
		return true
	}

	// check when numbers
	if topDiscardPileCard.kind == NUMBER && card.kind == NUMBER {
		if (topDiscardPileCard.number == card.number) || topDiscardPileCard.color == card.color {
			return true
		}
	}

	if (topDiscardPileCard.kind == SKIP || topDiscardPileCard.kind == REVERSE) &&
		topDiscardPileCard.color == card.color {
		return true
	}

	return false
}
