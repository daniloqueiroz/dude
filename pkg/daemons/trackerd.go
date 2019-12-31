package daemons

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"os"
	"path"
	"syscall"
	"time"
)

func TimeTrackd() {
	proc.CreatePidFile("trackerd")
	daemon := proc.NewDaemon(trackTime)
	daemon.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
}

func trackTime() {
	logger.Info("Starting trackerd")
	display := os.Getenv("DISPLAY")
	time_now := time.Now()
	year, wk_num := time_now.ISOWeek()
	trackingDir := path.Join(internal.Config.ScreenTimeDataDir, fmt.Sprintf("%d-w%d", year, wk_num))
	tracker, err := gone.NewTracker(display, trackingDir)
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	tracker.Start()
	logger.Info("Trackerd is running")
}
