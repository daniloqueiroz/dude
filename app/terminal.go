package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

func NewTmuxTerminal() error {
	return proc.NewProcess(
		system.Config.AppTerminal,
		"--command", system.Config.AppTmux, "new-session", "-t", "dude").FireAndForget()
}

func NewTerminalApp(cmd string) error {
	return proc.NewProcess(system.Config.AppTerminal, "--command", cmd).FireAndForget()
}
