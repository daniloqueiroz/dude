package actions

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/daniloqueiroz/dude/internal/laucher"
	"github.com/rkoesters/xdg/basedir"
	"path"
	"strings"
)

type App struct {
	name        string
	command     []string
	description string
}

func (p App) Input() string {
	return p.name
}

func (p App) Description() string {
	return p.description
}

func (p App) Exec() {
	proc.NewProcess(p.command[0], p.command[1:]...).FireAndForget()
}

func (p App) String() string {
	return p.Input()
}


func loadApplicationActions(actions map[string]laucher.Action) {
	dirs := append([]string(nil), basedir.DataDirs...)
	dirs = append(dirs, basedir.DataHome)
	for _, dir := range dirs {
		dir = path.Join(dir, "applications")
		share_apps := internal.LoadDesktopEntries(dir)
		for _, app := range share_apps {
			action := App{
				name:        strings.ToLower(app.Name),
				command:     strings.Fields(app.Exec),
				description: strings.TrimSpace(app.GenericName),
			}
			actions[action.Input()] = action
		}
	}
}
