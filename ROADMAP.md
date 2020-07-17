# TODO

- [x] Core: Move config file to ~/.config/dude.yml
- [ ] Launcher: system action "config" -> xdg-open ~/.config/dude.yml
- [ ] Launcher: Refactor to decouple UI/Controler/Results
- [ ] Launcher: Refactor results to multiple types
- [ ] Launcher: Refactor make Action Finders to break into subtokens when it's the case 
- [ ] Enable/Disable some integrations (such as pass)

# ROADMAP
## v1 
  * Session
      * wallpaper
      * compositor
      * screensaver
      * autostart
      * display daemon
      * power daemon
      * time tracker daemon
  * lock-screen
  * display profile 
  * brightness
  * terminal
  * !input
    * keyboard layout
    * !mouse settings
  * volume
  * !time tracker report
  * launcher
    * desktop files
    * console commands
    * pass store
    * internal
      * suspend
      * shutdown
      * ?display
      * brightness
      * lock-screen
      * !volume
      * !input/keyboard
      * !!config (xdg-open ~/.config/dude/config.yml)
## v2
  * iwctl/bluetooth ?
  * File search
  * launcher xdg-open
# v3
  * Notification Server
  * Polkit
