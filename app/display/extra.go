package display

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
)

func SetWallpaper() {
	logger.Info("Setting feh wallpaper")
	cmd := proc.NewProcess(system.Config.AppFeh, "--bg-fill", "--randomize", system.Config.WallpaperDir)
	cmd.FireAndForget()
}

func StartCompositor(wd *proc.Watchdog) {
	logger.Info("Starting compton compositor")
	cmd := proc.NewProcess(system.Config.AppCompton, "-d", ":0")
	wd.Supervise(cmd)
}
