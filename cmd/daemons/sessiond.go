package daemons

import (
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"syscall"
	"time"
)


func Sessiond() {
	proc.CreatePidFile("sessiond")
	wd := proc.NewWatchdog()
	display.StartCompositor(wd)
	display.StartScreensaver(wd)
	app.StartPolkit(wd)
	app.AutostartApps()
	proc.LaunchDaemon(wd, "powerd")

	if system.Config.DisplayAutoConfigEnabled {
		proc.LaunchDaemon(wd, "displayd")
	}

	if system.Config.AppUsageTrackEnable {
		proc.LaunchDaemon(wd, "trackerd")
	}
	system.SimpleNotification("Session started").Show()

	wd.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	time.Sleep(1 * time.Second)
	logger.Info("Session ended")
}
