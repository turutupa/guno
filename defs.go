package main

const INIT_NUMBER_OF_CARDS = 7

// cli
type Cli struct{}

const ADD_PLAYER = "Add new player"
const START = "Start game"
const USE_CARD = "Use card"
const PEAK_OPPONENTS = "Peak opponents"
const PEAK_PILE = "Peak top card in pile"
const DRAW_CARD = "Draw card"

// types of cards
const NUMBER = "number"
const WILD = "wild"                     // Change the color being played to any color.
const SKIP = "skip"                     // The next player loses his/her turn and is "skipped".
const REVERSE = "reverse"               // The direction of play is reversed.
const DRAW_TWO = "draw_two"             // The next player must draw 2 cards and forfeit the turn.
const WILD_DRAW_FOUR = "wild_draw_four" // Choose the next color played and force the next player to pick 4 cards and forfeit his/her turn.
var SPECIAL_CARDS = [5]string{
	WILD,
	SKIP,
	REVERSE,
	DRAW_TWO,
	WILD_DRAW_FOUR,
}

// colors
const RED = "red"
const YELLOW = "yellow"
const GREEN = "green"
const BLUE = "blue"

var COLORS = [4]string{RED, YELLOW, GREEN, BLUE}

// game
type Game struct {
	players    []Player
	turn       int
	cemetery   Deck // used cards
	pile       Deck // this is the pile of cards pre-dealing to players
	isDone     bool
	hasStarted bool
}

// player
type Player struct {
	game  Game
	name  string
	cards []Card
}

type playerActions interface {
	useCard()
	drawCard()
	hasColor(string)
	shoutUNO()
}

// card
type Card struct {
	kind   string // this I wanted to call "type" but that is a keyword
	number int
	color  string
}
