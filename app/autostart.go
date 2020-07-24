package app

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"path/filepath"
)

func AutostartApps() {
	entries := system.LoadDesktopEntries(filepath.Join(system.HomeDir(), "/.config/autostart"))
	for _, entry := range entries {
		logger.Info(fmt.Sprintf("Autostarting %s", entry.Exec))
		proc.NewProcess(entry.Exec).FireAndForget()
	}
}
