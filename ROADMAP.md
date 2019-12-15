> dude: danilo's unique desktop environment

* I3 as WM
* DE
  * session management
    * WindowManager?
    * ~wallpaper -> feh
    * ~compositor -> compton
    * ~screensaver -> xss-lock + xsecurelock
    * !!monitor
  * ~powerd -> daemon for battery, backlight and cpu
  * ?input config
  * !!time-tracker -> gone (https://github.com/dim13/gone)
  * launcher -> lighthouse
    * DE commands -> prefix ":"
      * :display [single, mirror, auto] 
      * :shutdown 
      * ~:suspend 
      * ~:lock-screen
      * :volume [up, down, mute, mic(?)]
      * :brightness [up, down]
      * :keyboard <layout> -> modifies keyboard layout
      * ~:terminal -> launches st
      > launcher only operations
      * :kill <program> 
      * ~:pass
      * ::-> window switch
      * :o <whatever> -> xdg-open 
      * :e <file> -> howl <file>
      * :! <cmd> -> execute command on terminal
    * ~applications -> desktop entries
    * ~command line apps
    * ?files -> file path
    * ?network-manager
