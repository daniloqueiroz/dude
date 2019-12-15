package daemons

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"os"
	"path"
)


func TimeTrackd() {
	logger.Info("Starting time track deamon")
	display := os.Getenv("DISPLAY")
	timeFile := path.Join(basedir.CacheHome, commons.Config.TimeTrackingFile)
	gone.StartTracker(display, timeFile)
	logger.Info("started")
	select { }
}
