package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/internal/system"
	"github.com/daniloqueiroz/dude/pkg"
	"github.com/daniloqueiroz/dude/pkg/display"
	"github.com/google/logger"
	"syscall"
	"time"
)


func Sessiond() {
	proc.CreatePidFile("sessiond")
	wd := proc.NewWatchdog()
	display.StartCompositor(wd)
	display.StartScreensaver(wd)
	pkg.StartPolkit(wd)
	pkg.AutostartApps()
	proc.LaunchDaemon(wd, "powerd")

	if internal.Config.DisplayAutoConfigEnabled {
		proc.LaunchDaemon(wd, "displayd")
	}

	if internal.Config.ScreenTimeEnabled {
		proc.LaunchDaemon(wd, "trackerd")
	}
	system.SimpleNotification("Session started").Show()

	wd.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	time.Sleep(1 * time.Second)
	logger.Info("Session ended")
}
