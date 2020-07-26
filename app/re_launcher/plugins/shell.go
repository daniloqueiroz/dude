package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/google/logger"
	"github.com/sahilm/fuzzy"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type shellAction struct {
	cmd    string
	params []string
	name   string
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
	app.NewTerminalAppWithArgs(sa.cmd, sa.params)
	return Empty{}
}

type shellPlugin struct {
	apps  []string
	paths map[string]string
}

func (s *shellPlugin) Category() Category {
	return ShellCommand
}

func (s *shellPlugin) FindActions(input string) Actions {
	tokens := strings.Split(input, " ")
	base := tokens[0]

	var params []string
	if len(tokens) > 1 {
		params = tokens[1:]
	}

	var results Actions
	matches := fuzzy.Find(base, s.apps)
	for _, match := range matches {
		if match.Score > 0 {
			cmd := s.apps[match.Index]
			results = append(results, &shellAction{
				name:   cmd,
				cmd:    s.paths[cmd],
				params: params,
			})
		}
	}
	return results
}

func ShellPluginNew() LauncherPlugin {
	plugin := &shellPlugin{
		apps:  make([]string, 10),
		paths: make(map[string]string, 10),
	}
	plugin.loadShellApps()
	return plugin
}

func (s *shellPlugin) loadShellApps() {
	dirname := "/usr/bin"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		logger.Error("Unable to load commands")
		return
	}

	for _, f := range files {
		if !f.IsDir() {
			fullFilepath := filepath.Join(dirname, f.Name())
			s.apps = append(s.apps, f.Name())
			s.paths[f.Name()] = fullFilepath
		}
	}
}
