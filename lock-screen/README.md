# Lock Screen

A polished custom lock screen for [Omarchy](https://omarchy.org) with a circular avatar, username display, clock, and date — plus an interactive avatar selector.

## Features

- **Circular avatar** with theme-colored border
- **Username display** on the lock screen
- **Clock and date** in the bottom-right corner
- **Dark overlay** on the wallpaper (configurable brightness)
- **Interactive avatar selector** via Elephant/Walker menu
- **Sample avatar included** (smoki)

## Screenshots

> TODO: Add screenshot of the lock screen

## Installation

```bash
git clone https://github.com/Bard-o/my_omarchy_config.git
cd my_omarchy_config/lock-screen
./install.sh
```

The installer will ask for the username to display on the lock screen.

## Usage

### Change Avatar

1. Open **Elephant** (Super + Space)
2. Go to **Style → Avatar**
3. Click an image to select it as your profile picture

### Add More Profile Images

Place images in `~/.config/omarchy/profile_images/`. Supported formats: JPG, PNG, GIF, BMP, WEBP.

The avatar selector will automatically pick up any new images you add to that folder.

### Customize the Lock Screen

Edit `~/.config/hypr/hyprlock.conf`:

```ini
# Avatar size (pixels)
size = 200

# Avatar border thickness (pixels)
border_size = 5

# Avatar position relative to center (px above center)
position = 0, 120

# Username font size
font_size = 25

# Dark overlay brightness (0.0 to 1.0, lower = darker)
brightness = 0.55

# Clock font size
font_size = 80

# Date font size
font_size = 22
```

For full documentation of all available options, see the [hyprlock wiki](https://wiki.hyprland.org/Hypr-Ecosystem/hyprlock).

### Display Timeout After Lock

By default, after locking the screen, the display turns off after **20 seconds** (this is an Omarchy system setting, not part of this config). This gives you time to appreciate your custom lock screen before it goes dark.

To change this delay:

```bash
# Edit the system lock script
nano ~/.local/share/omarchy/bin/omarchy-system-lock
```

Find this line near the bottom:

```bash
sleep 20
```

Change `20` to your preferred number of seconds (for example, `60` for one minute, or `0` to disable automatic display off entirely).

After editing, save and you're done — no restart needed.

> **Note:** This file is part of the Omarchy system and may be overwritten on updates. If that happens, simply edit it again.

## Files

| File | Destination | Description |
|------|-------------|-------------|
| `hyprlock.conf` | `~/.config/hypr/hyprlock.conf` | Lock screen config |
| `elephant/omarchy_profile_selector.lua` | `~/.config/elephant/menus/` | Avatar selector provider |
| `menu.sh` | `~/.config/omarchy/extensions/menu.sh` | Menu overrides (Style + Avatar, System) |
| `profile_images/smoki.jpg` | `~/.config/omarchy/profile_images/` + `~/.face` | Sample avatar |

## Requirements

- [Omarchy](https://omarchy.org)
- Hyprland
- Walker (for the avatar selector menu)
- Elephant (for the menu provider system)

## Uninstall

To restore your previous config:

```bash
# Restore from backup (check the timestamped backup folder)
cp ~/.config/omarchy/lock-screen-backup.*/hyprlock.conf ~/.config/hypr/
cp ~/.config/omarchy/lock-screen-backup.*/menu.sh ~/.config/omarchy/extensions/
cp ~/.config/omarchy/lock-screen-backup.*/omarchy_profile_selector.lua ~/.config/elephant/menus/

# Restart services
omarchy restart walker
hyprctl reload
```

## Credits

Built on [Omarchy](https://omarchy.org) with [Hyprlock](https://wiki.hyprland.org/Hypr-Ecosystem/hyprlock) and [Walker](https://github.com/stepanzwolinski/log).
