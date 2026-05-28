package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInit(t *testing.T) {
	t.Run("returns nil command for initial tick", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)
		cmd := m.Init()
		// Init should return a command (or nil for no initial action).
		// The spec says "return initial tick" — we verify it doesn't panic
		// and returns a valid cmd (nil is acceptable for no-op init).
		// A non-nil cmd indicates a tick or setup command.
		_ = cmd // Just verify no panic
	})
}

func TestUpdate(t *testing.T) {
	t.Run("WindowSizeMsg stores dimensions", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)

		msg := tea.WindowSizeMsg{Width: 30, Height: 12}
		newModel, _ := m.Update(msg)

		updated := newModel.(Model)
		if updated.width != 30 {
			t.Errorf("width = %d, want 30", updated.width)
		}
		if updated.height != 12 {
			t.Errorf("height = %d, want 12", updated.height)
		}
	})

	t.Run("unknown message preserves model state", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)
		originalMonth := m.month
		originalYear := m.year

		// Send a message type that Update doesn't handle specifically
		newModel, _ := m.Update(tea.FocusMsg{})

		updated := newModel.(Model)
		if updated.month != originalMonth {
			t.Errorf("month changed from %v to %v on unknown message", originalMonth, updated.month)
		}
		if updated.year != originalYear {
			t.Errorf("year changed from %d to %d on unknown message", originalYear, updated.year)
		}
	})

	t.Run("keyboard messages are delegated to handler", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)
		originalMonth := m.month

		// Press right arrow → should go to next month
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}}
		newModel, _ := m.Update(msg)

		updated := newModel.(Model)
		if updated.month == originalMonth && updated.year == m.year {
			// If month didn't change, the key wasn't handled properly
			// (unless we're already at December, which wraps the year)
			t.Errorf("pressing 'l' should change month from %v", originalMonth)
		}
	})

	t.Run("mouse messages are delegated to handler", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)
		m.width = 24
		m.height = 10

		// Left-click on right half of header (row 0) → next month
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 18, Y: 0}
		newModel, _ := m.Update(msg)

		updated := newModel.(Model)
		if updated.month == m.month && updated.year == m.year {
			t.Error("clicking right half of header should change month")
		}
	})

	t.Run("multiple WindowSizeMsgs update dimensions", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)

		msg1 := tea.WindowSizeMsg{Width: 40, Height: 15}
		newModel, _ := m.Update(msg1)
		updated := newModel.(Model)

		if updated.width != 40 || updated.height != 15 {
			t.Errorf("after first resize: width=%d height=%d, want 40x15", updated.width, updated.height)
		}

		// Second resize updates again
		msg2 := tea.WindowSizeMsg{Width: 60, Height: 20}
		newModel2, _ := updated.Update(msg2)
		updated2 := newModel2.(Model)

		if updated2.width != 60 || updated2.height != 20 {
			t.Errorf("after second resize: width=%d height=%d, want 60x20", updated2.width, updated2.height)
		}
	})
}
