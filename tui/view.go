package tui

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/Guilherme-Costa-Dev/tuiminesweeper/game"
)

const (
	tnBg       = "#1a1b26"
	tnBgDark   = "#1f2335"
	tnBgHigh   = "#3b4261"
	tnBgSel    = "#7aa2f7"
	tnBgReveal = "#24283b"
	tnFg       = "#c0caf5"
	tnBlue     = "#7aa2f7"
	tnCyan     = "#7dcfff"
	tnGreen    = "#9ece6a"
	tnOrange   = "#ff9e64"
	tnRed      = "#f7768e"
	tnPurple   = "#bb9af7"
	tnMagenta  = "#bb9af7"
	tnComment  = "#565f89"
)

func fgHex(hex string) string { return "\x1b[38;2;" + rgbOf(hex) + "m" }
func bgHex(hex string) string { return "\x1b[48;2;" + rgbOf(hex) + "m" }

func rgbOf(hex string) string {
	hex = strings.TrimPrefix(hex, "#")
	var r, g, b int
	fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	return fmt.Sprintf("%d;%d;%d", r, g, b)
}

func parseColor(hex string) color.Color {
	hex = strings.TrimPrefix(hex, "#")
	var r, g, b int
	fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

var numberColor = map[int]string{
	1: tnBlue,
	2: tnGreen,
	3: tnRed,
	4: tnPurple,
	5: tnOrange,
	6: tnCyan,
	7: tnFg,
	8: tnComment,
}

func glyph(c game.Cell, underCursor bool) (text, fg, bg string) {
	switch {
	case c.Revealed && c.Mine:
		return " *", tnRed, tnBgDark
	case c.Revealed && c.Around > 0:
		return " " + strconv.Itoa(c.Around), numberColor[c.Around], tnBgReveal
	case c.Revealed:
		return "  ", tnFg, tnBgReveal
	case c.Flagged:
		return " F", tnMagenta, tnBgHigh
	default:
		return "  ", tnFg, tnBgHigh
	}
}

func (m Model) View() tea.View {
	const (
		cellW = 2
		padX  = 1
	)

	width := m.Width
	if width <= 0 {
		width = 80
	}
	height := m.Height
	if height <= 0 {
		height = 24
	}

	boardW := m.Board.Size * cellW
	panelW := boardW + padX*2
	blockH := m.Board.Size + 2

	if panelW > width {
		panelW = width
	}
	topPad := (height - blockH - 2) / 2
	if topPad < 0 {
		topPad = 0
	}
	leftPad := (width - panelW) / 2
	if leftPad < 0 {
		leftPad = 0
	}

	var out strings.Builder
	out.WriteString(bgHex(tnBg))

	fullLine := func() {
		out.WriteString(strings.Repeat(" ", width))
	}

	row := func(content func()) {
		out.WriteString(bgHex(tnBg))
		out.WriteString(strings.Repeat(" ", leftPad))
		content()
		if used := leftPad + panelW; used < width {
			out.WriteString(bgHex(tnBg))
			out.WriteString(strings.Repeat(" ", width-used))
		}
		out.WriteString("\n")
	}

	for i := 0; i < topPad; i++ {
		fullLine()
		out.WriteString("\n")
		out.WriteString(bgHex(tnBg))
	}

	row(func() {
		out.WriteString(bgHex(tnBgDark))
		out.WriteString(strings.Repeat(" ", panelW))
	})

	for y := 0; y < m.Board.Size; y++ {
		row(func() {
			out.WriteString(bgHex(tnBgDark))
			out.WriteString(strings.Repeat(" ", padX))
			for x := 0; x < m.Board.Size; x++ {
				text, fg, bg := glyph(m.Board.Cells[y][x], x == m.CursorX && y == m.CursorY)
				if x == m.CursorX && y == m.CursorY {
					bg = tnBgSel
					fg = tnBg
				}
				out.WriteString(bgHex(bg))
				out.WriteString(fgHex(fg))
				out.WriteString(text)
			}
			out.WriteString(bgHex(tnBgDark))
			out.WriteString(strings.Repeat(" ", padX))
		})
	}

	row(func() {
		out.WriteString(bgHex(tnBgDark))
		out.WriteString(strings.Repeat(" ", panelW))
	})

	fullLine()
	out.WriteString("\n")
	out.WriteString(bgHex(tnBg))

	status, statusColor := fmt.Sprintf("%d mines  %dx%d", m.Board.MineCount, m.Board.Size, m.Board.Size), tnComment
	switch {
	case m.GameOver:
		status, statusColor = "Game Over   r restart   q quit", tnRed
	case m.Won:
		status, statusColor = "You Won!   r restart   q quit", tnGreen
	}
	centeredText(&out, status, statusColor, width)

	centeredText(&out, "arrows/hjkl move  ·  space reveal  ·  f flag  ·  r restart  ·  q quit", tnComment, width)

	out.WriteString("\x1b[0m")

	v := tea.NewView(out.String())
	v.AltScreen = true
	v.BackgroundColor = parseColor(tnBg)
	return v
}

func centeredText(out *strings.Builder, text, colorHex string, width int) {
	runes := []rune(text)
	if len(runes) > width {
		runes = runes[:width]
	}
	padLeft := (width - len(runes)) / 2

	out.WriteString(bgHex(tnBg))
	out.WriteString(strings.Repeat(" ", padLeft))
	out.WriteString(fgHex(colorHex))
	out.WriteString(string(runes))
	out.WriteString(bgHex(tnBg))
	if rest := width - padLeft - len(runes); rest > 0 {
		out.WriteString(strings.Repeat(" ", rest))
	}
	out.WriteString("\n")
}
