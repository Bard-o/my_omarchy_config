show_system_menu() {
  local options="󱄄  Screensaver\n  Lock"
  ! omarchy-toggle-enabled suspend-off && options="$options\n󰒲  Suspend"
  omarchy-hibernation-available && options="$options\n󰤁  Hibernate"
  options="$options\n󰍃  Logout\n󰜉  Restart\n󰐥  Shutdown"

  case $(menu "System" "$options") in
  *Screensaver*) omarchy-launch-screensaver force ;;
  *Lock*) omarchy-system-lock ;;
  *Suspend*) systemctl suspend ;;
  *Hibernate*) systemctl hibernate ;;
  *Logout*) omarchy-system-logout ;;
  *Restart*) omarchy-system-reboot ;;
  *Shutdown*) omarchy-system-shutdown ;;
  *) back_to show_main_menu ;;
  esac
}

show_style_menu() {
  case $(menu "Style" "󰸌  Theme\n󰅣  Avatar\n󰟵  Unlock\n  Font\n  Background\n  Hyprland\n󱄄  Screensaver\n  About") in
  *Theme*) show_theme_menu ;;
  *Avatar*) omarchy-launch-walker -m menus:omarchyProfileSelector --width 500 --minheight 300 ;;
  *Unlock*) omarchy-launch-walker -m menus:omarchyunlocks --width 800 --minheight 400 ;;
  *Font*) show_font_menu ;;
  *Background*) show_background_menu ;;
  *Hyprland*) open_in_editor ~/.config/hypr/looknfeel.conf ;;
  *Screensaver*) show_screensaver_menu ;;
  *About*) show_about_menu ;;
  *) show_main_menu ;;
  esac
}
