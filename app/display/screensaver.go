package display

import (
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
)

func StartScreensaver() *proc.Process {
	logger.Info("Configuring screensaver timeout and starting xss-lock")
	if err := app.XSetScreensaverTimeProc().FireAndWait(); err != nil {
		logger.Errorf("Unable to set screensaver timeout: %v", err)
	}
	return app.XSSLockProc()
}
