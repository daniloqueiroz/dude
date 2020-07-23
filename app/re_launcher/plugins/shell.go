package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/google/logger"
	"io/ioutil"
	"path/filepath"
)

type shellAction struct {
	cmd  string
	name string
}

func (sa *shellAction) Category() Category {
	return ShellCommand
}
func (sa *shellAction) Name() string {
	return sa.name
}
func (sa *shellAction) Description() string {
	return fmt.Sprintf("Execute '%s' in terminal", sa.Name())
}

func (sa *shellAction) Execute() Result {
	app.NewTerminalApp(sa.cmd)
	return Empty{}
}

type shellPlugin struct {
	shellApps Actions
}

func (s *shellPlugin) Category() Category {
	return ShellCommand
}

func (s *shellPlugin) FindActions(input string) Actions {
	return FilterAction(input, s.shellApps)
}

func ShellPluginNew() LauncherPlugin {
	return &shellPlugin{
		shellApps: loadShellActions(),
	}
}

func loadShellActions() Actions {
	dirname := "/usr/bin"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		logger.Error("Unable to load commands")
		return nil
	}
	var actions Actions
	for _, f := range files {
		if !f.IsDir() {
			fullFilepath := filepath.Join(dirname, f.Name())
			actions = append(actions, &shellAction{
				name: f.Name(),
				cmd:  fullFilepath,
			})
		}
	}
	return actions
}
