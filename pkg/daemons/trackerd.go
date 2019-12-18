package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"os"
	"path"
)


func TimeTrackd() {
	logger.Info("Starting time track deamon")
	display := os.Getenv("DISPLAY")
	timeFile := path.Join(basedir.CacheHome, internal.Config.TimeTrackingFile)
	tracker, err := gone.NewTracker(display, timeFile)
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	tracker.Start()
	logger.Info("Time trackerd started")
	select { }
}
