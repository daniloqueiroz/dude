package internal

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
)

func StartPolkit(wd *proc.Watchdog) {
	cmd := proc.NewProcess(commons.Config.AppPolkitAgent)
	wd.Supervise(cmd)
}