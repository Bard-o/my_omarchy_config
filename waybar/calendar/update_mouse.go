package main

import tea "github.com/charmbracelet/bubbletea"

// handleMouse processes mouse input for header-based month navigation.
// Clicking the left half of the header goes to the previous month.
// Clicking the right half of the header goes to the next month.
func (m Model) handleMouse(msg tea.MouseMsg) (Model, tea.Cmd) {
	if msg.Type == tea.MouseLeft {
		// Header is the first row (Y=0) in the borderless design
		if msg.Y == 0 {
			midX := m.width / 2
			if msg.X < midX {
				m = m.prevMonth()
			} else {
				m = m.nextMonth()
			}
		}
	}
	return m, nil
}
