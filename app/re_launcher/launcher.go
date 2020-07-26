package re_launcher

import (
	"github.com/daniloqueiroz/dude/app/re_launcher/plugins"
	"github.com/google/logger"
)

const (
	MainMenu Mode = "main_menu"
	SubMenu  Mode = "sub_menu"
)

type Launcher struct {
	plugins            map[plugins.Category]plugins.LauncherPlugin
	selectedCategories []plugins.Category
	availableActions   plugins.Actions
	mode               Mode
}

func (l *Launcher) Reset() {
	l.mode = MainMenu
	l.availableActions = nil
}
func (l *Launcher) GetSelectedCategories() []plugins.Category {
	return l.selectedCategories
}

func (l *Launcher) GetMode() Mode {
	return l.mode
}

func (l *Launcher) SelectCategories(categories []plugins.Category) {
	if l.mode == MainMenu {
		l.selectedCategories = categories
	}
}

func (l *Launcher) RefreshOptions(input string) {
	logger.Infof("Refreshing Launcher options: %s", input)
	if l.mode == MainMenu {
		var options plugins.Actions

		for _, name := range l.selectedCategories {
			plugin := l.plugins[name]
			if plugin != nil {
				suggested := plugin.FindActions(input)
				options = append(options, suggested...)
			}
		}

		l.availableActions = options
	} else {
		l.availableActions = plugins.FilterAction(input, l.availableActions)
	}
}

func (l *Launcher) AvailableActions() []plugins.Action {
	var actions []plugins.Action
	for _, action := range l.availableActions {
		actions = append(actions, action)
	}
	return actions
}

func (l *Launcher) ExecuteOption(position int) plugins.Result {
	selected := l.availableActions[position]
	l.selectedCategories = []plugins.Category{selected.Category()}
	logger.Infof("Selected action: %s::%s", selected.Category(), selected.Name())
	result := selected.Execute()
	switch res := result.(type) {
	case *plugins.SubActions:
		l.mode = SubMenu
		l.availableActions = res.SubActions
	default:
		l.mode = MainMenu
	}
	return result
}

func LauncherNew() *Launcher {
	// TODO make it configurable
	launcherPlugins := map[plugins.Category]plugins.LauncherPlugin{
		plugins.Application:  plugins.ApplicationsPluginNew(),
		plugins.Password:     plugins.PassPluginNew(),
		plugins.ShellCommand: plugins.ShellPluginNew(),
		plugins.System:       plugins.InternalPluginNew(),
		plugins.Web:          plugins.WebPluginNew(),
	}

	return &Launcher{
		plugins:          launcherPlugins,
		availableActions: nil,
		mode:             MainMenu,
	}
}
