package system

import (
	"github.com/hairyhenderson/go-which"
)

type external_app string

const (
	FEH           external_app = "feh"
	PICOM         external_app = "picom"
	XSET          external_app = "xset"
	XSS_LOCK      external_app = "xss-lock"
	XSECURELOCK   external_app = "xsecurelock"
	ACPI          external_app = "acpi"
	TMUX          external_app = "tmux"
	TERMINAL      external_app = "alacritty"
	BRIGHTNESSCTL external_app = "brightnessctl"
	XRANDR        external_app = "xrandr"
	SETXKBMAP     external_app = "setxkbmap"
	PULSE_CTL     external_app = "pactl"
	BLUETOOTHCTL  external_app = "bluetoothctl"
	IWCTL         external_app = "iwctl"
	XDG_OPEN      external_app = "xdg-open"
	UDISKIE       external_app = "udiskie"
)

func ExternalAppPath(app external_app) string {
	path := which.Which(string(app))
	return path
}
