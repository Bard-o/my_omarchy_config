package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Frame dimensions (in characters).
const frameWidth = 24

// Border characters (multi-byte UTF-8 box drawing).
const (
	borderTL     = "┌" // top-left
	borderTR     = "┐" // top-right
	borderBL     = "└" // bottom-left
	borderBR     = "┘" // bottom-right
	borderH      = "─" // horizontal
	borderV      = "│" // vertical
	borderPrefix = "─ " // top border: ┌─ title
	borderSuffix = ""    // top border: ┐ with no extra before

	frameInnerW = frameWidth - 2 // width inside vertical borders (chars)
)

// View renders the calendar TUI as a string inside a bordered frame.
// Top border has the month/year embedded like LazyGit/Btop/Impala:
//
//	┌─ May 2026 ──────────┐
//	│                      │
//	│  Su Mo Tu We Th Fr Sa│
//	│                1  2  │
//	│  ...                 │
//	│                      │
//	│  ← → navegar  q salir│
//	└──────────────────────┘
func (m Model) View() string {
	styles := NewStyles(m.theme)
	grid := monthGrid(m.year, m.month)

	// ▸ Top border with embedded title: ┌─ May 2026 ──────────┐
	plainTitle := fmt.Sprintf(" %s %d ", m.month.String(), m.year)
	title := styles.Header.Render(plainTitle)
	titleRuneLen := utf8.RuneCountInString(plainTitle)
	// "┌─ " = 3 chars, then title, then dashes, then "┐" = 1 char → total 24
	dashCount := frameWidth - 3 - titleRuneLen - 1
	top := styles.Frame.Render(borderTL+borderPrefix) + title +
		styles.Frame.Render(strings.Repeat(borderH, dashCount)+borderTR)

	// ▸ Day names row (already styled by Lipgloss)
	dayNames := styles.DayName.Render("Su Mo Tu We Th Fr Sa")

	// ▸ Grid rows: 6 rows × 7 cells, each 2 chars wide
	var gridRows []string
	for row := 0; row < 6; row++ {
		var cells []string
		for col := 0; col < 7; col++ {
			day := grid[row][col]
			if day == 0 {
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

	// ▸ Footer
	footer := styles.Hint.Render("← → navegar  q salir")

	// ▸ Bottom border
	bottom := styles.Frame.Render(borderBL + strings.Repeat(borderH, frameWidth-2) + borderBR)

	// ▸ Assemble: each content line is padded to contentWidth (20) chars
	// Layout: │ + space + content(20) + space + │ = 24 total
	const contentWidth = 20
	bl := styles.Frame.Render(borderV + " ") // "│ "
	br := styles.Frame.Render(" " + borderV)  // " │"
	blankLine := bl + strings.Repeat(" ", contentWidth) + br

	// Footer: strip ANSI, count runes for visual width, pad
	footerStripped := stripStyles(footer)
	footerPad := contentWidth - utf8.RuneCountInString(footerStripped)
	if footerPad < 0 {
		footerPad = 0
	}

	var lines []string
	lines = append(lines, top)
	lines = append(lines, blankLine)
	lines = append(lines, bl+dayNames+br)
	for _, row := range gridRows {
		lines = append(lines, bl+row+br)
	}
	lines = append(lines, bl+footerStripped+strings.Repeat(" ", footerPad)+br)
	lines = append(lines, bottom)

	return strings.Join(lines, "\n") + "\n"
}

// stripStyles removes Lipgloss ANSI escape sequences from a string.
func stripStyles(s string) string {
	var result strings.Builder
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
		} else if inEscape {
			if r == 'm' {
				inEscape = false
			}
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// isToday returns true if the given day matches today's date in the current view month/year.
func (m Model) isToday(day int) bool {
	return m.today.Year() == m.year && m.today.Month() == m.month && m.today.Day() == day
}
