# Design: Waybar Calendar TUI

## Technical Approach

A standalone Go+Bubbletea TUI app that renders a styled month-view calendar in a small floating terminal window. The binary is launched via a bash wrapper script triggered by waybar's clock `on-click-right`. Theme colors are read from `~/.config/omarchy/current/theme/colors.toml` (primary source) with `waybar.css` as fallback. The app is intentionally minimal — no title bar, no status bar, just the calendar grid with keyboard navigation.

## Architecture Decisions

| Decision | Options | Tradeoff | Choice |
|----------|---------|----------|--------|
| Color source | `colors.toml` vs `waybar.css` vs both | `colors.toml` has full palette (accent, color0-15); `waybar.css` only has fg/bg. TOML is structured and parseable with `BurntSushi/toml` or simple regex. CSS requires `@define-color` regex. | **`colors.toml` primary, `waybar.css` fallback** — richer palette, simpler parsing |
| Terminal emulator | Hardcode alacritty vs `xdg-terminal-exec` | Hardcoding is fragile; `xdg-terminal-exec` respects user's default | **`xdg-terminal-exec`** with `--app-id` for window rules |
| Window positioning | Hyprland window rule vs launch-script geometry | Window rules are static; launch script can query active monitor dynamically | **Launch script** queries `hyprctl monitors -j`, passes `--position` to terminal |
| Grid sizing | Fixed terminal size vs dynamic | Calendar is always 7×6 max — fixed size is simpler and predictable | **Fixed 24×10 terminal** (cols×rows) |
| Dependency: TOML parser | `BurntSushi/toml` vs regex | Adding a dependency for 23 lines of simple key=value is overkill | **Regex parser** — `colors.toml` is trivially parseable |
| No borders on calendar | Bordered vs borderless | Borders add 2 lines of height and visual noise for a popup | **Borderless** — cleaner popup aesthetic, simpler height math |

## Data Flow

```
waybar clock right-click
        │
        ▼
   run.sh (bash)
        │
        ├── hyprctl monitors -j  →  active monitor dims
        ├── calculate position   →  x,y for floating window
        │
        ▼
   xdg-terminal-exec --app-id=waybar-calendar \
       --position=x,y --size=24x10 \
       -e waybar-calendar (binary)
        │
        ▼
   main.go
        │
        ├── theme.go: read colors.toml → ThemeColors
        ├── model.go: init Model{month: now, year: now}
        │
        ▼
   Bubbletea loop
        │
        ├── update_keyboard.go: h/l/←/→ nav, q/Esc quit
        ├── view.go: render month grid with Lipgloss styles
        └── styles.go: apply ThemeColors to cells
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `waybar-calendar/go.mod` | Create | Go module: `github.com/bardo/waybar-calendar`, Go 1.26.3, deps: bubbletea, lipgloss |
| `waybar-calendar/main.go` | Create | Entry point (~25 lines): parse theme, create model, run tea.Program |
| `waybar-calendar/types.go` | Create | `Model`, `ThemeColors`, `DayCell` structs, constants |
| `waybar-calendar/model.go` | Create | `NewModel()`, `calculateLayout()`, month grid computation |
| `waybar-calendar/update.go` | Create | Message dispatcher: routes to keyboard handler |
| `waybar-calendar/update_keyboard.go` | Create | `handleKeyboard()`: h/l/arrows for nav, q/Esc for quit |
| `waybar-calendar/view.go` | Create | `View()`: renders header + day-name row + day grid |
| `waybar-calendar/styles.go` | Create | Lipgloss styles: header, dayName, normalDay, todayDay, otherMonthDay |
| `waybar-calendar/theme.go` | Create | `LoadTheme()`: reads `colors.toml`, parses with regex, returns `ThemeColors` |
| `waybar-calendar/theme_test.go` | Create | Tests for theme parsing (with sample TOML content) |
| `waybar-calendar/calendar_test.go` | Create | Tests for grid computation, today detection, month navigation |
| `waybar-calendar/run.sh` | Create | Launch script: queries monitor, calculates position, launches binary |
| `waybar-calendar/Makefile` | Create | `build`, `install`, `test`, `clean` targets |

## Interfaces / Contracts

### ThemeColors

```go
type ThemeColors struct {
    Foreground string // e.g. "#d4be98"
    Background string // e.g. "#282828"
    Accent     string // e.g. "#7daea3"
    Color0     string // dark gray — for other-month days
    Color3     string // yellow — for today highlight
}
```

### Model

```go
type Model struct {
    year     int
    month    time.Month
    today    time.Time
    theme    ThemeColors
    width    int  // terminal width (fixed: 24)
    height   int  // terminal height (fixed: 10)
}
```

### Calendar Grid Computation

```go
// monthGrid returns a 6×7 grid of day numbers.
// Days from prev/next month are 0 (rendered as spaces).
func monthGrid(year int, month time.Month) [6][7]int
```

### Theme Loading

```go
// LoadTheme reads ~/.config/omarchy/current/theme/colors.toml
// Falls back to hardcoded Gruvbox defaults on error.
func LoadTheme() ThemeColors
```

### Launch Script Contract

```bash
# run.sh — called by waybar on-click-right
# 1. Kill existing instance if running (pkill -f waybar-calendar)
# 2. Query: hyprctl monitors -j → extract focused monitor x,y,width
# 3. Calculate: pos_x = monitor_x + (monitor_width/2 - 12*char_width)
#               pos_y = monitor_y + 30  (below waybar's 26px + 4px gap)
# 4. Launch: xdg-terminal-exec --app-id=waybar-calendar \
#            --title=waybar-calendar \
#            -e ~/.local/bin/waybar-calendar
# 5. Hyprland window rule (in hyprctl dispatch): float, size, position
```

## Calendar Layout

```
     May 2026       ← header (centered, accent color)
Su Mo Tu We Th Fr Sa ← day names (foreground, dimmed)
             1  2    ← row 1 (offset by weekday)
 3  4  5  6  7  8  9 ← rows 2-6
10 11 12 13 14 15 16
17 18 19 20 21 22 23
24 25 26 27 28 29 30
31                   ← trailing days rendered as spaces
```

**Cell width**: 3 chars (`"XX "` or `" XX"` right-aligned per cell).
**Total grid width**: 20 chars + 2 padding each side = 24 cols.
**Total height**: 1 (header) + 1 (blank) + 1 (day names) + 6 (rows) + 1 (blank) = 10 rows.

### Golden Rules Compliance

1. **No borders** → no border subtraction needed (borderless design)
2. **No auto-wrap** → all day strings are fixed-width (2 chars), header is truncated to 20 chars
3. **No mouse** → keyboard-only, no mouse detection needed
4. **Fixed size** → terminal is launched at exact dimensions, no weight calculations needed

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | `monthGrid()` — correct offsets, leap years, month boundaries | Table-driven tests |
| Unit | `LoadTheme()` — regex parsing of colors.toml, fallback behavior | Test with embedded sample TOML |
| Unit | `NewModel()` — initializes to current month/year | Verify against `time.Now()` |
| Unit | Keyboard update — h/l/arrows change month, q/Esc returns `tea.Quit` | Test `Update()` with `tea.KeyMsg` |
| Unit | View rendering — output contains expected day numbers and header | String assertion on `View()` output |

## Migration / Rollout

No migration required. This is a new standalone component.

**Rollout steps** (post-implementation, requires user approval):
1. `make install` → compiles binary to `~/.local/bin/waybar-calendar`
2. User manually updates `~/.config/waybar/config.jsonc` clock module:
   ```json
   "on-click-right": "~/.config/waybar/../path/to/run.sh"
   ```
3. `omarchy restart waybar`

## Open Questions

- [ ] Should the launch script use `hyprctl dispatch windowrule` for floating/positioning, or rely on terminal geometry flags? (Depends on which terminal is default — alacritty supports `--position`, foot does not)
- [ ] Should we add a Hyprland window rule in `~/.config/hypr/hyprland.conf` for `waybar-calendar` class (float, no border, pin) instead of doing it in the script?
