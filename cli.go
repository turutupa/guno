package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
)

func (cli *Cli) run(game *Game) {
	for !(*game).hasStarted {
		initPhase(game)
	}

	// inform of top most discardPile card before starting
	peakPileSelect(game)

	for !game.isGameOver() {
		gamePhase(game)
	}
}

func initPhase(game *Game) {
	var label string
	if len(game.players) > 0 {
		label = "Choose one"
	} else {
		label = "Welcome to UNO! Please follow the instructions:"
	}

	items := []string{ADD_PLAYER}
	if game.canStart() {
		items = append(items, START)
	}

	_, selected := selector(label, items)

	if selected == START {
		startSelect(game)
	} else if selected == ADD_PLAYER {
		addPlayerSelect(game)
	}
}

func gamePhase(game *Game) {
	label := game.getPlayer().name + " what do you want to do?"
	items := []string{USE_CARD, DRAW_CARD, PEAK_OPPONENTS, PEAK_PILE}
	_, selected := selector(label, items)

	if selected == USE_CARD {
		useCardSelect(game)
	} else if selected == PEAK_OPPONENTS {
		peakOpponentsSelect(game)
	} else if selected == PEAK_PILE {
		peakPileSelect(game)
	} else if selected == DRAW_CARD {
		drawCardSelect(game)
	} else {
		fmt.Println("Oops! Didn't understand you. Please say again")
	}
}

func startSelect(game *Game) {
	game.hasStarted = true
	game.start()

	for _, p := range game.players {
		fmt.Println(p.name + " has " + strconv.Itoa(len(p.cards)))
	}

	// TODO: delete
	fmt.Println("discardPile has " + strconv.Itoa(len(game.discardPile)))
	fmt.Println("drawPile has " + strconv.Itoa(len(game.drawPile)))
}

func addPlayerSelect(game *Game) {
	name := prompt("Name of the player")
	game.addPlayer(name)
}

func useCardSelect(game *Game) {
	label := "What card do you want to use?"
	items := []string{}

	// player cards
	for _, card := range game.getPlayer().cards {
		items = append(items, getCardText(&card))
	}

	items = append(items, "Go back")
	i, selected := selector(label, items)
	if i == len(items)-1 {
		return
	}

	isValidCard := game.isValidCard(game.getPlayerCard(i))

	if !isValidCard {
		print("Selected card is not valid. Trying another one or draw a card", RED)
		return
	}

	fmt.Println("You selected: " + selected)
}

func peakOpponentsSelect(game *Game) {
	for _, player := range game.players {
		fmt.Println(player.name + " has " + strconv.Itoa(len(player.cards)) + " cards left")
	}
}

func peakPileSelect(game *Game) {
	card := game.getTopDiscardPileCard()
	print("Top card is: "+getCardText(card), card.color)
}

func drawCardSelect(game *Game) {}

func prompt(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, _ := prompt.Run()
	return result
}

func selector(label string, items []string) (int, string) {
	prompt := promptui.Select{
		Items: items,
		Size:  9999, // so every line displays
	}

	if len(label) > 0 {
		prompt.Label = label
	}

	i, result, err := prompt.Run()

	if err != nil {
		fmt.Println("Hope to see you back!")
		os.Exit(0)
	}

	return i, result
}
