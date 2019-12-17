package internal

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/google/logger"
)

func StartScreensaver(wd *proc.Watchdog) {
	logger.Info("Configuring screensaver timeout and starting xss-lock")
	if err := proc.NewProcess(commons.Config.AppXset, "s", "300").FireAndWait(); err != nil {
		logger.Errorf("Unable to set screensaver timeout: %v", err)
	}

	cmd := proc.NewProcess(commons.Config.AppXssLock, commons.Config.AppXsecurelock)
	wd.Supervise(cmd)
}
