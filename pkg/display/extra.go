package display

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/google/logger"
)

func SetWallpaper() {
	logger.Info("Setting feh wallpaper")
	cmd := proc.NewProcess(internal.Config.AppFeh, "--bg-fill", "--randomize", internal.Config.WallpaperDir)
	cmd.FireAndForget()
}

func StartCompositor(wd *proc.Watchdog) {
	logger.Info("Starting compton compositor")
	cmd := proc.NewProcess(internal.Config.AppCompton, "-d", ":0")
	wd.Supervise(cmd)
}
