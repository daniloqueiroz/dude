package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

func StartPolkit(wd *proc.Watchdog) {
	cmd := proc.NewProcess(system.Config.AppPolkitAgent)
	wd.Supervise(cmd)
}