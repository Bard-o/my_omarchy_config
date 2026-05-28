package main

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestHandleMouse(t *testing.T) {
	// Helper to create a model at a specific year/month with width set
	newModelAt := func(year int, month time.Month, width int) Model {
		theme := DefaultTheme()
		m := NewModel(theme)
		m.year = year
		m.month = month
		m.width = width
		m.height = 10
		return m
	}

	t.Run("left click on left half of header goes to previous month", func(t *testing.T) {
		m := newModelAt(2026, time.May, 24)
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 5, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.April {
			t.Errorf("click left half: month = %v, want April", result.month)
		}
	})

	t.Run("left click on right half of header goes to next month", func(t *testing.T) {
		m := newModelAt(2026, time.May, 24)
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 18, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.June {
			t.Errorf("click right half: month = %v, want June", result.month)
		}
	})

	t.Run("click below header does nothing", func(t *testing.T) {
		m := newModelAt(2026, time.May, 24)
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 5, Y: 2}
		result, _ := m.handleMouse(msg)
		if result.month != time.May {
			t.Errorf("click below header should not change month: got %v, want May", result.month)
		}
		if result.year != 2026 {
			t.Errorf("click below header should not change year: got %d, want 2026", result.year)
		}
	})

	t.Run("click exactly at midpoint goes to next month (right half inclusive)", func(t *testing.T) {
		m := newModelAt(2026, time.May, 24)
		// With width=24, midpoint is 12. X=12 is right half (>= midpoint)
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 12, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.June {
			t.Errorf("click at midpoint: month = %v, want June (right half inclusive)", result.month)
		}
	})

	t.Run("click just before midpoint goes to previous month", func(t *testing.T) {
		m := newModelAt(2026, time.May, 24)
		// With width=24, midpoint is 12. X=11 is left half (< midpoint)
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 11, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.April {
			t.Errorf("click just before midpoint: month = %v, want April", result.month)
		}
	})

	t.Run("header click does year wrapping January to December", func(t *testing.T) {
		m := newModelAt(2026, time.January, 24)
		// Click left half of header → previous month
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 5, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.December {
			t.Errorf("January left click: month = %v, want December", result.month)
		}
		if result.year != 2025 {
			t.Errorf("January left click: year = %d, want 2025", result.year)
		}
	})

	t.Run("right mouse button does nothing", func(t *testing.T) {
		m := newModelAt(2026, time.May, 24)
		msg := tea.MouseMsg{Type: tea.MouseRight, X: 18, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.May {
			t.Errorf("right mouse button should not change month: got %v, want May", result.month)
		}
	})

	t.Run("different width affects header region correctly", func(t *testing.T) {
		m := newModelAt(2026, time.May, 40)
		// With width=40, midpoint is 20. X=18 is left half.
		msg := tea.MouseMsg{Type: tea.MouseLeft, X: 18, Y: 0}
		result, _ := m.handleMouse(msg)
		if result.month != time.April {
			t.Errorf("width=40 click at X=18: month = %v, want April (left half)", result.month)
		}

		// X=22 is right half
		msg2 := tea.MouseMsg{Type: tea.MouseLeft, X: 22, Y: 0}
		m2 := newModelAt(2026, time.May, 40)
		result2, _ := m2.handleMouse(msg2)
		if result2.month != time.June {
			t.Errorf("width=40 click at X=22: month = %v, want June (right half)", result2.month)
		}
	})
}
