package main

import (
	"fmt"
	"log"
	"strconv"
)

// GAME
func (game *Game) init() {
	game.players = make([]Player, 0) // init players arr
	game.turn = 0                    // init turn
	game.discardPile = make(Deck, 0) // init discardPile
	game.drawPile = make(Deck, 0)    // init drawPile
}

func (game *Game) start() {
	game.drawPile = buildDeck() // get deck
	shuffleDeck(&game.drawPile) // shuffle deck
	game.dealCards()            // deal cards
}

func (game *Game) addPlayer(name string) {
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
			card := pop(&game.drawPile)
			cards := append(player.cards, *card)
			game.players[i].cards = cards
		}
		j++
	}

	card := pop(&game.drawPile)
	game.discardPile = append(game.discardPile, *card)
}

func (game *Game) getPlayer() Player {
	return game.players[game.turn]
}

func (game *Game) getPlayerCard(i int) *Card {
	return &game.getPlayer().cards[i]
}

func (game *Game) isValidCard(card *Card) bool {
	top := game.getTopDiscardPileCard()

	if card.kind == WILD || card.kind == WILD_DRAW_FOUR {
		return true
	}

	if card.kind != NUMBER && top.kind == card.kind {
		return true
	}

	if top.kind == NUMBER && card.kind == NUMBER {
		if (top.number == card.number) || top.color == card.color {
			return true
		}
	}

	return false
}

func (game *Game) canStart() bool {
	return len(game.players) > 1
}

func (game *Game) getTopDiscardPileCard() *Card {
	return peak(&game.discardPile)
}

func (game *Game) isGameOver() bool {
	playersWithCards := 0
	for _, p := range game.players {
		if len(p.cards) > 0 {
			playersWithCards++
		}
	}

	return playersWithCards == 0
}

func (game *Game) useCard(card *Card) {
	game.discardPile = append(game.discardPile, *card)
}

// PLAYER
func (player *Player) init(name string) Player {
	player.name = name
	player.cards = make([]Card, 0)
	return *player
}

func (player *Player) drawCard(card *Card) {
}

func (player *Player) printCards() {
	fmt.Println("What card do you want to use?")
	for i, card := range player.cards {
		index := i + 1
		prefix := "[" + strconv.Itoa(index) + "] "
		fmt.Println(prefix + getCardText(&card))
	}
}

func getCardText(card *Card) string {
	if card.kind == NUMBER {
		return "[" + strconv.Itoa(card.number) + "] [" + card.color + "]"
	} else if card.kind == WILD {
		return "[x] Wild: Change color"
	} else if card.kind == WILD_DRAW_FOUR {
		return "[x] Wild + draw four!"
	} else if card.kind == SKIP {
		return "[x] [" + card.color + "] Skip player"
	} else if card.kind == DRAW_TWO {
		return "[" + card.color + "] Draw two!"
	} else if card.kind == REVERSE {
		return "[" + card.color + "] Reverse!"
	} else {
		return ""
	}
}

func main() {
	game := Game{}
	cli := Cli{}

	game.init()
	cli.run(&game)
}
