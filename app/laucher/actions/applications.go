package actions

import (
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/rkoesters/xdg/basedir"
	"path"
	"strings"
)

type Applications struct {
	desktopApps laucher.Actions
}

func (a *Applications) Find(input string) laucher.Actions {
	if a.desktopApps == nil {
		a.loadApplicationActions()
	}
	return laucher.FilterAction(input, a.desktopApps)
}

func (a *Applications) loadApplicationActions() {
	dirs := append([]string(nil), basedir.DataDirs...)
	dirs = append(dirs, basedir.DataHome)
	var apps laucher.Actions
	for _, dir := range dirs {
		dir = path.Join(dir, "applications")
		share_apps := system.LoadDesktopEntries(dir)
		for _, app := range share_apps {
			action := laucher.Action{
				Details: laucher.ActionMeta{
					Name: strings.ToLower(app.Name),
					Description: strings.TrimSpace(app.GenericName),
					Category:laucher.Application,
				},
				Exec: func (){
					command := strings.Fields(app.Exec)
					proc.NewProcess(command[0], command[1:]...).FireAndForget()
				},
			}
			apps = append(apps, action)
		}
	}
	a.desktopApps = apps
}
