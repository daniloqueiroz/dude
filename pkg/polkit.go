package pkg

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
)

func StartPolkit(wd *proc.Watchdog) {
	cmd := proc.NewProcess(internal.Config.AppPolkitAgent)
	wd.Supervise(cmd)
}