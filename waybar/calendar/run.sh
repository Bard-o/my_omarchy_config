#!/bin/bash

# waybar-calendar: Launch the calendar TUI as a floating window
# Positioned by Hyprland window rules (~/.config/hypr/hyprland.conf)

# Kill existing calendar window
pkill -x waybar-calendar 2>/dev/null || true

# Build if binary is missing
BIN_DIR="$(cd "$(dirname "$0")" && pwd)"
if [[ ! -x "$BIN_DIR/waybar-calendar" ]]; then
    make -C "$BIN_DIR" build || exit 1
fi

# Launch with detected terminal emulator (uwsm-app integrates with Wayland session)
if command -v kitty &>/dev/null; then
    setsid -f uwsm-app -- kitty --class=waybar-calendar --title=Calendar -e waybar-calendar
elif command -v alacritty &>/dev/null; then
    setsid -f uwsm-app -- alacritty --class=waybar-calendar --title=Calendar -e waybar-calendar
else
    setsid -f xdg-terminal-exec --app-id=waybar-calendar -e waybar-calendar
fi
