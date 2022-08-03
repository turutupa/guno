package guno

import (
	"log"
	"math"
)

// GAME
func (game *Game) Init(deck *Deck) {
	game.players = make([]Player, 0) // Init players arr
	game.turn = 0                    // Init turn
	game.discardPile = make(Deck, 0) // Init discardPile
	game.drawPile = *deck            // Init drawPile
	game.direction = DEFAULT_DIRECTION
	game.useAnyCard = false
}

func (game *Game) start() {
	game.drawPile.shuffle()           // shuffle deck
	firstCard := game.drawPile.peak() // top discard pile card
	game.dealCards()                  // deal cards
	game.hasStarted = true
	game.color = firstCard.color

	for firstCard.kind == WILD || firstCard.kind == WILD_DRAW_FOUR {
		game.drawPile.shuffle() // shuffle deck
	}

	// set first card upside up
	card := game.drawPile.pop()
	game.discardPile.push(card)
}

func (game *Game) isGameOver() bool {
	playersWithCards := 0
	for _, p := range game.players {
		if len(p.cards) > 0 {
			playersWithCards++
		}
	}

	return playersWithCards == 1
}

func (game *Game) AddPlayer(name string) {
	player := (&Player{}).init(name)
	game.players = append(game.players, player)
}

func (game *Game) addDeck(d *Deck) {
	game.drawPile = *d
}

func (game *Game) dealCards() {
	if len(game.players) == 0 {
		log.Panic("You have to add players to the game first.")
	}

	// give INIT_NUMBER_OF_CARDS to each player
	j := 0
	for j < INIT_NUMBER_OF_CARDS {
		for i, player := range game.players {
			card := game.drawPile.pop()
			cards := append(player.cards, *card)
			game.players[i].cards = cards
		}
		j++
	}
}

func (game *Game) getPlayer() *Player {
	return &game.players[game.turn]
}

func (game *Game) getPlayerCard(i int) *Card {
	return &(game.getPlayer().cards[i])
}

func (game *Game) canStart() bool {
	return len(game.players) > 1
}

func (game *Game) peakDiscardPile() *Card {
	return game.discardPile.peak()
}

func (game *Game) pushDiscardPile(card *Card) {
	game.discardPile = append(game.discardPile, *card)
}

func (game *Game) popDrawCardPile() *Card {
	const emptyPile int = 0
	if len(game.drawPile) > emptyPile {
		return game.drawPile.pop()
	}

	// removes all cards from discard pile - except top most -
	// and are placed as draw pile - shuffled too.
	game.drawPile = game.discardPile[:len(game.discardPile)-1]
	game.discardPile = append([]Card{}, game.discardPile[len(game.discardPile)-1])
	game.drawPile.shuffle()

	return game.drawPile.pop()
}

func (game *Game) drawCards(player *Player) int {
	const minCardsToDraw float64 = 1
	numberDrawnCards := int(math.Max(minCardsToDraw, game.drawAcc))

	for i := 0; i < numberDrawnCards; i++ {
		player.drawCard(game.popDrawCardPile())
	}

	game.drawAcc = 0
	return numberDrawnCards
}

func (game *Game) usePlayerCardAt(i int) *Card {
	player := game.getPlayer()
	card := player.removeCardAt(i) // remove card from players hand
	game.pushDiscardPile(&card)    // place card on top of discard pile
	game.nextTurn(&card)           // calculate next turn
	game.useAnyCard = false        // next player has to follow this card
	game.color = card.color

	switch card.kind {
	case DRAW_TWO:
		game.drawAcc += 2
	case WILD_DRAW_FOUR:
		game.drawAcc += 4
	}

	if len(player.cards) == 0 {
		game.winners = append(game.winners, *player)
		player.hasWon = true
	}

	return &card
}

func (game *Game) skipPlayer() {
	game.turn += game.direction
	game.rebaseTurn()
}

func (game *Game) nextTurn(card *Card) {
	if card == nil { // player has drawn cards
		game.turn += game.direction
		game.rebaseTurn()

		topCard := *game.peakDiscardPile()
		if topCard.kind == DRAW_TWO || topCard.kind == WILD_DRAW_FOUR {
			game.useAnyCard = true // next player can use any card
		}
		game.drawAcc = 0
		return
	}

	if card.kind == REVERSE {
		game.direction *= -1 // reverse direction
	}

	game.turn += game.direction
	if card.kind == SKIP {
		game.turn += game.direction
	}
	game.rebaseTurn()
}

func (game *Game) rebaseTurn() {
	if game.turn < 0 {
		game.turn = game.turn + len(game.players)
	}
	if game.turn >= len(game.players) {
		game.turn = game.turn - len(game.players)
	}
}
