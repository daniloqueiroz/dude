package plugins

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/rkoesters/xdg/basedir"
	"path"
	"strings"
)

type AppAction struct {
	cmd string
	name string
	description string
}

func (ac AppAction) Category() Category {
	return Application
}
func (ac AppAction) Name() string {
	return ac.name
}
func (ac AppAction) Description() string {
	return ac.description
}

func (ac AppAction) Execute() Result {
	command := strings.Fields(ac.cmd)
	proc.NewProcess(command[0], command[1:]...).FireAndForget()
	return Result("la")
}

type Applications struct {
	desktopApps Actions
}

func (a Applications) Category() Category {
	return Application
}

func (a Applications) FindActions(input string) Actions {
		return FilterAction(input, a.desktopApps)
}

func ApplicationsPluginNew() Applications {
	dirs := append([]string(nil), basedir.DataDirs...)
	dirs = append(dirs, basedir.DataHome)
	var apps Actions
	for _, dir := range dirs {

		dir = path.Join(dir, "applications")
		share_apps := system.LoadDesktopEntries(dir)
		for _, app := range share_apps {
			action := AppAction{
				name: strings.ToLower(app.Name),
				description: strings.TrimSpace(app.GenericName),
				cmd: app.Exec,
			}
			apps = append(apps, action)
		}
	}
	return Applications{
		desktopApps: apps,
	}
}
