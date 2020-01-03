package display

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/google/logger"
)

func StartScreensaver(wd *proc.Watchdog) {
	logger.Info("Configuring screensaver timeout and starting xss-lock")
	if err := proc.NewProcess(internal.Config.AppXset, "s", string(internal.Config.ScreenSaverTimeoutSecs)).FireAndWait(); err != nil {
		logger.Errorf("Unable to set screensaver timeout: %v", err)
	}

	cmd := proc.NewProcess(internal.Config.AppXssLock, internal.Config.AppXsecurelock)
	wd.Supervise(cmd)
}
