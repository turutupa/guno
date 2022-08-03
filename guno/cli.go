package guno

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
)

func (cli *Cli) Run(game *Game) {
	for !(*game).hasStarted {
		initPhase(game)
	}

	// inform of top most discardPile card before starting
	peakPileSelect(game)

	for !game.isGameOver() {
		gamePhase(game, game.getPlayer())
	}

	showWinners(game)
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

	_, selected := selector(label, items, DEFAULT_CURSOR_POS)

	if selected == START {
		startGameSelect(game)
	} else if selected == ADD_PLAYER {
		addPlayerSelect(game)
	}
}

func gamePhase(game *Game, player *Player) {
	if player.hasWon {
		game.skipPlayer()
		return
	}

	if !game.useAnyCard && !player.hasValidCard(game.peakDiscardPile()) {
		drawnCards := game.drawCards(player)
		print(player.name+" has no valid cards - Drew "+strconv.Itoa(drawnCards)+" cards", RED)
		game.nextTurn(nil)
		return
	}

	player.cards.sortByColor()

	label := game.getPlayer().name + " what do you want to do?"
	items := []string{USE_CARD, DRAW_CARD, PEAK_OPPONENTS, PEAK_PILE}
	cursorPos := getCursorPos(player)

	_, userAction := selector(label, items, cursorPos)

	if userAction == USE_CARD {
		useCardSelect(game)
	} else if userAction == PEAK_OPPONENTS {
		peakOpponentsSelect(game)
	} else if userAction == PEAK_PILE {
		peakPileSelect(game)
	} else if userAction == DRAW_CARD {
		drawCardSelect(game)
	} else {
		fmt.Println("Oops! Didn't understand you. Please say again")
	}
}

func startGameSelect(game *Game) {
	game.start()
}

func addPlayerSelect(game *Game) {
	name := prompt("Name of the player")

	if name != "" {
		game.AddPlayer(name)
	}
}

func useCardSelect(game *Game) *Card {
	label := game.getPlayer().name + " what card do you want to use?"
	items := []string{}

	// player cards
	for _, card := range game.getPlayer().cards {
		items = append(items, getCardText(&card))
	}

	items = append(items, "Go back")
	i, _ := selector(label, items, DEFAULT_CURSOR_POS)
	if i == len(items)-1 {
		return nil
	}

	selectedCard := game.getPlayerCard(i)
	isValidCard := isValidCard(game.peakDiscardPile(), selectedCard)

	if !game.useAnyCard && !isValidCard {
		print("Selected card is not valid. Trying another one or draw a card", RED)
		return nil
	}

	if selectedCard != nil && (selectedCard.kind == WILD ||
		selectedCard.kind == WILD_DRAW_FOUR) {
		newColor := useWildCardSelect(game)
		if !newColor {
			return nil
		}
	}

	print(game.getPlayer().name+" used "+getCardText(selectedCard), selectedCard.color)
	return game.usePlayerCardAt(i)
}

func useWildCardSelect(game *Game) bool {
	items := []string{}

	for _, color := range COLORS {
		items = append(items, color)
	}

	goBackLabel := "Go back"
	items = append(items, "Go back")
	label := "What color do you want to change to?"
	_, selected := selector(label, items, DEFAULT_CURSOR_POS)

	if selected == goBackLabel {
		return false
	}

	game.color = selected
	return true
}

func peakOpponentsSelect(game *Game) {
	for _, player := range game.players {
		fmt.Println(player.name + " has " + strconv.Itoa(len(player.cards)) + " cards left")
	}
}

func peakPileSelect(game *Game) {
	card := game.peakDiscardPile()

	if card == nil {
		print("Where did the top card go?", RED)
	}

	prefix := "Top card is: "
	cardText := getCardText(card)
	if card.color == "" {
		fmt.Println(prefix + cardText)
	}
	print(prefix+cardText, card.color)
}

func drawCardSelect(game *Game) {
	drawn := game.drawCards(game.getPlayer())
	fmt.Printf("%s drew %d cards\n", game.getPlayer().name, drawn)
	game.nextTurn(nil) // no cards used
}

func showWinners(game *Game) {
	fmt.Println("Game ended! The winners are: ")
	for _, p := range game.players {
		fmt.Println(p.name)
	}

	fmt.Println("")
	fmt.Println("gUNO developed by turutupa")
}

func getCursorPos(player *Player) int {
	lastUserAction := player.lastAction
	cursorPos := 0
	if lastUserAction == PEAK_OPPONENTS || lastUserAction == PEAK_PILE {
		cursorPos = userActionsPos[lastUserAction]
	}

	return cursorPos
}

func prompt(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, _ := prompt.Run()
	return result
}

func selector(label string, items []string, cursorPos int) (int, string) {
	prompt := promptui.Select{
		Items:        items,
		Size:         9999, // so every line displays
		CursorPos:    cursorPos,
		HideSelected: true,
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
