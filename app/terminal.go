package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

func NewTmuxTerminal() error {
	return proc.NewProcess(
		system.ExternalAppPath(system.TERMINAL),
		"--command", system.ExternalAppPath(system.TMUX), "new-session", "-t", "dude").FireAndForget()
}

func NewTerminalApp(cmd string) error {
	return proc.NewProcess(system.ExternalAppPath(system.TERMINAL), "--command", cmd).FireAndForget()
}

func NewTerminalAppWithArgs(cmd string, params []string) error {
	args := []string{"--command", cmd}
	for _, arg := range params {
		args = append(args, arg)
	}
	return proc.NewProcess(system.ExternalAppPath(system.TERMINAL), args...).FireAndForget()
}
