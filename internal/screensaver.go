package internal

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/google/logger"
)

func StartScreensaver() {
	logger.Info("Configuring screensaver timeout and starting xss-lock")
	if err := proc.NewProcess(commons.Config.AppXset, "s", "300").FireAndWait(); err != nil {
		logger.Errorf("Unable to set screensaver timeout: %v", err)
	}

	process := proc.NewProcess(commons.Config.AppXssLock, commons.Config.AppXsecurelock)
	if err := process.FireAndKeepAlive(100); err != nil {
		logger.Errorf("Screensaver has died: %v", err)
	}
}
