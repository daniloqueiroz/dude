package app

import (
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
)

const POLKIT_AGENT = "/usr/lib/polkit-gnome/polkit-gnome-authentication-agent-1"

func PolkitProc() *proc.Process {
	logger.Info("Starting PolKit agent")
	return proc.NewProcess(POLKIT_AGENT)
}
