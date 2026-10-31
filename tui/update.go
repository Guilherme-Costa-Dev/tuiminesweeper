package tui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if ws, ok := msg.(tea.WindowSizeMsg); ok {
		m.Width = ws.Width
		m.Height = ws.Height
		return m, nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	s := key.String()

	if s == "q" || s == "ctrl+c" {
		return m, tea.Quit
	}
	if m.GameOver || m.Won {
		if s == "r" {
			return m.reset(), nil
		}
		return m, nil
	}

	switch s {
	case "up", "k":
		if m.CursorY > 0 {
			m.CursorY--
		}
	case "down", "j":
		if m.CursorY < m.Board.Size-1 {
			m.CursorY++
		}
	case "left", "h":
		if m.CursorX > 0 {
			m.CursorX--
		}
	case "right", "l":
		if m.CursorX < m.Board.Size-1 {
			m.CursorX++
		}
	case "space", " ":
		if err := m.Board.Reveal(m.CursorX, m.CursorY); err == nil {
			if m.Board.Lost {
				m.GameOver = true
				m.Board.RevealAll()
			} else if m.Board.Won() {
				m.Won = true
			}
		}
	case "f":
		_ = m.Board.Flag(m.CursorX, m.CursorY)
	case "r":
		return m.reset(), nil
	}

	return m, nil
}
