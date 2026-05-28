package main

import (
	"testing"
	"time"
)

func TestMonthGrid(t *testing.T) {
	t.Run("May 2026 starts on Friday", func(t *testing.T) {
		// May 2026: 1st is a Friday (weekday=5), 31 days
		grid := monthGrid(2026, time.May)

		// Friday = 5 (Sunday=0), so day 1 is at column 5
		if grid[0][5] != 1 {
			t.Errorf("May 2026 day 1 at [0][5] = %d, want 1", grid[0][5])
		}
		// Previous days should be 0
		for col := 0; col < 5; col++ {
			if grid[0][col] != 0 {
				t.Errorf("May 2026 [0][%d] = %d, want 0 (prev month fill)", col, grid[0][col])
			}
		}
		// May 31: row 5, column 0 (Saturday of the last week)
		// Row 0: 0 0 0 0 0 1 2
		// Row 4: 26 27 28 29 30 31 ... wait let me count properly:
		// Row 4: Sun=24 Mon=25 Tue=26 Wed=27 Thu=28 Fri=29 Sat=30
		// Row 5: Sun=31 ... → column 0
		if grid[5][0] != 31 {
			t.Errorf("May 2026 day 31 at [5][0] = %d, want 31", grid[5][0])
		}
		// After day 31, next cells should be 0
		if grid[5][1] != 0 {
			t.Errorf("May 2026 [5][1] = %d, want 0 (next month fill)", grid[5][1])
		}
	})

	t.Run("January 2026 has 31 days starting Thursday", func(t *testing.T) {
		// Jan 2026: 1st is a Thursday (weekday=4), 31 days
		grid := monthGrid(2026, time.January)

		// Count total non-zero entries (all days of January)
		dayCount := 0
		for row := 0; row < 6; row++ {
			for col := 0; col < 7; col++ {
				if grid[row][col] != 0 {
					dayCount++
				}
			}
		}
		if dayCount != 31 {
			t.Errorf("January 2026 has %d non-zero days, want 31", dayCount)
		}

		// Thursday start, so [0][4] = 1
		if grid[0][4] != 1 {
			t.Errorf("January 2026 day 1 at [0][4] = %d, want 1", grid[0][4])
		}
	})

	t.Run("September 2025 starts on Monday", func(t *testing.T) {
		// Sep 2025: 1st is a Monday (weekday=1), 30 days
		grid := monthGrid(2025, time.September)

		// Monday = 1, so day 1 is at column 1
		if grid[0][1] != 1 {
			t.Errorf("September 2025 day 1 at [0][1] = %d, want 1", grid[0][1])
		}
		// Sunday column (0) should be 0 (prev month fill)
		if grid[0][0] != 0 {
			t.Errorf("September 2025 [0][0] = %d, want 0", grid[0][0])
		}
	})

	t.Run("February 2024 is a leap year with 29 days", func(t *testing.T) {
		grid := monthGrid(2024, time.February)

		dayCount := 0
		for row := 0; row < 6; row++ {
			for col := 0; col < 7; col++ {
				if grid[row][col] != 0 {
					dayCount++
				}
			}
		}
		if dayCount != 29 {
			t.Errorf("February 2024 has %d non-zero days, want 29 (leap year)", dayCount)
		}
	})

	t.Run("year wraps from December to January", func(t *testing.T) {
		// Verify that December 2025 and January 2026 are handled independently
		decGrid := monthGrid(2025, time.December)
		janGrid := monthGrid(2026, time.January)

		// Dec 2025: 31 days, 1st is Monday (weekday=1)
		decDays := countNonZero(decGrid)
		if decDays != 31 {
			t.Errorf("December 2025 has %d days, want 31", decDays)
		}

		// Jan 2026: 31 days
		janDays := countNonZero(janGrid)
		if janDays != 31 {
			t.Errorf("January 2026 has %d days, want 31", janDays)
		}
	})

	t.Run("Monday start offsets correctly", func(t *testing.T) {
		// June 2026: 1st is a Monday (weekday=1)
		grid := monthGrid(2026, time.June)

		// Day 1 at column 1 (Monday)
		if grid[0][1] != 1 {
			t.Errorf("June 2026 day 1 at [0][1] = %d, want 1", grid[0][1])
		}
		// Sunday column (0) should be 0 (prev month)
		if grid[0][0] != 0 {
			t.Errorf("June 2026 [0][0] = %d, want 0", grid[0][0])
		}
		// Day 30 should be at end of the grid
		dayCount := countNonZero(grid)
		if dayCount != 30 {
			t.Errorf("June 2026 has %d days, want 30", dayCount)
		}
	})

	t.Run("Saturday start uses only 1 column in first row", func(t *testing.T) {
		// August 2026: 1st is a Saturday (weekday=6)
		grid := monthGrid(2026, time.August)

		// Saturday = 6, so day 1 is at column 6
		if grid[0][6] != 1 {
			t.Errorf("August 2026 day 1 at [0][6] = %d, want 1", grid[0][6])
		}
		// Columns 0-5 should be 0 (prev month fill)
		for col := 0; col < 6; col++ {
			if grid[0][col] != 0 {
				t.Errorf("August 2026 [0][%d] = %d, want 0", col, grid[0][col])
			}
		}
	})

	t.Run("Sunday start fills no leading zeros", func(t *testing.T) {
		// Need a month that starts on Sunday. Let's find one.
		// March 2026: 1st is Sunday (weekday=0)
		// Let me verify... Actually, let me test the general Sunday-start case
		// by picking a known Sunday month.
		// December 2024: 1st is Sunday
		grid := monthGrid(2024, time.December)

		// Sunday = 0, so day 1 is at column 0
		if grid[0][0] != 1 {
			t.Errorf("December 2024 day 1 at [0][0] = %d, want 1", grid[0][0])
		}
		// No prev-month fill needed
		dayCount := countNonZero(grid)
		if dayCount != 31 {
			t.Errorf("December 2024 has %d days, want 31", dayCount)
		}
	})
}

func TestNewModel(t *testing.T) {
	t.Run("initializes to current month and year", func(t *testing.T) {
		theme := DefaultTheme()
		m := NewModel(theme)

		now := time.Now()
		if m.year != now.Year() {
			t.Errorf("Model year = %d, want %d", m.year, now.Year())
		}
		if m.month != now.Month() {
			t.Errorf("Model month = %d, want %d", m.month, now.Month())
		}
	})

	t.Run("sets theme and fixed terminal dimensions", func(t *testing.T) {
		theme := ThemeColors{
			Foreground: "#aabbcc",
			Background: "#112233",
			Accent:     "#445566",
			Color0:     "#778899",
			Color3:     "#123456",
		}
		m := NewModel(theme)

		if m.theme != theme {
			t.Errorf("Model theme = %+v, want %+v", m.theme, theme)
		}
		if m.width != TerminalWidth {
			t.Errorf("Model width = %d, want %d", m.width, TerminalWidth)
		}
		if m.height != TerminalHeight {
			t.Errorf("Model height = %d, want %d", m.height, TerminalHeight)
		}
	})
}

// countNonZero counts the number of non-zero entries in a grid.
func countNonZero(grid [6][7]int) int {
	count := 0
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			if grid[row][col] != 0 {
				count++
			}
		}
	}
	return count
}
