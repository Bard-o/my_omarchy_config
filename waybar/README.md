# Waybar Custom

A polished waybar setup for [Omarchy](https://omarchy.org/) with custom modules. fairly similar to the default Omarchy waybar if you already like it.
Btw it's going to be a shame when Omarchy 4.8 comes with quickshell and this config that took me hours to make will be worth nothing...
## Features

- ** Calendar TUI** — Right-click the clock to open a month-view calendar built with Go + Bubbletea. Navigate with arrow keys or mouse.
- ** Media Player** — Shows current track from any MPRIS player. Left-click to toggle play/pause, right-click to cycle players. Auto-switches when the current player stops.
- ** Omarchy-themed** — Uses your current theme colors automatically :D.

## Requirements

- [Omarchy](https://omarchy.org/) Linux distribution
- **Go 1.21+** (for calendar TUI compilation)
- **jq** (for calendar monitor detection)
- **Alacritty** or **Kitty** terminal emulator (it's the default for Omarchy so don't worry)

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

## Troubleshooting 
If the display zone or the terminal windows size bugs to you, config your window rules by yourself  
```bash
nvim .config/hypr/hyprland.conf
```
For example
```toml
# waybar-calendar: floating overlay near top-center
# floating window properties
windowrule = float on, match:class ^(waybar-calendar)$
# size of the window, in pixels, width and heigth
windowrule = size 260 235, match:class ^(waybar-calendar)$
# initial possition, the first argument is to center (check how the number that subtract is half the window width), second argument is initial distance from the top, incliding waybar
windowrule = move (monitor_w*0.5-130) 37, match:class ^(waybar-calendar)$
```
 

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
