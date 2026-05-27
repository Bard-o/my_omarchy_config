# Video Wallpapers for Omarchy

Animated wallpaper support for Omarchy with video loop, audio mute, and toggle hotkey.

Demo:
https://github.com/user-attachments/assets/5dd8f3f4-8fff-4326-bc1b-c6b13ccfabf3

## Features

- MP4, MKV, WebM video support
- Infinite loop
- Audio muted by default
- Toggle with `SUPER + SHIFT + W`
- Per-theme video folders
- Remembers last used wallpaper
- Also supports static images and animated GIFs

## Installation

```bash
# Clone this repository
git clone https://github.com/Bard-o/my_omarchy_config.git ~/Projects/Omarchy_Setup

# Run installation script
cd ~/Projects/Omarchy_Setup/video-wallpaper
chmod +x install.sh
./install.sh
```

## Requirements

The following tools are required and will be installed:

- [mpvpaper-rs](https://github.com/BitYoungjae/mpvpaper-rs) - Video wallpaper player
- [swww](https://github.com/LGFae/swww) - Animated GIF wallpaper daemon
- [mpv](https://mpv.io/) - Media player (usually pre-installed)
- [elephant](https://github.com/pystardust/eleFant) - Menu system for Omarchy

## Usage

| Action | Hotkey | Description |
|--------|--------|-------------|
| Toggle video/static | `SUPER + SHIFT + W` | Switch between last video and last static wallpaper |
| Select wallpaper | `SUPER + CTRL + SPACE` | Open menu with all backgrounds (images + videos) |
| Cycle backgrounds | Via menu | Select any wallpaper from theme folders |

## File Locations

| Purpose | Path |
|---------|------|
| Videos | `~/.config/omarchy/themes/<theme>/videos/` |
| Static images | `~/.config/omarchy/current/theme/backgrounds/` |
| Scripts | `~/.local/share/omarchy/bin/` |
| Background menu | `~/.config/elephant/menus/` |
| State files | `~/.config/omarchy/current/` |

## How It Works

### Toggle Behavior
The toggle feature remembers your last choices:
- When switching **video → static**: Uses the last static wallpaper you had
- When switching **static → video**: Uses the last video wallpaper you had
- If you select a wallpaper from the menu, it saves as your new "last used"

### Folder Structure
Each Omarchy theme can have its own video folder:
```
~/.config/omarchy/themes/
├── gruvbox/
│   ├── backgrounds/   # Static wallpapers
│   └── videos/       # Video wallpapers
├── another-theme/
│   ├── backgrounds/
│   └── videos/
```

### Supported Formats

| Type | Extensions | Tool Used |
|------|------------|-----------|
| Video | MP4, MKV, WebM, AVI, MOV, WMV | mpvpaper-rs |
| Animated GIF | GIF | swww |
| Static Image | JPG, PNG, WebP, BMP | swaybg |

## Adding Videos

```bash
# Copy your video to the current theme's videos folder
cp ~/Downloads/my-video.mp4 ~/.config/omarchy/themes/gruvbox/videos/

# Then select it from SUPER + CTRL + SPACE menu
# Or use SUPER + SHIFT + W to toggle to it
```

## Uninstallation

To remove video wallpaper support:
```bash
# Remove scripts
rm ~/.local/share/omarchy/bin/omarchy-theme-bg-set
rm ~/.local/share/omarchy/bin/omarchy-theme-bg-next
rm ~/.local/share/omarchy/bin/omarchy-toggle-video-wallpaper

# Restore default background selector
omarchy-refresh-config elephant/background_selector.lua

# Restart services
systemctl --user restart elephant.service
```

## Credits

Created by [Bard-o](https://github.com/Bard-o)
License: MIT

Tools used:
- [mpvpaper-rs](https://github.com/BitYoungjae/mpvpaper-rs) by BitYoungjae
- [swww](https://github.com/LGFae/swww) by LGFae
- [mpv](https://mpv.io/) by mpv team
