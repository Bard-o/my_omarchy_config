package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// handleKeyboard processes keyboard input for month navigation and quitting.
func (m Model) handleKeyboard(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "left", "h":
		m = m.prevMonth()
	case "right", "l":
		m = m.nextMonth()
	case "up", "k":
		m = m.prevMonth()
	case "down", "j":
		m = m.nextMonth()
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

// prevMonth decrements the month, wrapping from January back to December of the previous year.
func (m Model) prevMonth() Model {
	m.month--
	if m.month < time.January {
		m.month = time.December
		m.year--
	}
	return m
}

// nextMonth increments the month, wrapping from December to January of the next year.
func (m Model) nextMonth() Model {
	m.month++
	if m.month > time.December {
		m.month = time.January
		m.year++
	}
	return m
}
