package internal

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
)

func NewTmuxTerminal() error {
	return proc.NewProcess(
		commons.Config.AppTerminal,
		fmt.Sprintf("-f %s:size=%s", commons.Config.TerminalFont, commons.Config.TerminalFontSize),
		"-T", "terminal", commons.Config.AppTmux, "new-session", "-t", "dude").FireAndForget()
}

func NewTerminalApp(cmd string) error {
	return proc.NewProcess(
		commons.Config.AppTerminal,
		fmt.Sprintf("-f %s:size=%s", commons.Config.TerminalFont, commons.Config.TerminalFontSize),
		"-T", cmd, "-e", cmd).FireAndForget()
}