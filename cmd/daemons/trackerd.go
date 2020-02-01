package daemons

import (
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/pkg/appusage"
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
	appusage.TrackAppUsage()
	logger.Info("Trackerd is running")
}
