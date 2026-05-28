package main

import "time"

// NewModel creates a calendar Model initialized to the current month and year.
func NewModel(theme ThemeColors) Model {
	now := time.Now()
	return Model{
		year:   now.Year(),
		month:  now.Month(),
		today:  now,
		theme:  theme,
		width:  TerminalWidth,
		height: TerminalHeight,
	}
}

// monthGrid returns a 6×7 grid of day numbers for the given year and month.
// Days from the previous or next month are represented as 0.
// The grid always has 6 rows to maintain consistent height.
// Columns represent weekdays with Sunday=0 through Saturday=6.
func monthGrid(year int, month time.Month) [6][7]int {
	var grid [6][7]int

	// First day of the month: which weekday does it fall on?
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	weekday := int(firstDay.Weekday()) // Sunday=0, Saturday=6

	// Number of days in this month
	daysInMonth := daysIn(year, month)

	// Fill the grid
	day := 1
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			if row == 0 && col < weekday {
				// Before the first day of the month
				grid[row][col] = 0
			} else if day > daysInMonth {
				// After the last day of the month
				grid[row][col] = 0
			} else {
				grid[row][col] = day
				day++
			}
		}
	}

	return grid
}

// daysIn returns the number of days in the given month of the given year.
func daysIn(year int, month time.Month) int {
	// time.Date with day=0 of the next month gives the last day of the current month
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
