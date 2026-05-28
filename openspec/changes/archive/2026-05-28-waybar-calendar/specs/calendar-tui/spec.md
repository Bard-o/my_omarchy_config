# Calendar TUI Specification

## Purpose

A standalone Go+Bubbletea TUI application that renders a month-view calendar as a floating overlay, triggered by right-clicking the waybar clock module. Displays the current month with correct day alignment, today highlighting, and month navigation.

## Requirements

### Requirement: Calendar Display

The system MUST render a 7-column month grid with correct weekday alignment, current day highlighting, and a header showing month name and year.

| Field | Value |
|-------|-------|
| Grid | 7 columns (Mon–Sun), up to 6 rows |
| Header | Month name + year, centered |
| Weekday labels | Abbreviated (Mo, Tu, We, Th, Fr, Sa, Su) |
| Today | Visually distinct (bold + inverted or colored background) |
| Empty cells | Blank — no placeholder characters |

#### Scenario: Current month renders correctly

- GIVEN the current date is 2026-05-28
- WHEN the application starts
- THEN the header displays "May 2026"
- AND day 28 appears in the correct weekday column (Thursday)
- AND day 28 is visually highlighted as today

#### Scenario: Month with 31 days renders full grid

- GIVEN the current month is January 2026
- WHEN the calendar renders
- THEN days 1–31 appear in correct weekday positions
- AND the grid fills 5 or 6 rows as needed

#### Scenario: Month starting on Monday aligns correctly

- GIVEN September 2025 (starts on Monday)
- WHEN the calendar renders
- THEN day 1 appears in the first column (Monday)

### Requirement: Keyboard Navigation

The system MUST allow month navigation via arrow keys and Vim-style keys (h/l), and exit via Escape or q.

| Key | Action |
|-----|--------|
| Left / h | Previous month |
| Right / l | Next month |
| Up / k | Previous month (alias) |
| Down / j | Next month (alias) |
| Tab | Toggle focus (future: no-op for v1) |
| Escape | Exit application |
| q | Exit application |

#### Scenario: Navigate to next month

- GIVEN the calendar displays May 2026
- WHEN the user presses Right arrow or l
- THEN the calendar displays June 2026
- AND the grid updates with correct day positions

#### Scenario: Navigate to previous month

- GIVEN the calendar displays January 2026
- WHEN the user presses Left arrow or h
- THEN the calendar displays December 2025
- AND the year in the header updates to 2025

#### Scenario: Exit with Escape

- GIVEN the calendar is displaying any month
- WHEN the user presses Escape
- THEN the application exits with code 0

#### Scenario: Exit with q

- GIVEN the calendar is displaying any month
- WHEN the user presses q
- THEN the application exits with code 0

### Requirement: Mouse Navigation

The system MUST support clicking to navigate months. Clicking the left side of the header navigates to the previous month; clicking the right side navigates to the next month.

#### Scenario: Click right side of header advances month

- GIVEN the calendar displays May 2026
- WHEN the user clicks on the right half of the header area
- THEN the calendar displays June 2026

#### Scenario: Click left side of header goes back

- GIVEN the calendar displays May 2026
- WHEN the user clicks on the left half of the header area
- THEN the calendar displays April 2026

### Requirement: Theme Integration

The system MUST read waybar theme colors from the active Omarchy theme CSS and apply them to the TUI styles. It MUST fallback to hardcoded defaults if parsing fails.

| Variable | Source | Fallback |
|----------|--------|----------|
| Background | `@define-color background` in theme CSS | `#282828` |
| Foreground | `@define-color foreground` in theme CSS | `#d4be98` |
| Accent | Derived from foreground (brighter variant) | `#e8d5b5` |

The CSS file path is `~/.config/omarchy/current/theme/waybar.css`.

#### Scenario: Theme colors load successfully

- GIVEN the theme CSS contains `@define-color background #282828` and `@define-color foreground #d4be98`
- WHEN the application starts
- THEN the TUI background is `#282828` and text is `#d4be98`

#### Scenario: Missing CSS falls back to defaults

- GIVEN the theme CSS file does not exist
- WHEN the application starts
- THEN the TUI uses hardcoded fallback colors
- AND no error is displayed to the user

#### Scenario: Malformed CSS falls back to defaults

- GIVEN the theme CSS file exists but contains no `@define-color` declarations
- WHEN the application starts
- THEN the TUI uses hardcoded fallback colors

### Requirement: Window Positioning

The launch script MUST detect the active monitor dimensions via `hyprctl monitors -j` and position the terminal window as a floating overlay near the top-center of the screen.

| Parameter | Value |
|-----------|-------|
| Terminal width | 40 columns |
| Terminal height | 18 rows |
| Horizontal position | Centered on active monitor |
| Vertical position | Near top (y-offset ~10% of monitor height) |
| Multi-monitor | Position on the monitor with focused workspace |

#### Scenario: Single monitor positions correctly

- GIVEN a single monitor at 1920x1080
- WHEN the launch script runs
- THEN the terminal opens at approximately x=740, y=108 (centered horizontally, near top)

#### Scenario: Multi-monitor uses focused display

- GIVEN two monitors, with the focused workspace on the HDMI display
- WHEN the launch script runs
- THEN the terminal positions relative to the HDMI monitor dimensions

### Requirement: Performance

The system MUST start and render the first frame within 200ms with no visible flicker or blank screen.

#### Scenario: Fast startup

- WHEN the user triggers the calendar via waybar right-click
- THEN the calendar is fully rendered within 200ms
- AND no blank terminal frame is visible

#### Scenario: No flicker on month change

- GIVEN the calendar displays May 2026
- WHEN the user navigates to June 2026
- THEN the view updates without flicker or intermediate blank state
