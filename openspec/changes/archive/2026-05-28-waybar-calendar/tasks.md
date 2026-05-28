# Tasks: Waybar Calendar TUI

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 600–750 |
| 400-line budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1: engine+tests (~300) → PR 2: UI+launch (~350) |
| Delivery strategy | ask-on-risk |
| Chain strategy | pending |

Decision needed before apply: Yes
Chained PRs recommended: Yes
Chain strategy: pending
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Notes |
|------|------|-----------|-------|
| 1 | Go module, types, calendar grid, theme parser, all unit tests | PR 1 | Self-contained, compiles and tests independently |
| 2 | Bubbletea model/update/view, styles, launch script, Makefile | PR 2 | Depends on PR 1 types; adds TUI loop and deployment |

## Phase 1: Foundation

- [x] 1.1 `waybar-calendar/go.mod` — `go mod init github.com/bardo/waybar-calendar`, Go 1.26.3, add bubbletea/lipgloss/bubbles deps. [code]
- [x] 1.2 `waybar-calendar/types.go` — Define `Model{year,month,today,theme,width,height}`, `ThemeColors{Foreground,Background,Accent,Color0,Color3}`, `DayCell` struct, weekday constants. [code]
- [x] 1.3 `waybar-calendar/theme.go` — `LoadTheme()` reads `~/.config/omarchy/current/theme/colors.toml` via regex, parses key=value pairs, returns `ThemeColors`. Fallback to Gruvbox defaults on missing/malformed file. [code]
- [x] 1.4 `waybar-calendar/theme_test.go` — Table-driven tests: valid TOML parses colors, missing file returns defaults, malformed TOML returns defaults, partial TOML fills gaps with defaults. [test]
- [x] 1.5 `waybar-calendar/model.go` — `NewModel(theme)` initializes to `time.Now()`. `monthGrid(year,month)` returns `[6][7]int` with correct weekday offsets, prev/next month spill as 0. [code]
- [x] 1.6 `waybar-calendar/calendar_test.go` — Table-driven tests: May 2026 grid (Fri), Jan 2026 (Thu), Sep 2025 (Mon), leap year Feb 2024, year wrap Dec→Jan, Monday/Saturday/Sunday starts. [test]
- [x] 1.7 Run `go test ./...` — All Phase 1 tests pass. [test]

## Phase 2: TUI Core

- [x] 2.1 `waybar-calendar/styles.go` — Lipgloss styles: `headerStyle` (accent, centered), `dayNameStyle` (dimmed fg), `normalDayStyle`, `todayDayStyle` (bold+inverted), `emptyDayStyle`. All use `ThemeColors`. [code]
- [x] 2.2 `waybar-calendar/update.go` — Message dispatcher: routes `tea.KeyMsg` to keyboard handler, `tea.MouseMsg` to mouse handler, returns `tea.Quit` on quit signal. [code]
- [x] 2.3 `waybar-calendar/update_keyboard.go` — `handleKeyboard(m, msg)`: Left/h → prev month, Right/l → next month, Up/k → prev month, Down/j → next month, q/Esc → `tea.Quit`. Wrap year at Jan/Dec boundary. [code]
- [x] 2.4 `waybar-calendar/update_mouse.go` — `handleMouse(m, msg)`: click left half of header → prev month, click right half → next month. Only active in header Y range. [code]
- [x] 2.5 `waybar-calendar/view.go` — `View(m)`: renders header (month+year centered), blank line, day-name row (Su–Sa), 6-row day grid. Each cell 2 chars wide, right-aligned. Empty cells are spaces. [code]
- [x] 2.6 `go test ./...` — Add keyboard update tests (h/l/arrows change month, q/Esc returns Quit), mouse tests (click halves), view tests (output contains header text, day numbers, today highlighted). [test]

## Phase 3: Entry Point

- [x] 3.1 `waybar-calendar/main.go` — Parse flags (none for v1), call `LoadTheme()`, create `NewModel()`, enable mouse with `tea.WithMouseCellMotion()`, run `tea.NewProgram()`. ~25 lines. [code]
- [x] 3.2 `go vet ./... && gofmt -w .` — No lint errors, formatted. [config]
- [x] 3.3 `go test ./...` — Full suite passes. [test]

## Phase 4: Build & Launch

- [x] 4.1 `waybar-calendar/Makefile` — Targets: `build` (go build to bin/), `install` (copy to ~/.local/bin/), `test` (go test -cover ./...), `clean`. [config]
- [x] 4.2 `waybar-calendar/run.sh` — Bash script: `pkill -f waybar-calendar` (kill existing), `hyprctl monitors -j` → extract focused monitor x/y/width, calculate center-x and top-y, terminal detection (kitty/alacritty/xdg-terminal-exec), Hyprland window rules for floating position. [config]
- [x] 4.3 `make build && make test` — Binary compiles, all tests pass. [test]
- [x] 4.4 Manual integration: binary exists and starts without panicking. Full TUI verification requires display. [integration]

## Phase 5: Cleanup

- [x] 5.1 Verify `go test ./...` coverage ≥80% on calendar logic (model.go, theme.go). [test]
- [x] 5.2 Confirm no modification to `~/.config/waybar/config.jsonc`. [config]
