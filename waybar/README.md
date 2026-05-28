# Waybar Custom

A polished waybar setup for [Omarchy](https://omarchy.org/) with custom modules.

## Features

- **📅 Calendar TUI** — Right-click the clock to open a month-view calendar built with Go + Bubbletea. Navigate with arrow keys or mouse.
- **🎵 Media Player** — Shows current track from any MPRIS player. Left-click to toggle play/pause, right-click to cycle players. Auto-switches when the current player stops.
- **🎨 Omarchy-themed** — Uses your current theme colors automatically.

## Requirements

- [Omarchy](https://omarchy.org/) Linux distribution
- **Go 1.21+** (for calendar TUI compilation)
- **jq** (for calendar monitor detection)
- **Alacritty** or **Kitty** terminal emulator

## Quick Install

```bash
git clone https://github.com/Bard-o/my_omarchy_config.git
cd my_omarchy_config/waybar
./install.sh
omarchy restart waybar
hyprctl reload
```

The install script will:
1. Back up your existing waybar configs
2. Symlink the custom configs to `~/.config/waybar/`
3. Build and install the calendar binary
4. Show you what to add to your hyprland config

## Files

| File | Description |
|---|---|
| `config.jsonc` | Waybar layout with custom modules (clock → calendar, media player) |
| `style.css` | Waybar styling (imports current omarchy theme) |
| `media.sh` | Multi-player MPRIS media indicator script |
| `calendar/` | Go + Bubbletea month-view calendar TUI |
| `hyprland.conf` | Window rules for the calendar floating overlay |
| `install.sh` | One-command setup |

## Calendar TUI

- **Open**: Right-click the clock in waybar
- **Navigate**: ← → or `h` `l` for previous/next month
- **Mouse**: Click left half of header for previous month, right half for next
- **Close**: `q`, `Esc`, or `Ctrl+C`
- **Style**: Automatically uses your omarchy theme colors

## License

MIT — see [LICENSE](../LICENSE)
