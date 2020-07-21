package re_launcher

import (
	"github.com/daniloqueiroz/dude/app/re_launcher/plugins"
	"github.com/google/logger"
)

type Launcher struct {
	plugins          map[plugins.Category]plugins.LauncherPlugin
	availableActions plugins.Actions
}

func (l *Launcher) RefreshOptions(input string, categories []plugins.Category) {
	logger.Infof("Refreshing Launcher options: %s", input)
	var options plugins.Actions

	for _, name := range categories {
		plugin := l.plugins[name]
		if plugin != nil {
			suggested := plugin.FindActions(input)
			options = append(options, suggested...)
		}
	}

	l.availableActions = options
}

func (l *Launcher) AvailableActions() []plugins.Action {
	var actions []plugins.Action
	for _, action := range l.availableActions {
		actions = append(actions, action)
	}
	return actions
}

func (l *Launcher) ExecuteOption(position int) {
	selected := l.availableActions[position]
	logger.Infof("Selected action: %s::%s", selected.Category(), selected.Name())
	selected.Execute()
	// Get the option at that position
	// Clear current options
	// Execute option
	// It result is options, update the current options
	// Returns the result
}

func LauncherNew() *Launcher {
	// TODO make it configurable
	launcherPlugins := map[plugins.Category]plugins.LauncherPlugin{
		plugins.Application:  plugins.ApplicationsPluginNew(),
		plugins.Password:     plugins.PassPluginNew(),
		plugins.ShellCommand: plugins.ShellPluginNew(),
		plugins.System:       plugins.InternalPluginNew(),
	}

	return &Launcher{
		plugins:          launcherPlugins,
		availableActions: nil,
	}
}
