package system

import (
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"github.com/spf13/viper"
	"path"
)

type config struct {
	DudeIcon                  string
	WallpaperDir              string
	AppFeh                    string
	AppCompositor             string
	AppXset                   string
	AppXssLock                string
	AppXsecurelock            string
	AppAcpi                   string
	AppPolkitAgent            string
	AppTmux                   string
	AppTerminal               string
	AppBrightness             string
	AppXrandr                 string
	AppSetxkbmap              string
	AppPactl                  string
	AppXdgOpen                string
	AppBluetoothCtl           string
	AppWifiCtl                string
	AppUsageDataDir           string
	BrightnessAC              int
	BrightnessBattery         int
	LauncherHeight            int
	LauncherWidth             int
	LauncherUIFolder          string
	LauncherDefaultCategories []string
	Profiles                  map[string]interface{}
	Keyboards                 map[string]interface{}
	AppUsageTrackEnable       bool
	DisplayAutoConfigEnabled  bool
	ScreenSaverTimeoutSecs    int
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
		DudeIcon:                  viper.GetString("icon"),
		LauncherUIFolder:          viper.GetString("launcher.ui_folder"),
		AppFeh:                    viper.GetString("apps.feh"),
		AppCompositor:             viper.GetString("apps.picom"),
		AppXset:                   viper.GetString("apps.xset"),
		AppXssLock:                viper.GetString("apps.xss_lock"),
		AppXsecurelock:            viper.GetString("apps.xsecurelock"),
		AppAcpi:                   viper.GetString("apps.acpi"),
		AppPolkitAgent:            viper.GetString("apps.polkit-agent"),
		AppTmux:                   viper.GetString("apps.tmux"),
		AppTerminal:               viper.GetString("apps.terminal"),
		AppBrightness:             viper.GetString("apps.brightness"),
		AppXrandr:                 viper.GetString("apps.xrandr"),
		AppSetxkbmap:              viper.GetString("apps.setxkbmap"),
		AppPactl:                  viper.GetString("apps.pactl"),
		AppBluetoothCtl:           viper.GetString("apps.bluetoothctl"),
		AppWifiCtl:                viper.GetString("apps.iwctl"),
		AppXdgOpen:                viper.GetString("apps.xdg_open"),
		AppUsageDataDir:           viper.GetString("app_usage.data_dir"),
		AppUsageTrackEnable:       viper.GetBool("app_usage.enabled"),
		LauncherWidth:             viper.GetInt("launcher.width"),
		LauncherHeight:            viper.GetInt("launcher.height"),
		LauncherDefaultCategories: viper.GetStringSlice("launcher.default_categories"),
		Profiles:                  viper.GetStringMap("display.profiles"),
		WallpaperDir:              viper.GetString("display.wallpapers_dir"),
		DisplayAutoConfigEnabled:  viper.GetBool("display.autoconfig_enabled"),
		ScreenSaverTimeoutSecs:    viper.GetInt("display.screensaver_timeout_secs"),
		BrightnessAC:              viper.GetInt("display.brightness.ac"),
		BrightnessBattery:         viper.GetInt("display.brightness.battery"),
		Keyboards:                 viper.GetStringMap("input.keyboards"),
	}
}

func loadDefaults() {
	viper.SetDefault("icon", "/usr/share/dude/dude.png")
	viper.SetDefault("apps.feh", "/usr/bin/feh")
	viper.SetDefault("apps.picom", "/usr/bin/picom")
	viper.SetDefault("apps.xset", "/usr/bin/xset")
	viper.SetDefault("apps.xss_lock", "/usr/bin/xss-lock")
	viper.SetDefault("apps.xsecurelock", "/usr/bin/xsecurelock")
	viper.SetDefault("apps.acpi", "/usr/bin/acpi")
	viper.SetDefault("apps.polkit-agent", "/usr/bin/lxpolkit")
	viper.SetDefault("apps.tmux", "/usr/bin/tmux")
	viper.SetDefault("apps.terminal", "/usr/bin/alacritty")
	viper.SetDefault("apps.brightness", "/usr/bin/brightnessctl")
	viper.SetDefault("apps.xrandr", "/usr/bin/xrandr")
	viper.SetDefault("apps.setxkbmap", "/usr/bin/setxkbmap")
	viper.SetDefault("apps.pactl", "/usr/bin/pactl")
	viper.SetDefault("apps.bluetoothctl", "/usr/bin/bluetoothctl")
	viper.SetDefault("apps.iwctl", "/usr/bin/iwctl")
	viper.SetDefault("apps.xdg_open", "/usr/bin/xdg-open")
	viper.SetDefault("app_usage.data_dir", "appusage")
	viper.SetDefault("app_usage.enabled", "true")
	viper.SetDefault("launcher.width", "700")
	viper.SetDefault("launcher.height", "250")
	viper.SetDefault("launcher.ui_folder", "/usr/share/dude/ui")
	viper.SetDefault("launcher.default_categories", []string{"applications", "passwords", "system"})
	viper.SetDefault("display.wallpapers_dir", path.Join(basedir.ConfigHome, "/wallpapers"))
	viper.SetDefault("display.autoconfig_enabled", "true")
	viper.SetDefault("display.screensaver_timeout_secs", "300")
	viper.SetDefault("display.brightness.ac", "100")
	viper.SetDefault("display.brightness.battery", "45")

}
