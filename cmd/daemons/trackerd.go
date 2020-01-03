package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/pkg"
	"github.com/google/logger"
	"syscall"
)

func TimeTrackerd() {
	proc.CreatePidFile("trackerd")
	daemon := proc.NewDaemon(trackTime)
	daemon.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
}

func trackTime() {
	logger.Info("Starting trackerd")
	tracker := pkg.NewTracker(internal.Config.ScreenTimeDataDir)
	tracker.Track()
	logger.Info("Trackerd is running")
}
