package actions

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/google/logger"
	"io/ioutil"
	"path/filepath"
)

type Shell struct {
	shellApps laucher.Actions
}

func (s *Shell) Find(input string) laucher.Actions {
	if s.shellApps == nil {
		s.loadShellActions()
	}
	return laucher.FilterAction(input, s.shellApps)
}

func (s *Shell) loadShellActions() {
	dirname := "/usr/bin"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		logger.Error("Unable to load commands")
		return
	}
	var actions laucher.Actions
	for _, f := range files {
		if !f.IsDir() {
			fullFilepath := filepath.Join(dirname, f.Name())
			actions = append(actions, laucher.Action{
				Details: laucher.ActionMeta{
					Name:        f.Name(),
					Description: fmt.Sprintf("Execute '%s' in terminal", f.Name()),
					Category:    laucher.ShellCommand,
				},
				Exec: wrapShell(fullFilepath),
			})
		}
	}
	s.shellApps = actions
}

func wrapShell(cmd string) (func()) {
	return func() {
		app.NewTerminalApp(cmd)
	}
}