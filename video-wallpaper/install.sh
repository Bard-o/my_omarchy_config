#!/bin/bash

# Video Wallpaper Installation Script for Omarchy
# Author: Bard-o

set -e

echo "=========================================="
echo "Video Wallpaper Installation for Omarchy"
echo "=========================================="
echo ""

# Backup existing files
echo "Creating backups..."
mkdir -p ~/.config/omarchy/current/backups/$(date +%Y%m%d_%H%M%S)

if [[ -f ~/.local/share/omarchy/bin/omarchy-theme-bg-set ]]; then
    cp ~/.local/share/omarchy/bin/omarchy-theme-bg-set ~/.config/omarchy/current/backups/$(date +%Y%m%d_%H%M%S)/
fi

if [[ -f ~/.local/share/omarchy/bin/omarchy-theme-bg-next ]]; then
    cp ~/.local/share/omarchy/bin/omarchy-theme-bg-next ~/.config/omarchy/current/backups/$(date +%Y%m%d_%H%M%S)/
fi

if [[ -f ~/.local/share/omarchy/bin/omarchy-toggle-video-wallpaper ]]; then
    cp ~/.local/share/omarchy/bin/omarchy-toggle-video-wallpaper ~/.config/omarchy/current/backups/$(date +%Y%m%d_%H%M%S)/
fi

if [[ -f ~/.config/elephant/menus/omarchy_background_selector.lua ]]; then
    cp ~/.config/elephant/menus/omarchy_background_selector.lua ~/.config/omarchy/current/backups/$(date +%Y%m%d_%H%M%S)/
fi

echo "Backups created in ~/.config/omarchy/current/backups/"
echo ""

# Install scripts
echo "Installing scripts..."
cp scripts/omarchy-theme-bg-set ~/.local/share/omarchy/bin/
cp scripts/omarchy-theme-bg-next ~/.local/share/omarchy/bin/
cp scripts/omarchy-toggle-video-wallpaper ~/.local/share/omarchy/bin/
chmod +x ~/.local/share/omarchy/bin/omarchy-theme-bg-set
chmod +x ~/.local/share/omarchy/bin/omarchy-theme-bg-next
chmod +x ~/.local/share/omarchy/bin/omarchy-toggle-video-wallpaper
echo "  ✓ Scripts installed"
echo ""

# Install Elephant menu extension
echo "Installing Elephant menu extension..."
if [[ -f ~/.config/elephant/menus/omarchy_background_selector.lua ]]; then
    rm ~/.config/elephant/menus/omarchy_background_selector.lua
fi
cp scripts/omarchy_background_selector.lua ~/.config/elephant/menus/
echo "  ✓ Menu extension installed"
echo ""

# Install mpv configuration
echo "Installing mpv configuration..."
mkdir -p ~/.config/mpv
cp configs/mpv.conf ~/.config/mpv/mpv.conf
echo "  ✓ mpv config installed"
echo ""

# Restart services
echo "Restarting services..."
systemctl --user restart elephant.service
echo "  ✓ Elephant restarted"
echo ""

# Create videos folder for current theme
THEME_NAME=$(cat ~/.config/omarchy/current/theme.name 2>/dev/null || echo "default")
mkdir -p ~/.config/omarchy/themes/$THEME_NAME/videos
mkdir -p ~/.config/omarchy/themes/$THEME_NAME/backgrounds
echo "  ✓ Created videos folder for theme: $THEME_NAME"
echo ""

echo "=========================================="
echo "Installation complete!"
echo "=========================================="
echo ""
echo "Usage:"
echo "  • SUPER + SHIFT + W  - Toggle video/static wallpaper"
echo "  • SUPER + CTRL + SPACE - Select wallpaper from menu"
echo ""
echo "Add videos to:"
echo "  ~/.config/omarchy/themes/$THEME_NAME/videos/"
echo ""
echo "Enjoy your animated wallpapers! 🎬"