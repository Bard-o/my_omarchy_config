#!/bin/bash

# =============================================================================
# Waybar Media Indicator
# =============================================================================
# Displays current media player info and handles cycling between players.
# Supports MPRIS-compatible players via playerctl.
#
# Usage:
#   ./media.sh          - Display current track info (for waybar)
#   ./media.sh cycle    - Cycle to next player
#   ./media.sh toggle   - Play/pause current player
#
# Dependencies:
#   - playerctl (MPRIS control)
#
# Files:
#   - /tmp/waybar-media-current    - Stores current player name
#   - /tmp/waybar-media-just-cycled - Flag to prevent auto-switch after manual cycle
#   - /tmp/waybar-media-prev-status - Tracks previous status for auto-switch detection
# =============================================================================

# -----------------------------------------------------------------------------
# Constants
# -----------------------------------------------------------------------------
IGNORE_PLAYERS="blanket"
STATE_FILE="/tmp/waybar-media-current"
CYCLES_FLAG="/tmp/waybar-media-just-cycled"
PREV_STATUS_FILE="/tmp/waybar-media-prev-status"

# =============================================================================
# State Management
# =============================================================================

get_players() {
    playerctl -l 2>/dev/null | grep -v -i "^$IGNORE_PLAYERS$"
}

get_playing_players() {
    for player in $(get_players); do
        if [[ "$(playerctl -p "$player" status 2>/dev/null)" == "Playing" ]]; then
            echo "$player"
        fi
    done
}

set_current_player() {
    local player="$1"
    echo "$player" > "$STATE_FILE"
}

get_current_player() {
    if [[ -f "$STATE_FILE" ]]; then
        cat "$STATE_FILE"
    fi
}

# =============================================================================
# Player Actions
# =============================================================================

cycle_player() {
    local players
    players=($(get_players))
    local current
    current=$(get_current_player)

    if [[ -z "$current" ]] || [[ ! " ${players[*]} " =~ " ${current} " ]]; then
        current="${players[0]}"
    fi

    local found=false
    for i in "${!players[@]}"; do
        if [[ "${players[$i]}" == "$current" ]]; then
            local next=$(( (i + 1) % ${#players[@]} ))
            set_current_player "${players[$next]}"
            touch "$CYCLES_FLAG"
            found=true
            break
        fi
    done

    if [[ "$found" == "false" ]] && [[ ${#players[@]} -gt 0 ]]; then
        set_current_player "${players[0]}"
    fi
}

toggle_current() {
    local current
    current=$(get_current_player)
    if [[ -n "$current" ]]; then
        playerctl -p "$current" play-pause 2>/dev/null
    fi
}

# =============================================================================
# Utilities
# =============================================================================

# Converts player instance names to display names.
# e.g., "vivaldi.instance12345" -> "vivaldi"
normalize_player_name() {
    local player="$1"
    case "$player" in
        vivaldi.*|vivaldi) echo "vivaldi" ;;
        chromium.*|chromium) echo "chromium" ;;
        brave.*|brave) echo "brave" ;;
        firefox.*|firefox) echo "firefox" ;;
        *) echo "$player" ;;
    esac
}

# Returns playback icon based on status.
get_icon() {
    local status="$1"
    if [[ "$status" == "Playing" ]]; then
        echo "⏸"
    else
        echo "▶"
    fi
}

# Truncates text to max length with ellipsis.
truncate_text() {
    local text="$1"
    local max_len=40

    if [[ ${#text} -gt $max_len ]]; then
        echo "${text:0:$max_len}…"
    else
        echo "$text"
    fi
}

# =============================================================================
# JSON Escaping
# =============================================================================

# Bash-only escaping for simple strings (no special characters expected).
escape_json() {
    local str="$1"
    str="${str//\\/\\\\}"
    str="${str//\"/\\\"}"
    str="${str//$'\n'/\\n}"
    str="${str//$'\r'/\\r}"
    str="${str//$'\t'/\\t}"
    echo -n "$str"
}

# Python-based escaping for complex strings (handles all edge cases).
json_escape() {
    python3 -c "import sys,json; sys.stdout.write(json.dumps(sys.stdin.read(), ensure_ascii=False))"
}

# =============================================================================
# Output Generation
# =============================================================================

# Builds tooltip with all active players.
get_tooltip() {
    local players
    players=($(get_players))
    local current
    current=$(get_current_player)
    local lines=""

    # Current player line
    local current_status
    current_status=$(playerctl -p "$current" status 2>/dev/null)
    local current_icon
    current_icon=$(get_icon "$current_status")
    local current_title
    current_title=$(playerctl -p "$current" metadata --format '{{title}}' 2>/dev/null)
    local current_artist
    current_artist=$(playerctl -p "$current" metadata --format '{{artist}}' 2>/dev/null)
    local current_display
    current_display=$(normalize_player_name "$current")

    if [[ -n "$current_artist" && -n "$current_title" ]]; then
        lines="[$current_display] $current_icon $current_title - $current_artist"
    else
        lines="[$current_display] $current_icon $current_title"
    fi

    # Other players
    for p in "${players[@]}"; do
        if [[ "$p" != "$current" ]]; then
            local p_status
            p_status=$(playerctl -p "$p" status 2>/dev/null)

            if [[ "$p_status" == "Playing" || "$p_status" == "Paused" ]]; then
                local p_icon
                p_icon=$(get_icon "$p_status")
                local p_title
                p_title=$(playerctl -p "$p" metadata --format '{{title}}' 2>/dev/null)
                local p_artist
                p_artist=$(playerctl -p "$p" metadata --format '{{artist}}' 2>/dev/null)
                local p_display
                p_display=$(normalize_player_name "$p")

                if [[ -n "$p_artist" && -n "$p_title" ]]; then
                    lines="$lines"$'\n'"[$p_display] $p_icon $p_title - $p_artist"
                else
                    lines="$lines"$'\n'"[$p_display] $p_icon $p_title"
                fi
            fi
        fi
    done

    echo "$lines" | json_escape
}

# =============================================================================
# Main
# =============================================================================

main() {
    local players
    players=($(get_players))

    if [[ ${#players[@]} -eq 0 ]]; then
        rm -f "$PREV_STATUS_FILE" 2>/dev/null
        echo '{"text": "⏹ No media playing...", "tooltip": "No media players running", "class": "stopped"}'
        exit 0
    fi

    local current
    current=$(get_current_player)

    if [[ -z "$current" ]] || [[ ! " ${players[*]} " =~ " ${current} " ]]; then
        current="${players[0]}"
        set_current_player "$current"
    fi

    local status
    status=$(playerctl -p "$current" status 2>/dev/null)

    # Read previous status for transition detection
    local prev_status=""
    if [[ -f "$PREV_STATUS_FILE" ]]; then
        prev_status=$(cat "$PREV_STATUS_FILE")
    fi

    # Auto-switch logic: Only when player transitions from Playing to Paused/Stopped
    # This triggers when:
    #   - Previous was "Playing"
    #   - Current is NOT "Playing"
    #   - Not protected by cycle flag
    local should_auto_switch=false
    if [[ "$prev_status" == "Playing" && "$status" != "Playing" && ! -f "$CYCLES_FLAG" ]]; then
        should_auto_switch=true
    fi

    if [[ "$should_auto_switch" == "true" ]]; then
        local playing
        playing=$(get_playing_players | head -1)
        if [[ -n "$playing" ]]; then
            set_current_player "$playing"
            current="$playing"
            status=$(playerctl -p "$current" status 2>/dev/null)
        fi
    fi

    # Clear cycle flag
    if [[ -f "$CYCLES_FLAG" ]]; then
        rm -f "$CYCLES_FLAG" &
    fi

    if [[ -z "$current" ]]; then
        echo '{"text": "⏹ No media playing...", "tooltip": "No media playing", "class": "stopped"}'
        exit 0
    fi

    # Store current status for next interval
    echo "$status" > "$PREV_STATUS_FILE"

    # Build display text
    local title
    local artist
    title=$(playerctl -p "$current" metadata --format '{{title}}' 2>/dev/null)
    artist=$(playerctl -p "$current" metadata --format '{{artist}}' 2>/dev/null)
    local display_player
    display_player=$(normalize_player_name "$current")
    local icon
    icon=$(get_icon "$status")
    local class
    class=$([ "$status" == "Playing" ] && echo "playing" || echo "paused")

    local full_text
    if [[ -n "$artist" && -n "$title" ]]; then
        full_text="$title - $artist ($display_player)"
    else
        full_text="$title ($display_player)"
    fi

    local display_text
    display_text=$(truncate_text "$full_text")
    local full_text_with_icon="${icon} ${display_text}"

    # Build JSON output
    local escaped_text
    escaped_text=$(escape_json "$full_text_with_icon")
    local escaped_tooltip
    escaped_tooltip=$(get_tooltip)

    echo "{\"text\": \"$escaped_text\", \"tooltip\": $escaped_tooltip, \"class\": \"$class\"}"
}

# =============================================================================
# CLI Handler
# =============================================================================

case "$1" in
    cycle)
        cycle_player
        echo '{}'
        ;;
    toggle)
        toggle_current
        echo '{}'
        ;;
    *)
        main
        ;;
esac