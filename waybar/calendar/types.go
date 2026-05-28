package main

import "time"

// Weekday order for calendar grid: Sunday through Saturday.
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// ThemeColors holds the color values parsed from the Omarchy theme.
type ThemeColors struct {
	Foreground string // primary text color (e.g. "#d4be98")
	Background string // background color (e.g. "#282828")
	Accent     string // highlight/accent color (e.g. "#7daea3")
	Color0     string // dark gray — used for other-month days
	Color3     string // yellow — used for today highlight
}

// DefaultTheme returns the hardcoded Gruvbox fallback colors.
func DefaultTheme() ThemeColors {
	return ThemeColors{
		Foreground: "#d4be98",
		Background: "#282828",
		Accent:     "#7daea3",
		Color0:     "#3c3836",
		Color3:     "#d8a657",
	}
}

// DayCell represents a single cell in the calendar grid.
// Day is 0 for days belonging to the previous or next month.
type DayCell struct {
	Day       int
	IsToday   bool
	IsCurrent bool // true if this day belongs to the displayed month
}

// Model is the Bubbletea model for the calendar TUI.
type Model struct {
	year   int
	month  time.Month
	today  time.Time
	theme  ThemeColors
	width  int // terminal width (fixed: 24)
	height int // terminal height (fixed: 10)
}

// Fixed terminal dimensions for the calendar popup.
const (
	TerminalWidth  = 24
	TerminalHeight = 12
	GridRows       = 6
	GridCols       = 7
)
