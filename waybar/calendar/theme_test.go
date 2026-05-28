package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadTheme(t *testing.T) {
	t.Run("parses valid colors.toml", func(t *testing.T) {
		dir := t.TempDir()
		content := []byte(
			`accent = "#7daea3"
foreground = "#d4be98"
background = "#282828"
color0 = "#3c3836"
color3 = "#d8a657"
`)
		tomlPath := filepath.Join(dir, "colors.toml")
		if err := os.WriteFile(tomlPath, content, 0644); err != nil {
			t.Fatalf("failed to write test fixture: %v", err)
		}

		theme := LoadThemeFromDir(dir)

		want := ThemeColors{
			Foreground: "#d4be98",
			Background: "#282828",
			Accent:     "#7daea3",
			Color0:     "#3c3836",
			Color3:     "#d8a657",
		}
		if theme != want {
			t.Errorf("LoadThemeFromDir() = %+v, want %+v", theme, want)
		}
	})

	t.Run("returns defaults on missing file", func(t *testing.T) {
		dir := t.TempDir()
		// no colors.toml file created

		theme := LoadThemeFromDir(dir)

		want := DefaultTheme()
		if theme != want {
			t.Errorf("LoadThemeFromDir() with missing file = %+v, want %+v", theme, want)
		}
	})

	t.Run("returns defaults on malformed TOML", func(t *testing.T) {
		dir := t.TempDir()
		content := []byte("this is not valid toml at all {{{}}}")
		tomlPath := filepath.Join(dir, "colors.toml")
		if err := os.WriteFile(tomlPath, content, 0644); err != nil {
			t.Fatalf("failed to write test fixture: %v", err)
		}

		theme := LoadThemeFromDir(dir)

		want := DefaultTheme()
		if theme != want {
			t.Errorf("LoadThemeFromDir() with malformed file = %+v, want %+v", theme, want)
		}
	})

	t.Run("fills gaps with defaults on partial TOML", func(t *testing.T) {
		dir := t.TempDir()
		// Only accent and foreground provided; background, color0, color3 missing
		content := []byte(
			`accent = "#ff0000"
foreground = "#00ff00"
`)
		tomlPath := filepath.Join(dir, "colors.toml")
		if err := os.WriteFile(tomlPath, content, 0644); err != nil {
			t.Fatalf("failed to write test fixture: %v", err)
		}

		theme := LoadThemeFromDir(dir)

		defaults := DefaultTheme()
		if theme.Accent != "#ff0000" {
			t.Errorf("Accent = %q, want %q", theme.Accent, "#ff0000")
		}
		if theme.Foreground != "#00ff00" {
			t.Errorf("Foreground = %q, want %q", theme.Foreground, "#00ff00")
		}
		// Missing fields should fall back to defaults
		if theme.Background != defaults.Background {
			t.Errorf("Background = %q, want default %q", theme.Background, defaults.Background)
		}
		if theme.Color0 != defaults.Color0 {
			t.Errorf("Color0 = %q, want default %q", theme.Color0, defaults.Color0)
		}
		if theme.Color3 != defaults.Color3 {
			t.Errorf("Color3 = %q, want default %q", theme.Color3, defaults.Color3)
		}
	})

	t.Run("ignores comments and unknown keys", func(t *testing.T) {
		dir := t.TempDir()
		// Lines with comments and unknown keys should be skipped
		content := []byte(
			`# This is a comment
accent = "#aabbcc"
cursor = "#ffffff"
foreground = "#112233"
background = "#445566"
some_other_key = "#notvalid"
color0 = "#778899"
color3 = "#123456"
`)
		tomlPath := filepath.Join(dir, "colors.toml")
		if err := os.WriteFile(tomlPath, content, 0644); err != nil {
			t.Fatalf("failed to write test fixture: %v", err)
		}

		theme := LoadThemeFromDir(dir)

		if theme.Accent != "#aabbcc" {
			t.Errorf("Accent = %q, want %q", theme.Accent, "#aabbcc")
		}
		if theme.Foreground != "#112233" {
			t.Errorf("Foreground = %q, want %q", theme.Foreground, "#112233")
		}
		if theme.Background != "#445566" {
			t.Errorf("Background = %q, want %q", theme.Background, "#445566")
		}
		if theme.Color0 != "#778899" {
			t.Errorf("Color0 = %q, want %q", theme.Color0, "#778899")
		}
		if theme.Color3 != "#123456" {
			t.Errorf("Color3 = %q, want %q", theme.Color3, "#123456")
		}
	})

	t.Run("ignores invalid hex values", func(t *testing.T) {
		dir := t.TempDir()
		// Values without proper #hex format should be ignored
		content := []byte(
			`accent = "not-a-color"
foreground = "#d4be98"
background = "#282828"
`)
		tomlPath := filepath.Join(dir, "colors.toml")
		if err := os.WriteFile(tomlPath, content, 0644); err != nil {
			t.Fatalf("failed to write test fixture: %v", err)
		}

		theme := LoadThemeFromDir(dir)

		// accent was invalid, should fall back to default
		defaults := DefaultTheme()
		if theme.Accent != defaults.Accent {
			t.Errorf("Accent = %q, want default %q (invalid hex ignored)", theme.Accent, defaults.Accent)
		}
		if theme.Foreground != "#d4be98" {
			t.Errorf("Foreground = %q, want %q", theme.Foreground, "#d4be98")
		}
	})
}
