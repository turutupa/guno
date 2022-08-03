package guno

const INIT_NUMBER_OF_CARDS = 7

// cli
type Cli struct{}

const DEFAULT_CURSOR_POS int = 0

// cli - select options
const ADD_PLAYER = "Add new player"
const START = "Start game"
const USE_CARD = "Use card"
const PEAK_OPPONENTS = "Peak opponents"
const PEAK_PILE = "Peak top card in discard pile"
const DRAW_CARD = "Draw card"

var userActionsPos = map[string]int{
	USE_CARD:       0,
	DRAW_CARD:      1,
	PEAK_OPPONENTS: 2,
	PEAK_PILE:      3,
}

// colors
const FMT_RED = "\033[31m"
const FMT_GREEN = "\033[32m"
const FMT_YELLOW = "\033[33m"
const FMT_BLUE = "\033[34m"

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

// direction
const DEFAULT_DIRECTION = 1

type Game struct {
	players     []Player // list of players
	turn        int      // specified whose turn it is
	discardPile Deck     // used cards
	drawPile    Deck     // unused cards
	hasStarted  bool     // game has started
	winners     []Player // ordered of winning players
	drawAcc     float64  // how much drawing cards next player has to draw
	direction   int      // -1 for left, +1 for right
	useAnyCard  bool     // player drew cards, next player may use any card
	color       string
}

type Player struct {
	name       string
	cards      Deck
	lastAction string
	hasWon     bool
}

type Card struct {
	kind   string // this I wanted to call "type" but that is a keyword
	number int
	color  string
}
