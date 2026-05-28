package main

import (
	"fmt"
	"strings"
)

// View renders the calendar TUI as a string.
// Layout: header (centered month+year) + blank line + day names + 6 grid rows.
// Follows Bubbletea golden rules: borderless (no border subtraction),
// fixed cell width (no auto-wrap), and theme-driven styling.
func (m Model) View() string {
	styles := NewStyles(m.theme)
	grid := monthGrid(m.year, m.month)

	// Header: centered month name + year
	header := styles.Header.Render(fmt.Sprintf("%s %d", m.month.String(), m.year))

	// Day names row: Su Mo Tu We Th Fr Sa
	dayNames := styles.DayName.Render("Su Mo Tu We Th Fr Sa")

	// Grid rows: 6 rows of 7 cells, each cell 2 chars wide, right-aligned
	var gridRows []string
	for row := 0; row < 6; row++ {
		var cells []string
		for col := 0; col < 7; col++ {
			day := grid[row][col]
			if day == 0 {
				// Empty cell (prev/next month day)
				cells = append(cells, "  ")
			} else {
				dayStr := fmt.Sprintf("%2d", day)
				if m.isToday(day) {
					cells = append(cells, styles.TodayDay.Render(dayStr))
				} else {
					cells = append(cells, styles.NormalDay.Render(dayStr))
				}
			}
		}
		gridRows = append(gridRows, strings.Join(cells, " "))
	}

	// Assemble the view: header + blank line + day names + grid rows
	var lines []string
	lines = append(lines, header)
	lines = append(lines, "") // blank line between header and day names
	lines = append(lines, dayNames)
	lines = append(lines, gridRows...)

	return strings.Join(lines, "\n") + "\n"
}

// isToday returns true if the given day matches today's date in the current view month/year.
func (m Model) isToday(day int) bool {
	return m.today.Year() == m.year && m.today.Month() == m.month && m.today.Day() == day
}
