# Proposal: Waybar Calendar TUI

## Intent

Replace the current `omarchy-tz-select` (gum filter in floating terminal) with a native Go+Bubbletea TUI calendar popup triggered via waybar right-click on the clock module. The current UX is a terminal-based timezone selector — this change delivers a proper styled calendar overlay that matches the Omarchy waybar aesthetic and provides month navigation at a glance.

## Scope

### In Scope
- Go module at `waybar-calendar/` with its own `go.mod`
- Month-view calendar TUI with keyboard navigation (arrows for month, Esc/q to quit)
- Lipgloss styling reading waybar theme colors (background/foreground from CSS variables)
- Floating overlay positioning near top-center via `hyprctl` window rules
- Launch script (`waybar-calendar/run.sh`) for the waybar `on-click-right` action
- Unit tests following strict TDD (`go test ./...` must pass)
- Build/install script to compile binary to `~/.local/bin/`

### Out of Scope
- Modifying waybar config (explicitly deferred — user approval required)
- Calendar event integration (CalDAV, Google Calendar, etc.)
- Year view or week view (month view only for v1)
- Timezone selection functionality (the tz-select use case is replaced, not replicated)
- Editing anything in `~/.local/share/omarchy/`

## Capabilities

### New Capabilities
- `calendar-tui`: Month-view calendar with keyboard navigation, today highlighting, and styled rendering via Bubbletea/Lipgloss

### Modified Capabilities
_None_ — this is a new standalone component with no existing specs to modify.

## Approach

1. **Module scaffolding**: Initialize `waybar-calendar/go.mod` with dependencies (`bubbletea`, `lipgloss`, `bubbles`). Follow the Bubbletea skill file architecture: `main.go`, `types.go`, `model.go`, `update.go`, `update_keyboard.go`, `view.go`, `styles.go`

2. **Calendar rendering**: Render a month grid (7 columns × 6 rows max) with day numbers, aligned to correct weekday offsets. Highlight today. Show month/year header with navigation hints.

3. **Theme integration**: Read waybar's active theme CSS (`~/.config/omarchy/current/theme/waybar.css` or the imported path) to extract `@background` and `@foreground` color variables for Lipgloss styles. Fallback to sensible defaults if parsing fails.

4. **Overlay positioning**: The TUI launches as a small terminal window. Use a Hyprland window rule (via `hyprctl`) to position it as a floating overlay near the top-center of the screen. The launch script (`run.sh`) will:
   - Determine screen dimensions via `hyprctl monitors`
   - Calculate position for a centered floating window near the top
   - Launch the compiled binary in a sized/positioned terminal

5. **Keyboard controls**: Left/Right arrows or `h`/`l` for month navigation, `Escape`/`q` to quit. Today highlighted distinctly.

6. **4 Golden Rules compliance**: All layout calculations follow the Bubbletea skill rules — account for borders, truncate text (prevent auto-wrap), match mouse to layout (not needed for this simple view, but pattern followed), use weights for sizing.

7. **TDD workflow**: Write tests first for date calculations (month grid, weekday offsets, today detection), then for model updates (navigation, quit), then for view rendering.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `waybar-calendar/` | New | Entire Go module — new directory |
| `~/.local/bin/waybar-calendar` | New | Compiled binary (install target) |
| `~/.config/waybar/config.jsonc` | Deferred | `on-click-right` change requires explicit user approval |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Terminal emulator differences affect rendering | Medium | Use `TERM=xterm-256color`; test with default Omarchy terminal (foot/alacritty) |
| Theme CSS parsing fails on non-standard formats | Low | Robust regex parser + hard-coded fallback palette |
| `hyprctl` positioning doesn't work on multi-monitor setups | Medium | Query active monitor from `hyprctl monitors -j`; position relative to active display |
| Waybar config change breaks existing clock behavior | Low | Config change is explicitly out of scope; user must approve separately |

## Rollback Plan

1. Remove `waybar-calendar/` directory
2. Remove compiled binary from `~/.local/bin/waybar-calendar`
3. Waybar config is never modified, so no rollback needed there
4. If user manually changed `on-click-right`, revert to: `omarchy-launch-floating-terminal-with-presentation omarchy-tz-select`

## Dependencies

- Go 1.26.3 (available via mise)
- `hyprctl` (available on Omarchy/Hyprland systems)
- Waybar theme CSS at `~/.config/omarchy/current/theme/waybar.css` (imported by waybar style)
- Bubbletea, Lipgloss, Bubbles (Go modules, pulled at build time)

## Success Criteria

- [ ] `go test ./...` passes with 80%+ coverage on calendar logic
- [ ] Calendar renders current month with correct day layout
- [ ] Keyboard navigation switches months without visual glitches
- [ ] Today's date is visually distinct
- [ ] Styling matches waybar theme (reads `@background`/`@foreground`)
- [ ] Launch script positions overlay near top-center on single monitor
- [ ] `q`/`Escape` exits cleanly
- [ ] Waybar config is NOT modified during implementation