package main

import (
	"log"
	"tuiminesweeper/internal"
)

func main() {
	size := 30
	mineCount := 200

	board, err := internal.NewBoard(size)
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
