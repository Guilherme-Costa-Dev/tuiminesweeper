package game

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Size        int
	Cells       [][]Cell
	MineCount   int
	RevealCount int
	Lost        bool
}

func NewBoard(size int) (*Board, error) {
	if size == 0 {
		return nil, fmt.Errorf("invalid size of board: 0")
	}
	var b Board
	b.Size = size
	b.Cells = make([][]Cell, size)
	for i := range b.Cells {
		b.Cells[i] = make([]Cell, size)
	}
	return &b, nil
}

func (b *Board) PlaceMines(mines int) error {
	if mines == 0 || mines > b.Size*b.Size {
		return fmt.Errorf("invalid number of mines: 0 or bigger than the board")
	}
	pos := rand.Perm(b.Size * b.Size)
	pos = pos[0:mines]

	for _, p := range pos {
		y := p / b.Size
		x := p % b.Size
		b.Cells[y][x].Mine = true
		b.MineCount += 1
	}
	return nil
}

func (b *Board) CalcAdjacents() {
	for i := range b.Cells {
		for j, cell := range b.Cells[i] {
			var count int
			if !cell.Mine {
				if j != 0 {
					if b.Cells[i][j-1].Mine {
						count += 1
					}
				}
				if j != b.Size-1 {
					if b.Cells[i][j+1].Mine {
						count += 1
					}
				}
				if i != 0 {
					if b.Cells[i-1][j].Mine {
						count += 1
					}
					if j != 0 {
						if b.Cells[i-1][j-1].Mine {
							count += 1
						}
					}
					if j != b.Size-1 {
						if b.Cells[i-1][j+1].Mine {
							count += 1
						}
					}
				}
				if i != b.Size-1 {
					if b.Cells[i+1][j].Mine {
						count += 1
					}
					if j != 0 {
						if b.Cells[i+1][j-1].Mine {
							count += 1
						}
					}
					if j != b.Size-1 {
						if b.Cells[i+1][j+1].Mine {
							count += 1
						}
					}
				}
			}
			b.Cells[i][j].Around = count
		}
	}
}

func (b *Board) PrintDebugBoard() {
	for i := range b.Cells {
		for j := range b.Cells[i] {
			if b.Cells[i][j].Mine {
				fmt.Print("X ")
			} else if b.Cells[i][j].Revealed {
				fmt.Print(b.Cells[i][j].Around, " ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
	fmt.Println("Mine Count:", b.MineCount)
	fmt.Println("Board Size:", b.Size*b.Size)
	fmt.Println("Reveal Count:", b.RevealCount)
}

func (b *Board) Reveal(x, y int) error {
	if x > b.Size-1 || y > b.Size-1 || x < 0 || y < 0 {
		return fmt.Errorf("invalid coordinates to reveal: (%d,%d)", x, y)
	}
	if b.Cells[y][x].Flagged || b.Cells[y][x].Revealed {
		return nil
	}
	if b.Cells[y][x].Mine {
		b.Lost = true
		return nil
	}
	b.Cells[y][x].Revealed = true
	b.RevealCount++
	if b.Cells[y][x].Around == 0 {
		if y != 0 {
			if !b.Cells[y-1][x].Revealed {
				b.Reveal(x, y-1)
			}
		}
		if x != 0 {
			if !b.Cells[y][x-1].Revealed {
				b.Reveal(x-1, y)
			}
		}
		if x != b.Size-1 {
			if !b.Cells[y][x+1].Revealed {
				b.Reveal(x+1, y)
			}
		}
		if y != b.Size-1 {
			if !b.Cells[y+1][x].Revealed {
				b.Reveal(x, y+1)
			}
		}
	}
	return nil
}

func (b *Board) Flag(x, y int) error {
	if x < 0 || y < 0 || x > b.Size-1 || y > b.Size-1 {
		return fmt.Errorf("invalid coordinates to flag: (%d,%d)", x, y)
	}
	if b.Cells[y][x].Revealed {
		return nil
	}
	if b.Cells[y][x].Flagged {
		b.Cells[y][x].Flagged = false
		return nil
	}
	b.Cells[y][x].Flagged = true
	return nil
}

func (b *Board) Won() bool {
	if b.RevealCount == b.Size*b.Size-b.MineCount {
		return true
	}
	return false
}

func (b *Board) RevealAll() {
	for i := range b.Cells {
		for j := range b.Cells[i] {
			b.Cells[i][j].Revealed = true
		}
	}
}
