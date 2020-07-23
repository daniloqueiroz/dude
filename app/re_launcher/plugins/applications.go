package plugins

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/rkoesters/xdg/basedir"
	"path"
	"strings"
)

type appAction struct {
	cmd         string
	name        string
	description string
}

func (ac *appAction) Category() Category {
	return Application
}
func (ac *appAction) Name() string {
	return ac.name
}
func (ac *appAction) Description() string {
	return ac.description
}

func (ac *appAction) Execute() Result {
	command := strings.Fields(ac.cmd)
	proc.NewProcess(command[0], command[1:]...).FireAndForget()
	return Empty{}
}

type appPlugin struct {
	desktopApps Actions
}

func (a *appPlugin) Category() Category {
	return Application
}

func (a *appPlugin) FindActions(input string) Actions {
	return FilterAction(input, a.desktopApps)
}

func ApplicationsPluginNew() LauncherPlugin {
	dirs := append([]string(nil), basedir.DataDirs...)
	dirs = append(dirs, basedir.DataHome)
	var apps Actions
	for _, dir := range dirs {

		dir = path.Join(dir, "applications")
		share_apps := system.LoadDesktopEntries(dir)
		for _, app := range share_apps {
			action := &appAction{
				name:        strings.ToLower(app.Name),
				description: strings.TrimSpace(app.GenericName),
				cmd:         app.Exec,
			}
			apps = append(apps, action)
		}
	}
	return &appPlugin{
		desktopApps: apps,
	}
}
