package main

import (
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func TestMain(m *testing.M) {
	// Force truecolor rendering so lipgloss emits ANSI codes in tests
	lipgloss.DefaultRenderer().SetColorProfile(termenv.TrueColor)
	os.Exit(m.Run())
}

func TestNewStyles(t *testing.T) {
	t.Run("header style renders accent-colored centered text", func(t *testing.T) {
		theme := DefaultTheme()
		styles := NewStyles(theme)

		rendered := styles.Header.Render("May 2026")
		plain := "May 2026"

		// Rendered output should contain the original text
		if !strings.Contains(rendered, "May 2026") {
			t.Errorf("Header rendered output should contain text 'May 2026', got: %q", rendered)
		}
		// Rendered output should be longer than plain text (ANSI escape codes added)
		if len(rendered) <= len(plain) {
			t.Errorf("Header rendered output should be longer than plain text, got len=%d, plain=%d", len(rendered), len(plain))
		}
		// Should contain ANSI escape sequence (proves color was applied)
		if !strings.Contains(rendered, "\x1b[") {
			t.Error("Header rendered output should contain ANSI escape sequences")
		}
	})

	t.Run("today style uses bold and distinct background", func(t *testing.T) {
		theme := ThemeColors{
			Foreground: "#d4be98",
			Background: "#282828",
			Accent:     "#7daea3",
			Color0:     "#3c3836",
			Color3:     "#d8a657",
		}
		styles := NewStyles(theme)

		rendered := styles.TodayDay.Render("28")
		plain := "28"

		// Today style should add bold (ANSI code 1)
		if !strings.Contains(rendered, "\x1b[1") {
			t.Errorf("TodayDay should be bold, got: %q", rendered)
		}
		// Rendered output should be longer than plain text
		if len(rendered) <= len(plain) {
			t.Errorf("TodayDay rendered output should be longer than plain text, got len=%d, plain=%d", len(rendered), len(plain))
		}
		// Should contain background escape sequence (48;2; for 24-bit background)
		if !strings.Contains(rendered, "48;2;") {
			t.Errorf("TodayDay should have background color (24-bit), got: %q", rendered)
		}
	})

	t.Run("different themes produce different styled output", func(t *testing.T) {
		theme1 := ThemeColors{
			Foreground: "#111111",
			Background: "#222222",
			Accent:     "#333333",
			Color0:     "#444444",
			Color3:     "#555555",
		}
		theme2 := ThemeColors{
			Foreground: "#aaaaaa",
			Background: "#bbbbbb",
			Accent:     "#cccccc",
			Color0:     "#dddddd",
			Color3:     "#eeeeee",
		}

		styles1 := NewStyles(theme1)
		styles2 := NewStyles(theme2)

		header1 := styles1.Header.Render("Jan 2026")
		header2 := styles2.Header.Render("Jan 2026")

		// Different accent colors should produce different ANSI codes
		if header1 == header2 {
			t.Error("Headers with different accent colors should produce different styled output")
		}
		// Both should contain their text
		if !strings.Contains(header1, "Jan 2026") {
			t.Error("Header1 should contain text 'Jan 2026'")
		}
		if !strings.Contains(header2, "Jan 2026") {
			t.Error("Header2 should contain text 'Jan 2026'")
		}
	})

	t.Run("normal day style uses foreground color", func(t *testing.T) {
		theme := ThemeColors{
			Foreground: "#d4be98",
			Background: "#282828",
			Accent:     "#7daea3",
			Color0:     "#3c3836",
			Color3:     "#d8a657",
		}
		styles := NewStyles(theme)

		rendered := styles.NormalDay.Render("15")
		plain := "15"

		// Should contain ANSI foreground escape sequence (38;2; for 24-bit)
		if !strings.Contains(rendered, "38;2;") {
			t.Errorf("NormalDay should have foreground color (24-bit), got: %q", rendered)
		}
		// Should be longer than plain text
		if len(rendered) <= len(plain) {
			t.Errorf("NormalDay rendered output should be longer than plain text, got len=%d, plain=%d", len(rendered), len(plain))
		}
	})

	t.Run("empty day style adds minimal or no styling", func(t *testing.T) {
		theme := DefaultTheme()
		styles := NewStyles(theme)

		rendered := styles.EmptyDay.Render("  ")

		// EmptyDay should not contain theme color ANSI codes
		// (it may contain a reset code but no foreground/background colors)
		if strings.Contains(rendered, "38;2;") {
			t.Error("EmptyDay should not contain foreground color ANSI codes")
		}
		if strings.Contains(rendered, "48;2;") {
			t.Error("EmptyDay should not contain background color ANSI codes")
		}
	})

	t.Run("header style centers text at width 20", func(t *testing.T) {
		theme := DefaultTheme()
		styles := NewStyles(theme)

		// The Header style should have width 20 for centering
		width := styles.Header.GetWidth()
		if width != 20 {
			t.Errorf("Header width = %d, want 20", width)
		}
		// Verify center alignment
		align := styles.Header.GetAlign()
		if align != lipgloss.Center {
			t.Errorf("Header horizontal alignment = %v, want Center", align)
		}
	})
}
