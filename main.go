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

	err := validate(*size, *mineCount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	p := tea.NewProgram(tui.New(*size, *mineCount))
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func validate(size, mines int) error {
	if size <= 0 {
		return fmt.Errorf("size must be positive, got %d", size)
	}
	total := size * size
	if mines <= 0 {
		return fmt.Errorf("mines must be positive, got %d", mines)
	}
	if mines >= total {
		return fmt.Errorf("mines must be at most %d for a %dx%d board, got %d", total-1, size, size, mines)
	}
	return nil
}
