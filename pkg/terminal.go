package pkg

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
)

func NewTmuxTerminal() error {
	return proc.NewProcess(
		internal.Config.AppTerminal,
		fmt.Sprintf("-f %s:size=%s", internal.Config.TerminalFont, internal.Config.TerminalFontSize),
		"-T", "terminal", internal.Config.AppTmux, "new-session", "-t", "dude").FireAndForget()
}

func NewTerminalApp(cmd string) error {
	return proc.NewProcess(
		internal.Config.AppTerminal,
		fmt.Sprintf("-f %s:size=%s", internal.Config.TerminalFont, internal.Config.TerminalFontSize),
		"-T", cmd, "-e", cmd).FireAndForget()
}