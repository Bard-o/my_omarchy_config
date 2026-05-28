package main

import "github.com/charmbracelet/lipgloss"

// Styles holds all Lipgloss styles for the calendar TUI, derived from ThemeColors.
type Styles struct {
	Header    lipgloss.Style
	DayName   lipgloss.Style
	NormalDay lipgloss.Style
	TodayDay  lipgloss.Style
	EmptyDay  lipgloss.Style
	Border    lipgloss.Style
}

// NewStyles creates a Styles struct populated with Lipgloss styles using the given theme colors.
func NewStyles(theme ThemeColors) Styles {
	return Styles{
		Header: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Accent)).
			Width(20).
			Align(lipgloss.Center),
		DayName: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Foreground)),
		NormalDay: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Foreground)),
		TodayDay: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Background)).
			Background(lipgloss.Color(theme.Color3)).
			Bold(true),
		EmptyDay: lipgloss.NewStyle(),
		Border: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Foreground)),
	}
}
