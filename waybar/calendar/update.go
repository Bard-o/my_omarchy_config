package main

import tea "github.com/charmbracelet/bubbletea"

// Init returns the initial command for the Bubbletea program.
// For the calendar, no initial action is needed — the view renders immediately.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles incoming Bubbletea messages and delegates to the appropriate handler.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyboard(msg)
	case tea.MouseMsg:
		return m.handleMouse(msg)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}
	return m, nil
}
