# TODO

- [x] Core: Move config file to ~/.config/dude.yml
- [x] Launcher: system action "config" -> xdg-open ~/.config/dude.yml
- [x] Launcher: Refactor to decouple UI/Controler/Results
- [x] Launcher: Refactor results to multiple types
- [x] Launcher: Refactor make Action Finders to break into subtokens when it's the case (support params for commands) 
- [x] Launcher: Open url
- [ ] Launcher: File manager integration (pick file manager)
- [ ] Keyboard and Mouse/Touchpad config
- [ ] Enable/Disable some integrations (such as pass, possibly some bluetooth and wifi as well)
- [ ] Improve bluetooth and wifi integration (bluetooth pair, wifi disconnect, wifi connect unknown)
- [ ] Launcher: System action: Airplane mode toggle

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
      * config (xdg-open ~/.config/dude/config.yml)
      * iwctl/bluetooth ?
# v2
  * Notification Server
  * Polkit
