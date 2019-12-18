package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"os"
	"path"
	"syscall"
)

func TimeTrackd() {
	daemon := proc.NewDaemon(trackTime)
	daemon.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
}

func trackTime() {
	logger.Info("Starting trackerd")
	display := os.Getenv("DISPLAY")
	timeFile := path.Join(basedir.CacheHome, internal.Config.TimeTrackingFile)
	tracker, err := gone.NewTracker(display, timeFile)
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	tracker.Start()
	logger.Info("Trackerd is running")
}
