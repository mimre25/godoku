package main

import (
	"fmt"
	"godoku/game"
)

func main() {
	var g = game.GameFromFile("s1.txt")

	g.PrintGame()
	fmt.Println(g.IsFinished())
	g.Solve()

	g.PrintGame()
	g = game.GameFromFile("s2.txt")

	fmt.Println()

	g.PrintGame()
	fmt.Println(g.IsFinished())

}
