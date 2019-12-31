package pkg

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/internal/system"
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
