package main

import (
	"log"
	"tuiminesweeper/game"
)

func main() {
	size := 30
	mineCount := 100

	board, err := game.NewBoard(size)
	if err != nil {
		log.Fatalf("invalid config:", err)
	}
	err = board.PlaceMines(mineCount)
	if err != nil {
		log.Fatalf("invalid config:", err)
	}
	board.CalcAdjacents()
	board.Reveal(15, 15)
	board.PrintDebugBoard()
}
