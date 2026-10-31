package tui

import "tuiminesweeper/game"

type Model struct {
	Board    *game.Board
	CursorX  int
	CursorY  int
	GameOver bool
	Won      bool

	Width  int
	Height int
}

func New(size, mines int) Model {
	return NewWithSize(size, mines, 80, 24)
}

func NewWithSize(size, mines, width, height int) Model {
	b, err := game.NewBoard(size)
	if err != nil {
		panic(err)
	}
	if err := b.PlaceMines(mines); err != nil {
		panic(err)
	}
	b.CalcAdjacents()

	return Model{
		Board:   b,
		CursorX: size / 2,
		CursorY: size / 2,
		Width:   width,
		Height:  height,
	}
}

func (m Model) reset() Model {
	return NewWithSize(m.Board.Size, m.Board.MineCount, m.Width, m.Height)
}
