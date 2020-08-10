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
	BrightnessAC              int
	BrightnessBattery         int
	LauncherHeight            int
	LauncherWidth             int
	LauncherUIFolder          string
	LauncherDefaultCategories []string
	LauncherWebQueryURL       string
	Profiles                  map[string]interface{}
	Keyboards                 map[string]interface{}
	BatteryMonitorEnabled     bool
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
	viper.SetConfigName("dude.yml")
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
		LauncherWidth:             viper.GetInt("launcher.width"),
		LauncherHeight:            viper.GetInt("launcher.height"),
		LauncherDefaultCategories: viper.GetStringSlice("launcher.default_categories"),
		LauncherWebQueryURL:       viper.GetString("launcher.plugins.web_search_url"),
		Profiles:                  viper.GetStringMap("display.profiles"),
		WallpaperDir:              viper.GetString("display.wallpapers_dir"),
		BatteryMonitorEnabled:     viper.GetBool("power.monitor_enabled"),
		DisplayAutoConfigEnabled:  viper.GetBool("display.autoconfig_enabled"),
		ScreenSaverTimeoutSecs:    viper.GetInt("display.screensaver_timeout_secs"),
		BrightnessAC:              viper.GetInt("display.brightness.ac"),
		BrightnessBattery:         viper.GetInt("display.brightness.battery"),
		Keyboards:                 viper.GetStringMap("input.keyboards"),
	}
}

func loadDefaults() {
	viper.SetDefault("icon", "/usr/share/dude/dude.png")
	viper.SetDefault("launcher.width", "700")
	viper.SetDefault("launcher.height", "250")
	viper.SetDefault("launcher.ui_folder", "/usr/share/dude/ui")
	viper.SetDefault("launcher.default_categories", []string{"applications", "passwords", "system"})
	viper.SetDefault("launcher.plugins.web_search_url", "https://www.google.com/search")
	viper.SetDefault("display.wallpapers_dir", path.Join(basedir.ConfigHome, "/wallpapers"))
	viper.SetDefault("power.monitor_enabled", "true")
	viper.SetDefault("display.autoconfig_enabled", "true")
	viper.SetDefault("display.screensaver_timeout_secs", "300")
	viper.SetDefault("display.brightness.ac", "100")
	viper.SetDefault("display.brightness.battery", "45")
}
