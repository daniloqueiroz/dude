package internal

import (
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"github.com/spf13/viper"
	"path"
)

type config struct {
	DudeIcon                 string
	WallpaperDir             string
	AppFeh                   string
	AppCompton               string
	AppXset                  string
	AppXssLock               string
	AppXsecurelock           string
	AppAcpi                  string
	AppXdotool               string
	AppPass                  string
	AppPolkitAgent           string
	AppTmux                  string
	AppTerminal              string
	AppBacklight             string
	AppXrandr                string
	TerminalFont             string
	TerminalFontSize         string
	ScreenTimeDataDir        string
	BackLightAC              int
	BackLightBattery         int
	LauncherHeight           int
	LauncherWidth            int
	Profiles                 map[string]interface{}
	ScreenTimeEnabled        bool
	DisplayAutoConfigEnabled bool
}

var Config config

func InitConfig() {
	loadDefaults()
	loadFromFile()
	loadConfig()
}

func loadFromFile() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(basedir.ConfigHome, "dude"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info("config file not found...")
		} else {
			logger.Fatalf("Error loading config: %s", err)
		}
	}
}

func loadConfig() {
	Config = config{
		DudeIcon:                 viper.GetString("icon"),
		AppFeh:                   viper.GetString("apps.feh"),
		AppCompton:               viper.GetString("apps.compton"),
		AppXset:                  viper.GetString("apps.xset"),
		AppXssLock:               viper.GetString("apps.xss_lock"),
		AppXsecurelock:           viper.GetString("apps.xsecurelock"),
		AppAcpi:                  viper.GetString("apps.acpi"),
		AppXdotool:               viper.GetString("apps.xdotool"),
		AppPass:                  viper.GetString("apps.pass"),
		AppPolkitAgent:           viper.GetString("apps.polkit-agent"),
		AppTmux:                  viper.GetString("apps.tmux"),
		AppTerminal:              viper.GetString("apps.terminal"),
		AppBacklight:             viper.GetString("apps.xbacklight"),
		AppXrandr:                viper.GetString("apps.xrandr"),
		TerminalFont:             viper.GetString("terminal.font"),
		TerminalFontSize:         viper.GetString("terminal.font_size"),
		ScreenTimeDataDir:        viper.GetString("screen_time.data_dir"),
		ScreenTimeEnabled:        viper.GetBool("screen_time.enabled"),
		BackLightAC:              viper.GetInt("backlight.ac"),
		BackLightBattery:         viper.GetInt("backlight.battery"),
		LauncherWidth:            viper.GetInt("launcher.width"),
		LauncherHeight:           viper.GetInt("launcher.height"),
		Profiles:                 viper.GetStringMap("display.profiles"),
		WallpaperDir:             viper.GetString("display.wallpapers_dir"),
		DisplayAutoConfigEnabled: viper.GetBool("display.autoconfig_enabled"),
	}
}

func loadDefaults() {
	viper.SetDefault("icon", "/usr/share/dude/dude.png")
	viper.SetDefault("apps.feh", "/usr/bin/feh")
	viper.SetDefault("apps.compton", "/usr/bin/compton")
	viper.SetDefault("apps.xset", "/usr/bin/xset")
	viper.SetDefault("apps.xss_lock", "/usr/bin/xss-lock")
	viper.SetDefault("apps.xsecurelock", "/usr/bin/xsecurelock")
	viper.SetDefault("apps.acpi", "/usr/bin/acpi")
	viper.SetDefault("apps.xdotool", "/usr/bin/xdotool")
	viper.SetDefault("apps.pass", "/usr/share/dude/passtype.sh")
	viper.SetDefault("apps.polkit-agent", "/usr/bin/lxpolkit")
	viper.SetDefault("apps.tmux", "/usr/bin/tmux")
	viper.SetDefault("apps.terminal", "/usr/bin/st")
	viper.SetDefault("apps.xbacklight", "/usr/bin/xbacklight")
	viper.SetDefault("apps.xrandr", "/usr/bin/xrandr")
	viper.SetDefault("terminal.font", "Source Code Pro")
	viper.SetDefault("terminal.font_size", "12")
	viper.SetDefault("screen_time.data_dir", path.Join(basedir.DataHome, "screen-time"))
	viper.SetDefault("screen_time.enabled", "true")
	viper.SetDefault("backlight.ac", "100")
	viper.SetDefault("backlight.battery", "20")
	viper.SetDefault("launcher.width", "600")
	viper.SetDefault("launcher.height", "250")
	viper.SetDefault("display.wallpapers_dir", path.Join(basedir.ConfigHome, "/dude/wallpapers"))
	viper.SetDefault("display.autoconfig_enabled", "true")
}
