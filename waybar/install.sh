#!/bin/bash
set -e

# waybar-custom — Install script
# Symlinks custom waybar configs + builds the calendar TUI
# https://github.com/Bard-o/my_omarchy_config

REPO_DIR="$(cd "$(dirname "$0")" && pwd)"
WAYBAR_CONFIG_DIR="$HOME/.config/waybar"
HYPR_CONFIG="$HOME/.config/hypr/hyprland.conf"

echo "==> Installing waybar-custom from: $REPO_DIR"
echo

# ── Back up existing configs ──────────────────────────────────────────
echo "→ Backing up existing waybar configs..."
BACKUP_DIR="$WAYBAR_CONFIG_DIR/backup.$(date +%s)"
mkdir -p "$BACKUP_DIR"
for f in config.jsonc style.css media.sh; do
    [ -f "$WAYBAR_CONFIG_DIR/$f" ] && cp "$WAYBAR_CONFIG_DIR/$f" "$BACKUP_DIR/" && echo "   $f → $BACKUP_DIR/"
done
[ -L "$WAYBAR_CONFIG_DIR/calendar" ] || [ -d "$WAYBAR_CONFIG_DIR/calendar" ] && mv "$WAYBAR_CONFIG_DIR/calendar" "$BACKUP_DIR/calendar" 2>/dev/null && echo "   calendar/ → $BACKUP_DIR/"
echo

# ── Symlink config files ──────────────────────────────────────────────
echo "→ Creating symlinks to $WAYBAR_CONFIG_DIR..."
ln -sf "$REPO_DIR/config.jsonc" "$WAYBAR_CONFIG_DIR/config.jsonc"
ln -sf "$REPO_DIR/style.css"   "$WAYBAR_CONFIG_DIR/style.css"
ln -sf "$REPO_DIR/media.sh"    "$WAYBAR_CONFIG_DIR/media.sh"
chmod +x "$WAYBAR_CONFIG_DIR/media.sh"
echo "   config.jsonc → $WAYBAR_CONFIG_DIR/"
echo "   style.css    → $WAYBAR_CONFIG_DIR/"
echo "   media.sh     → $WAYBAR_CONFIG_DIR/"
echo

# ── Symlink calendar directory ────────────────────────────────────────
echo "→ Linking calendar TUI..."
ln -sfn "$REPO_DIR/calendar" "$WAYBAR_CONFIG_DIR/calendar"
echo "   calendar/ → $WAYBAR_CONFIG_DIR/calendar/"
echo

# ── Build and install calendar binary ─────────────────────────────────
echo "→ Building calendar TUI..."
if command -v go &>/dev/null; then
    make -C "$REPO_DIR/calendar" install
    echo "   Binary installed to ~/.local/bin/waybar-calendar"
else
    echo "   ⚠ Go not found — skipping calendar build."
    echo "     Install Go and run: make -C $REPO_DIR/calendar install"
fi
echo

# ── Hyprland window rules ─────────────────────────────────────────────
echo "→ Hyprland window rules for the calendar:"
echo
echo "   Add this line to your $HYPR_CONFIG:"
echo "     source = $REPO_DIR/hyprland.conf"
echo
echo "   Or append the rules manually (see hyprland.conf in this directory)."
echo

# ── Done ──────────────────────────────────────────────────────────────
echo "==> Installation complete!"
echo
echo "   Restart waybar:   omarchy restart waybar"
echo "   Reload hyprland:  hyprctl reload"
echo
echo "   Right-click the clock in your waybar to open the calendar."
echo
echo "   Backups saved in: $BACKUP_DIR"
