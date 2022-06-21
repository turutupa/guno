package main

import "fmt"

func print(str string, color string) {
	if color == RED {
		printColor(str, FMT_RED)
	}

	if color == GREEN {
		printColor(str, FMT_GREEN)
	}

	if color == BLUE {
		printColor(str, FMT_BLUE)
	}

	if color == YELLOW {
		printColor(str, FMT_YELLOW)
	}
}

func printColor(str string, color string) {
	fmt.Println(color, str)
}
