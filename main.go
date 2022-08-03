package main

import "uno/guno"

func main() {
	game := guno.Game{}
	deck := guno.Build() // get deck
	game.Init(&deck)

	// For testing only
	// game.AddPlayer("Alberto")
	// game.AddPlayer("suby")
	// game.AddPlayer("zarreta")

	cli := guno.Cli{}
	cli.Run(&game)

	main()
}
