package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Guilherme-Costa-Dev/tuiminesweeper/tui"
)

func main() {
	size := flag.Int("size", 30, "Defines the lateral size of the board")
	mineCount := flag.Int("mines", 120, "Defines the amount of mines in the board")
	flag.Parse()

	p := tea.NewProgram(tui.New(*size, *mineCount))
	_, err := p.Run()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
