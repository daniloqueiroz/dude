package app

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

func NewTmuxTerminal() error {
	return proc.NewProcess(
		system.Config.AppTerminal,
		fmt.Sprintf("-f %s:size=%s", system.Config.TerminalFont, system.Config.TerminalFontSize),
		"-T", "terminal", system.Config.AppTmux, "new-session", "-t", "dude").FireAndForget()
}

func NewTerminalApp(cmd string) error {
	return proc.NewProcess(
		system.Config.AppTerminal,
		fmt.Sprintf("-f %s:size=%s", system.Config.TerminalFont, system.Config.TerminalFontSize),
		"-T", cmd, "-e", cmd).FireAndForget()
}