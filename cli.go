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

	for !(*game).isDone {
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
	items := []string{USE_CARD, PEAK_OPPONENTS}
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
}

func addPlayerSelect(game *Game) {
	name := prompt("Name of the player")
	game.addPlayer(name)
}

func useCardSelect(game *Game) {
	// game.players[game.turn].printCards()
	label := "What card do you want to use?"
	items := []string{}
	for _, card := range game.getPlayer().cards {
		items = append(items, getCardText(card))
	}
	items = append(items, "Go back")
	i, selected := selector(label, items)
	if i == len(items)-1 {
		return
	}

	fmt.Println("You selected: " + selected)
}

func peakOpponentsSelect(game *Game) {
	for _, player := range game.players {
		fmt.Println(player.name + " has " + strconv.Itoa(len(player.cards)) + " cards left")
	}
}

func peakPileSelect(game *Game) {}

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
