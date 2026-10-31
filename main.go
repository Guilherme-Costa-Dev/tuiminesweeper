package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"tuiminesweeper/tui"
)

func main() {
	var size int
	flag.IntVar(&size, "width", 30, "Defines the width of the board")
	flag.IntVar(&size, "w", 30, "Defines the width of the board")
	var mineCount int
	flag.IntVar(&mineCount, "mines", 150, "Defines the amount of mines in the board")
	flag.IntVar(&mineCount, "m", 150, "Defines the ammount of mines in the board")
	flag.Parse()

	p := tea.NewProgram(tui.New(size, mineCount))
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
