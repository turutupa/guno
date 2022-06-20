package main

import (
	"fmt"
	"log"
	"strconv"
)

func (game *Game) init() {
	game.players = make([]Player, 0) // init players arr
	game.turn = 0                    // init turn
	game.cemetery = make(Deck, 0)    // init cemetery
	game.pile = make(Deck, 0)        // init pile
}

func (game *Game) start() {
	game.pile = buildDeck() // get deck
	shuffleDeck(&game.pile) // shuffle deck
	game.dealCards()        // deal cards
}

func (game *Game) addPlayer(name string) {
	player := (&Player{}).init(name)
	game.players = append(game.players, player)
}

func (game *Game) addDeck(d *Deck) {
	game.pile = *d
}

func (game *Game) dealCards() {
	if len(game.players) == 0 {
		log.Panic("You have to add players to the game first.")
	}

	j := 0
	for j < INIT_NUMBER_OF_CARDS {
		for i, player := range game.players {
			card := game.pile[len(game.pile)-1]
			cards := append(player.cards, card)
			game.players[i].cards = cards
			game.pile = game.pile[:len(game.pile)-1]
		}
		j++
	}
}

func (game *Game) getPlayer() Player {
	return game.players[game.turn]
}

func (game *Game) canStart() bool {
	return len(game.players) > 1
}

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
		fmt.Println(prefix + getCardText(card))
	}
}

func getCardText(card Card) string {
	if card.kind == NUMBER {
		return "[" + strconv.Itoa(card.number) + "] " + card.color
	} else if card.kind == WILD {
		return "Wild: Change color"
	} else if card.kind == WILD_DRAW_FOUR {
		return "Wild + draw four!"
	} else if card.kind == SKIP {
		return "Skip player"
	} else if card.kind == DRAW_TWO {
		return "Draw two! " + card.color
	} else if card.kind == REVERSE {
		return "Reverse! " + card.color
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
