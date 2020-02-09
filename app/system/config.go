package system

import (
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"github.com/spf13/viper"
	"path"
)

type config struct {
	DudeIcon                 string
	LauncherIconsFolder      string
	WallpaperDir             string
	AppFeh                   string
	AppCompton               string
	AppXset                  string
	AppXssLock               string
	AppXsecurelock           string
	AppAcpi                  string
	AppPolkitAgent           string
	AppTmux                  string
	AppTerminal              string
	AppBacklight             string
	AppXrandr                string
	AppSetxkbmap             string
	AppPactl                 string
	TerminalFont             string
	TerminalFontSize         string
	AppUsageDataDir          string
	BackLightAC              int
	BackLightBattery         int
	LauncherHeight           int
	LauncherWidth            int
	Profiles                 map[string]interface{}
	Keyboards                map[string]interface{}
	AppUsageTrackEnable      bool
	DisplayAutoConfigEnabled bool
	ScreenSaverTimeoutSecs   int
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
		LauncherIconsFolder:      viper.GetString("launcher.icons_folder"),
		AppFeh:                   viper.GetString("apps.feh"),
		AppCompton:               viper.GetString("apps.compton"),
		AppXset:                  viper.GetString("apps.xset"),
		AppXssLock:               viper.GetString("apps.xss_lock"),
		AppXsecurelock:           viper.GetString("apps.xsecurelock"),
		AppAcpi:                  viper.GetString("apps.acpi"),
		AppPolkitAgent:           viper.GetString("apps.polkit-agent"),
		AppTmux:                  viper.GetString("apps.tmux"),
		AppTerminal:              viper.GetString("apps.terminal"),
		AppBacklight:             viper.GetString("apps.xbacklight"),
		AppXrandr:                viper.GetString("apps.xrandr"),
		AppSetxkbmap:             viper.GetString("apps.setxkbmap"),
		AppPactl:                 viper.GetString("apps.pactl"),
		TerminalFont:             viper.GetString("terminal.font"),
		TerminalFontSize:         viper.GetString("terminal.font_size"),
		AppUsageDataDir:          viper.GetString("app_usage.data_dir"),
		AppUsageTrackEnable:      viper.GetBool("app_usage.enabled"),
		LauncherWidth:            viper.GetInt("launcher.width"),
		LauncherHeight:           viper.GetInt("launcher.height"),
		Profiles:                 viper.GetStringMap("display.profiles"),
		WallpaperDir:             viper.GetString("display.wallpapers_dir"),
		DisplayAutoConfigEnabled: viper.GetBool("display.autoconfig_enabled"),
		ScreenSaverTimeoutSecs:   viper.GetInt("display.screensaver_timeout_secs"),
		BackLightAC:              viper.GetInt("display.brightness.ac"),
		BackLightBattery:         viper.GetInt("display.brightness.battery"),
		Keyboards:                viper.GetStringMap("input.keyboards"),
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
	viper.SetDefault("apps.polkit-agent", "/usr/bin/lxpolkit")
	viper.SetDefault("apps.tmux", "/usr/bin/tmux")
	viper.SetDefault("apps.terminal", "/usr/bin/st")
	viper.SetDefault("apps.xbacklight", "/usr/bin/xbacklight")
	viper.SetDefault("apps.xrandr", "/usr/bin/xrandr")
	viper.SetDefault("apps.setxkbmap", "/usr/bin/setxkbmap")
	viper.SetDefault("apps.pactl", "/usr/bin/pactl")
	viper.SetDefault("terminal.font", "Source Code Pro")
	viper.SetDefault("terminal.font_size", "12")
	viper.SetDefault("app_usage.data_dir", "appusage")
	viper.SetDefault("app_usage.enabled", "true")
	viper.SetDefault("launcher.width", "600")
	viper.SetDefault("launcher.height", "250")
	viper.SetDefault("launcher.icons_folder", "/usr/share/dude/launcher")
	viper.SetDefault("display.wallpapers_dir", path.Join(basedir.ConfigHome, "/dude/wallpapers"))
	viper.SetDefault("display.autoconfig_enabled", "true")
	viper.SetDefault("display.screensaver_timeout_secs", "300")
	viper.SetDefault("display.brightness.ac", "100")
	viper.SetDefault("display.brightness.battery", "45")

}
