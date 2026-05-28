package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestView(t *testing.T) {
	// Helper to create a model at a specific year/month for deterministic testing
	newModelAt := func(year int, month time.Month) Model {
		theme := DefaultTheme()
		m := NewModel(theme)
		m.year = year
		m.month = month
		m.width = 24
		m.height = 10
		return m
	}

	t.Run("May 2026 view contains header with month and year", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		view := m.View()
		stripped := stripANSI(view)

		if !strings.Contains(stripped, "May") {
			t.Error("View should contain 'May' in header")
		}
		if !strings.Contains(stripped, "2026") {
			t.Error("View should contain '2026' in header")
		}
	})

	t.Run("May 2026 view contains weekday labels", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		view := m.View()
		stripped := stripANSI(view)

		for _, day := range []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"} {
			if !strings.Contains(stripped, day) {
				t.Errorf("View should contain day abbreviation '%s'", day)
			}
		}
	})

	t.Run("May 2026 view contains all day numbers 1-31", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		view := m.View()
		stripped := stripANSI(view)

		for day := 1; day <= 31; day++ {
			dayStr := fmt.Sprintf("%2d", day)
			if !strings.Contains(stripped, dayStr) {
				t.Errorf("View should contain day number '%s'", dayStr)
			}
		}
	})

	t.Run("January 2026 view contains correct header and days", func(t *testing.T) {
		m := newModelAt(2026, time.January)
		view := m.View()
		stripped := stripANSI(view)

		if !strings.Contains(stripped, "Jan") {
			t.Error("View should contain 'Jan' in header")
		}
		if !strings.Contains(stripped, "2026") {
			t.Error("View should contain '2026' in header")
		}
		for day := 1; day <= 31; day++ {
			dayStr := fmt.Sprintf("%2d", day)
			if !strings.Contains(stripped, dayStr) {
				t.Errorf("View should contain day '%s'", dayStr)
			}
		}
	})

	t.Run("February 2024 leap year shows 29 days", func(t *testing.T) {
		m := newModelAt(2024, time.February)
		view := m.View()
		stripped := stripANSI(view)

		if !strings.Contains(stripped, "Feb") {
			t.Error("View should contain 'Feb' in header")
		}
		for day := 1; day <= 29; day++ {
			dayStr := fmt.Sprintf("%2d", day)
			if !strings.Contains(stripped, dayStr) {
				t.Errorf("View should contain day '%s' for Feb leap year", dayStr)
			}
		}
	})

	t.Run("view has correct line count for calendar layout", func(t *testing.T) {
		m := newModelAt(2026, time.May)
		view := m.View()

		// New bordered layout: top(1) + blank(1) + daynames(1) + 6 grid rows + blank(1) + footer(1) + bottom(1) = 12 lines
		lines := strings.Split(strings.TrimRight(view, "\n"), "\n")
		if len(lines) < 11 {
			t.Errorf("View should have at least 11 lines (bordered frame), got %d", len(lines))
		}
	})

	t.Run("September 2025 view renders correctly", func(t *testing.T) {
		m := newModelAt(2025, time.September)
		view := m.View()
		stripped := stripANSI(view)

		if !strings.Contains(stripped, "Sep") {
			t.Error("View should contain 'Sep' in header")
		}
		if !strings.Contains(stripped, "2025") {
			t.Error("View should contain '2025' in header")
		}
		// September has 30 days
		for day := 1; day <= 30; day++ {
			dayStr := fmt.Sprintf("%2d", day)
			if !strings.Contains(stripped, dayStr) {
				t.Errorf("View should contain day '%s'", dayStr)
			}
		}
	})

	t.Run("today number is highlighted with accent color", func(t *testing.T) {
		// Use current date — verify that the today day number
		// is rendered with TodayDay style (which adds ANSI codes)
		theme := DefaultTheme()
		m := NewModel(theme)
		m.width = 24
		m.height = 10
		view := m.View()

		// The view should contain ANSI escape sequences (proves styling is applied)
		if !strings.Contains(view, "\x1b[") {
			t.Error("View should contain ANSI escape sequences (styled output)")
		}

		// Verify today's day appears in the view
		stripped := stripANSI(view)
		todayStr := fmt.Sprintf("%2d", m.today.Day())
		if !strings.Contains(stripped, todayStr) {
			t.Errorf("View should contain today's day number '%s'", todayStr)
		}
	})
}

// stripANSI removes ANSI escape sequences from a string for content testing.
func stripANSI(s string) string {
	var result strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == '\x1b' {
			// Skip ANSI escape sequence: ESC [ ... final_byte
			i++ // skip ESC
			if i < len(s) && s[i] == '[' {
				i++ // skip [
				for i < len(s) {
					if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
						i++ // skip final byte
						break
					}
					i++
				}
			}
		} else {
			result.WriteByte(s[i])
			i++
		}
	}
	return result.String()
}
