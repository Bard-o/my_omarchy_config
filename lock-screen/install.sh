#!/bin/bash
set -e

# lock-screen — Interactive install script
# Adds a custom lock screen with avatar selector to Omarchy
# https://github.com/Bard-o/my_omarchy_config

REPO_DIR="$(cd "$(dirname "$0")" && pwd)"
BACKUP_DIR="$HOME/.config/omarchy/lock-screen-backup.$(date +%s)"

echo
echo "==> Lock Screen Installer for Omarchy"
echo "    https://github.com/Bard-o/my_omarchy_config"
echo

# ── Check we're on Omarchy ─────────────────────────────────────────────
if [[ ! -d "$HOME/.config/hypr" ]]; then
    echo "Error: This script is designed for Omarchy."
    echo "       Could not find ~/.config/hypr"
    exit 1
fi

# ── Ask for username ───────────────────────────────────────────────────
echo "This will install a custom lock screen with:"
echo "  - Circular avatar with profile image selector"
echo "  - Username display on the lock screen"
echo "  - Clock and date in the bottom-right corner"
echo "  - Dark overlay on the wallpaper"
echo
read -r -p "Enter the username to display on the lock screen: " username
if [[ -z "$username" ]]; then
    echo "Error: Username cannot be empty."
    exit 1
fi
echo

# ── Backup ─────────────────────────────────────────────────────────────
echo "==> Backing up existing configs to:"
echo "   $BACKUP_DIR"
mkdir -p "$BACKUP_DIR"

[[ -f "$HOME/.config/hypr/hyprlock.conf" ]] && cp "$HOME/.config/hypr/hyprlock.conf" "$BACKUP_DIR/" && echo "   hyprlock.conf → backup"
[[ -f "$HOME/.config/omarchy/extensions/menu.sh" ]] && cp "$HOME/.config/omarchy/extensions/menu.sh" "$BACKUP_DIR/menu.sh" && echo "   menu.sh → backup"
[[ -f "$HOME/.config/elephant/menus/omarchy_profile_selector.lua" ]] && cp "$HOME/.config/elephant/menus/omarchy_profile_selector.lua" "$BACKUP_DIR/omarchy_profile_selector.lua" && echo "   omarchy_profile_selector.lua → backup"
[[ -d "$HOME/.config/omarchy/profile_images" ]] && cp -r "$HOME/.config/omarchy/profile_images" "$BACKUP_DIR/profile_images" && echo "   profile_images/ → backup"
[[ -f "$HOME/.face" ]] && cp "$HOME/.face" "$BACKUP_DIR/.face" && echo "   .face → backup"
echo

# ── Create directories ─────────────────────────────────────────────────
echo "==> Creating directories..."
mkdir -p "$HOME/.config/omarchy/extensions"
mkdir -p "$HOME/.config/elephant/menus"
mkdir -p "$HOME/.config/omarchy/profile_images"
echo

# ── Install hyprlock.conf ──────────────────────────────────────────────
echo "==> Installing hyprlock.conf..."
# Replace USERNAME placeholder with actual username
sed "s/USERNAME/$username/g" "$REPO_DIR/hyprlock.conf" > "$HOME/.config/hypr/hyprlock.conf"
echo "   ~/.config/hypr/hyprlock.conf"

# ── Install menu.sh overrides ───────────────────────────────────────────
echo "==> Installing menu extensions..."

# Read the new functions from the repo's menu.sh
new_style_func=$(sed -n '/^show_style_menu/,/^}/p' "$REPO_DIR/menu.sh")
new_system_func=$(sed -n '/^show_system_menu/,/^}/p' "$REPO_DIR/menu.sh")

# Check if menu.sh already exists and has these functions
if [[ -f "$HOME/.config/omarchy/extensions/menu.sh" ]]; then
    if grep -q "^show_style_menu" "$HOME/.config/omarchy/extensions/menu.sh" 2>/dev/null; then
        echo "   show_style_menu() already exists — skipping"
    else
        echo "$new_style_func" >> "$HOME/.config/omarchy/extensions/menu.sh"
        echo "   Added show_style_menu() to menu.sh"
    fi

    if grep -q "^show_system_menu" "$HOME/.config/omarchy/extensions/menu.sh" 2>/dev/null; then
        echo "   show_system_menu() already exists — skipping"
    else
        echo "$new_system_func" >> "$HOME/.config/omarchy/extensions/menu.sh"
        echo "   Added show_system_menu() to menu.sh"
    fi
else
    cp "$REPO_DIR/menu.sh" "$HOME/.config/omarchy/extensions/menu.sh"
    echo "   ~/.config/omarchy/extensions/menu.sh"
fi
echo

# ── Install elephant provider ──────────────────────────────────────────
echo "==> Installing avatar selector (Elephant provider)..."
cp "$REPO_DIR/elephant/omarchy_profile_selector.lua" "$HOME/.config/elephant/menus/omarchy_profile_selector.lua"
echo "   ~/.config/elephant/menus/omarchy_profile_selector.lua"
echo

# ── Install sample avatar ───────────────────────────────────────────────
echo "==> Installing sample avatar..."
cp "$REPO_DIR/profile_images/smoki.jpg" "$HOME/.config/omarchy/profile_images/"
cp "$REPO_DIR/profile_images/smoki.jpg" "$HOME/.face"
echo "   smoki.jpg → ~/.face (default avatar)"
echo "   smoki.jpg → ~/.config/omarchy/profile_images/"
echo

# ── Restart services ────────────────────────────────────────────────────
echo "==> Restarting services..."
omarchy restart walker 2>/dev/null || true
hyprctl reload 2>/dev/null || true
echo

# ── Done ───────────────────────────────────────────────────────────────
echo "==> Installation complete!"
echo
echo "   Your lock screen now shows: $username"
echo
echo "   To change your avatar:"
echo "     Elephant → Style → Avatar"
echo "     (Click an image to select it as your profile picture)"
echo
echo "   To add more profile images:"
echo "     Place images in ~/.config/omarchy/profile_images/"
echo
echo "   To customize the lock screen:"
echo "     Edit ~/.config/hypr/hyprlock.conf"
echo "     (avatar size, positions, colors, font, dark overlay brightness)"
echo
echo "   To test the lock screen:"
echo "     omarchy system lock"
echo
echo "   Backups saved in: $BACKUP_DIR"
echo
