package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
)

func PolkitProc() *proc.Process {
	logger.Info("Starting PolKit agent")
	return proc.NewProcess(system.Config.PolkitAgent)
}
