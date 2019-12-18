package internal

import (
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"github.com/spf13/viper"
)

type config struct {
	DudeIcon         string
	AppFeh           string
	AppCompton       string
	AppXset          string
	AppXssLock       string
	AppXsecurelock   string
	AppAcpi          string
	AppXdotool       string
	AppPass          string
	AppPolkitAgent   string
	AppTmux          string
	AppTerminal      string
	AppBacklight     string
	TerminalFont     string
	TerminalFontSize string
	TimeTrackingFile string
	BackLightAC      int
	BackLightBattery int
	LauncherHeight   int
	LauncherWidth    int
}

var Config config

func InitConfig() {
	loadDefaults()
	loadFromFile()
	loadConfig()
}

func loadFromFile() {
	viper.SetConfigName("dude.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(basedir.ConfigHome)
	viper.AddConfigPath(".") // optionally look for config in the working directory
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
		DudeIcon:         viper.GetString("internal.icon"),
		AppFeh:           viper.GetString("internal.apps.feh"),
		AppCompton:       viper.GetString("internal.apps.compton"),
		AppXset:          viper.GetString("internal.apps.xset"),
		AppXssLock:       viper.GetString("internal.apps.xss_lock"),
		AppXsecurelock:   viper.GetString("internal.apps.xsecurelock"),
		AppAcpi:          viper.GetString("internal.apps.acpi"),
		AppXdotool:       viper.GetString("internal.apps.xdotool"),
		AppPass:          viper.GetString("internal.apps.pass"),
		AppPolkitAgent:   viper.GetString("internal.apps.polkit-agent"),
		AppTmux:          viper.GetString("internal.apps.tmux"),
		AppTerminal:      viper.GetString("internal.apps.terminal"),
		AppBacklight:     viper.GetString("internal.apps.xbacklight"),
		TerminalFont:     viper.GetString("terminal.font"),
		TerminalFontSize: viper.GetString("terminal.font_size"),
		TimeTrackingFile: viper.GetString("time_tracking.file"),
		BackLightAC:      viper.GetInt("backlight.ac"),
		BackLightBattery: viper.GetInt("backlight.battery"),
		LauncherWidth: viper.GetInt("launcher.width"),
		LauncherHeight: viper.GetInt("launcher.height"),
	}
}

func loadDefaults() {
	viper.SetDefault("internal.icon", "/usr/share/dude/fire-fist-colored.png")
	viper.SetDefault("internal.apps.feh", "/usr/bin/feh")
	viper.SetDefault("internal.apps.compton", "/usr/bin/compton")
	viper.SetDefault("internal.apps.xset", "/usr/bin/xset")
	viper.SetDefault("internal.apps.xss_lock", "/usr/bin/xss-lock")
	viper.SetDefault("internal.apps.xsecurelock", "/usr/bin/xsecurelock")
	viper.SetDefault("internal.apps.acpi", "/usr/bin/acpi")
	viper.SetDefault("internal.apps.xdotool", "/usr/bin/xdotool")
	viper.SetDefault("internal.apps.pass", "/usr/bin/pass")
	viper.SetDefault("internal.apps.polkit-agent", "/usr/bin/lxpolkit")
	viper.SetDefault("internal.apps.tmux", "/usr/bin/tmux")
	viper.SetDefault("internal.apps.terminal", "/usr/bin/st")
	viper.SetDefault("internal.apps.xbacklight", "/usr/bin/xbacklight")
	viper.SetDefault("terminal.font", "Source Code Pro")
	viper.SetDefault("terminal.font_size", "12")
	viper.SetDefault("time_tracking.file", "time-tracking.bin")
	viper.SetDefault("backlight.ac", "100")
	viper.SetDefault("backlight.battery", "20")
	viper.SetDefault("launcher.width", "600")
	viper.SetDefault("launcher.height", "250")
}
