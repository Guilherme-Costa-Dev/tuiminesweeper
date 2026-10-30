package tuiminesweeper

import (
	"log"
)

func main() {
	size := 30
	mineCount := 200

	board, err := NewBoard(size)
	if err != nil {
		log.Fatalf("invalid config:", err)
	}
	err = board.PlaceMines(mineCount)
	if err != nil {
		log.Fatalf("invalid config:", err)
	}
	board.CalcAdjacents()
}
