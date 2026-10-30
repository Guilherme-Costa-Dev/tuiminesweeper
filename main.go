package main

import (
	"log"
	"tuiminesweeper/game"
)

func main() {
	size := 30
	mineCount := 150

	board, err := game.NewBoard(size)
	if err != nil {
		log.Fatalf("invalid config:", err)
	}
	err = board.PlaceMines(mineCount)
	if err != nil {
		log.Fatalf("invalid config:", err)
	}
	board.CalcAdjacents()
	board.PrintBoard()
}
