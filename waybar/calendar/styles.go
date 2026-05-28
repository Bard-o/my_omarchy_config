package main

import "github.com/charmbracelet/lipgloss"

// Styles holds all Lipgloss styles for the calendar TUI, derived from ThemeColors.
type Styles struct {
	Header    lipgloss.Style
	DayName   lipgloss.Style
	NormalDay lipgloss.Style
	TodayDay  lipgloss.Style
	EmptyDay  lipgloss.Style
	Frame     lipgloss.Style // border characters
	Hint      lipgloss.Style // footer hints
}

// NewStyles creates a Styles struct populated with Lipgloss styles using the given theme colors.
func NewStyles(theme ThemeColors) Styles {
	return Styles{
		Header: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Accent)).
			Bold(true),
		DayName: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Foreground)),
		NormalDay: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Foreground)),
		TodayDay: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Background)).
			Background(lipgloss.Color(theme.Color3)).
			Bold(true),
		EmptyDay: lipgloss.NewStyle(),
		Frame: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Foreground)),
		Hint: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Color0)),
	}
}
