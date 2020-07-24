D's Ultimate Desktop Environment

# About it
Dude's principles:

* For X11, WM agnostic
* Plain text file (yaml) configurable
* Simple consistent UX - cli & GTK3 launcher
* Modular Monolith

# Features
  - GTK Launcher
  - Brightness control - with AC/Battery brightness auto adjust
  - Startup apps support (~/.config/autostart/)
  - Display profiles
  - Pass integration
  - App Time tracking
  - bluetooth and wifi (iwd) support

# Dependencies
- feh - for background image
- picom - compositor for x11
- xss-lock - bridges an x screensaver to systemd's login manager
- xsecurelock - X11 screen lock utility designed for security
- acpi - for battery info and AC/Battery adjustments
- pass - password manager
- xdotool - command-line X11 automation tool - pass dependency for auto fill 
- lxpolkit - gtk polkit-agent for lxde 
- brightnessctl - brightness control for X/wayland and driver agnostic
- tmux
- alacritty
- bluez
- iwd

# Using dude with i3

Add the lines below to your `~/.config/i3/config`:

```
set $dude /usr/bin/dude

## Start dude session
exec --no-startup-id $dude session
# Keybinds
bindsym $mod+space exec --no-startup-id $dude launcher
bindsym $mod+Return exec $dude terminal
bindsym $mod+l exec $dude lock-screen

bindsym XF86MonBrightnessUp exec $dude brightness up
bindsym XF86MonBrightnessDown exec $dude brightness down

bindsym XF86AudioRaiseVolume exec --no-startup-id $dude audio vol-up
bindsym XF86AudioLowerVolume exec --no-startup-id $dude audio vol-down
bindsym XF86AudioMute exec --no-startup-id $dude audio vol-mute
bindsym XF86AudioMicMute exec --no-startup-id $dude audio mic-mute
```

# Credits
- Dude uses a modified verion of [Gone](https://github.com/dim13/gone) and [iwd](https://github.com/shibumi/iwd) go lib.
- Icons made by [Freepik](https://www.flaticon.com/authors/freepik) from [www.flaticon.com/](https://www.flaticon.com/)
- Icons used on Launcher are from [FontAwesome](https://fontawesome.com/license/free) and are under [CC BY 4.0 License](https://creativecommons.org/licenses/by/4.0/)
