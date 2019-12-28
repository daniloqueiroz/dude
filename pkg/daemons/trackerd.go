package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"os"
	"syscall"
)

func TimeTrackd() {
	daemon := proc.NewDaemon(trackTime)
	daemon.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
}

func trackTime() {
	logger.Info("Starting trackerd")
	display := os.Getenv("DISPLAY")
	tracker, err := gone.NewTracker(display, internal.Config.ScreenTimeDataDir)
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	tracker.Start()
	logger.Info("Trackerd is running")
}
