package daemons

import (
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/internal/system"
	"github.com/daniloqueiroz/dude/pkg"
	"github.com/google/logger"
	"syscall"
	"time"
)


func Sessiond() {
	wd := proc.NewWatchdog()
	pkg.StartCompositor(wd)
	pkg.StartScreensaver(wd)
	pkg.StartPolkit(wd)
	pkg.SetWallpaper()
	pkg.AutostartApps()
	proc.DaemonExec(wd, "powerd")
	proc.DaemonExec(wd, "trackerd")
	system.SimpleNotification("Session started").Show()

	wd.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	time.Sleep(1 * time.Second)
	logger.Info("Session ended")
}
