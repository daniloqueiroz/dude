package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

const (
	PASS_DIR    = ".password-store"
	PASS_SCRIPT = "/usr/share/dude/scripts/passtype.sh"
)

type PassAction struct {
	pass string
}

func (pa *PassAction) Category() Category {
	return Password
}
func (pa *PassAction) Name() string {
	return pa.pass
}
func (pa *PassAction) Description() string {
	return fmt.Sprintf("Password for %s", pa.pass)
}

func (pa *PassAction) Execute() Result {
	err := proc.NewProcess(PASS_SCRIPT, pa.pass).FireAndForget()
	if err != nil {
		logger.Errorf("Error launching passtype")
	}
	return Result("la")
}

type Pass struct {
	passwords Actions
}

func (p *Pass) Category() Category {
	return Password
}

func (p *Pass) FindActions(input string) Actions {
	return FilterAction(input, p.passwords)
}

func PassPluginNew() LauncherPlugin {
	dirname := filepath.Join(system.HomeDir(), PASS_DIR)
	return &Pass{
		passwords: loadPassFromDir(dirname, true),
	}

}

func loadPassFromDir(dirname string, isRootDir bool) Actions {
	var entries Actions
	if strings.HasSuffix(dirname, ".git") {
		return entries
	}

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		logger.Infof("Error loading files from dir %s", dirname)
		return entries
	}
	for _, f := range files {
		if f.IsDir() {
			entries = append(entries, loadPassFromDir(filepath.Join(dirname, f.Name()), false)...)
		} else {
			if strings.HasSuffix(f.Name(), ".gpg") {
				pass := strings.TrimSuffix(f.Name(), ".gpg")
				if !isRootDir {
					pass = filepath.Join(path.Base(dirname), pass)
				}
				entries = append(entries, &PassAction{pass: pass})
			}
		}
	}
	return entries
}
