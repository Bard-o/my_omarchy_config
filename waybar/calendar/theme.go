package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// LoadTheme reads the Omarchy theme colors.toml and returns ThemeColors.
// Falls back to Gruvbox defaults on missing or malformed file.
func LoadTheme() ThemeColors {
	home, err := os.UserHomeDir()
	if err != nil {
		return DefaultTheme()
	}
	dir := filepath.Join(home, ".config", "omarchy", "current", "theme")
	return LoadThemeFromDir(dir)
}

// LoadThemeFromDir reads colors.toml from the given directory and returns ThemeColors.
// Missing fields are filled with Gruvbox defaults. Malformed files return full defaults.
func LoadThemeFromDir(dir string) ThemeColors {
	path := filepath.Join(dir, "colors.toml")
	data, err := os.ReadFile(path)
	if err != nil {
		return DefaultTheme()
	}

	result := DefaultTheme()
	parsed := parseColorValues(string(data))

	if v, ok := parsed["foreground"]; ok {
		result.Foreground = v
	}
	if v, ok := parsed["background"]; ok {
		result.Background = v
	}
	if v, ok := parsed["accent"]; ok {
		result.Accent = v
	}
	if v, ok := parsed["color0"]; ok {
		result.Color0 = v
	}
	if v, ok := parsed["color3"]; ok {
		result.Color3 = v
	}

	return result
}

// colorPattern matches lines like: key = "#rrggbb"
// Group 1: key name, Group 2: hex color value (with #)
var colorPattern = regexp.MustCompile(`^(\w+)\s*=\s*"(#[0-9a-fA-F]{6})"`)

// parseColorValues extracts key=value pairs from colors.toml content.
// Each line should match: key = "#rrggbb"
func parseColorValues(content string) map[string]string {
	result := make(map[string]string)
	lines := splitLines(content)
	for _, line := range lines {
		matches := colorPattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			result[matches[1]] = matches[2]
		}
	}
	return result
}

// splitLines splits content into non-empty lines, trimming \r for Windows line endings.
func splitLines(content string) []string {
	var lines []string
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
