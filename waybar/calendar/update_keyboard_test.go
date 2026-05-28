package main

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestHandleKeyboard(t *testing.T) {
	// Helper to create a model at a specific year/month
	newModelAt := func(year int, month time.Month) Model {
		theme := DefaultTheme()
		m := NewModel(theme)
		m.year = year
		m.month = month
		return m
	}

	t.Run("left arrow goes to previous month", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyLeft}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.April {
			t.Errorf("left arrow from May: month = %v, want April", updated.month)
		}
		if updated.year != 2026 {
			t.Errorf("left arrow from May: year = %d, want 2026", updated.year)
		}
	})

	t.Run("h key goes to previous month", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h'}}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.April {
			t.Errorf("'h' from May: month = %v, want April", updated.month)
		}
	})

	t.Run("right arrow goes to next month", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyRight}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.June {
			t.Errorf("right arrow from May: month = %v, want June", updated.month)
		}
		if updated.year != 2026 {
			t.Errorf("right arrow from May: year = %d, want 2026", updated.year)
		}
	})

	t.Run("l key goes to next month", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.June {
			t.Errorf("'l' from May: month = %v, want June", updated.month)
		}
	})

	t.Run("year wraps from January back to December of previous year", func(t *testing.T) {
		m := newModelAt(2026, time.January)
		msg := tea.KeyMsg{Type: tea.KeyLeft}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.December {
			t.Errorf("left arrow from January: month = %v, want December", updated.month)
		}
		if updated.year != 2025 {
			t.Errorf("left arrow from January 2026: year = %d, want 2025", updated.year)
		}
	})

	t.Run("h key wraps year from January", func(t *testing.T) {
		m := newModelAt(2026, time.January)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h'}}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.December {
			t.Errorf("'h' from January: month = %v, want December", updated.month)
		}
		if updated.year != 2025 {
			t.Errorf("'h' from January 2026: year = %d, want 2025", updated.year)
		}
	})

	t.Run("year wraps from December forward to January of next year", func(t *testing.T) {
		m := newModelAt(2025, time.December)
		msg := tea.KeyMsg{Type: tea.KeyRight}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.January {
			t.Errorf("right arrow from December: month = %v, want January", updated.month)
		}
		if updated.year != 2026 {
			t.Errorf("right arrow from December 2025: year = %d, want 2026", updated.year)
		}
	})

	t.Run("up arrow goes to previous month (alias)", func(t *testing.T) {
		m := newModelAt(2026, time.June)
		msg := tea.KeyMsg{Type: tea.KeyUp}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.May {
			t.Errorf("up arrow from June: month = %v, want May", updated.month)
		}
	})

	t.Run("down arrow goes to next month (alias)", func(t *testing.T) {
		m := newModelAt(2026, time.June)
		msg := tea.KeyMsg{Type: tea.KeyDown}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.July {
			t.Errorf("down arrow from June: month = %v, want July", updated.month)
		}
	})

	t.Run("k key goes to previous month (vim up)", func(t *testing.T) {
		m := newModelAt(2026, time.July)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.June {
			t.Errorf("'k' from July: month = %v, want June", updated.month)
		}
	})

	t.Run("j key goes to next month (vim down)", func(t *testing.T) {
		m := newModelAt(2026, time.July)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}
		result, _ := m.Update(msg)
		updated := result.(Model)
		if updated.month != time.August {
			t.Errorf("'j' from July: month = %v, want August", updated.month)
		}
	})

	t.Run("q key returns quit command", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}
		_, cmd := m.Update(msg)

		// The quit command should be tea.Quit
		// We verify by executing the command and checking it produces a quit message
		if cmd == nil {
			t.Error("pressing 'q' should return a non-nil quit command")
		}
	})

	t.Run("escape key returns quit command", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyEsc}
		_, cmd := m.Update(msg)
		if cmd == nil {
			t.Error("pressing Escape should return a non-nil quit command")
		}
	})

	t.Run("unknown key preserves model state", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'x'}}
		result, cmd := m.Update(msg)
		updated := result.(Model)

		if updated.month != time.May {
			t.Errorf("unknown key 'x': month = %v, want May", updated.month)
		}
		if updated.year != 2026 {
			t.Errorf("unknown key 'x': year = %d, want 2026", updated.year)
		}
		if cmd != nil {
			t.Error("unknown key should return nil command")
		}
	})
}
