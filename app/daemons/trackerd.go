package daemons

import (
	"github.com/daniloqueiroz/dude/app/appusage"
	"github.com/daniloqueiroz/dude/app/system/proc"
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
