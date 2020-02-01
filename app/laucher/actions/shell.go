package actions

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/google/logger"
	"io/ioutil"
	"path/filepath"
)

const (
	ShellPrefix = "!"
)

type Shell struct {
	name        string
	fullpath    string
	description string
}

func (p Shell) Input() string {
	return fmt.Sprintf("%s%s", ShellPrefix, p.name)
}

func (p Shell) Description() string {
	return p.description
}

func (p Shell) Exec() {
	app.NewTerminalApp(p.fullpath)
}

func (p Shell) String() string {
	return p.Input()
}

func loadShellActions(actions map[string]laucher.Action) {
	dirname := "/usr/bin"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		logger.Error("Unable to load commands")
		return
	}
	for _, f := range files {
		if !f.IsDir() {
			fullFilepath := filepath.Join(dirname, f.Name())
			action := Shell{
				name:        f.Name(),
				fullpath:    fullFilepath,
				description: fmt.Sprintf("Execute '%s' in terminal", f.Name()),
			}
			actions[action.Input()] = action
		}
	}
}
