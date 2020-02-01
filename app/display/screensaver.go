package display

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
)

func StartScreensaver(wd *proc.Watchdog) {
	logger.Info("Configuring screensaver timeout and starting xss-lock")
	if err := proc.NewProcess(system.Config.AppXset, "s", string(system.Config.ScreenSaverTimeoutSecs)).FireAndWait(); err != nil {
		logger.Errorf("Unable to set screensaver timeout: %v", err)
	}

	cmd := proc.NewProcess(system.Config.AppXssLock, system.Config.AppXsecurelock)
	wd.Supervise(cmd)
}
