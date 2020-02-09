package actions

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/laucher"
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

type Pass struct {
	passwords laucher.Actions
}

func (p *Pass) Find(input string) laucher.Actions {
	if p.passwords == nil {
		p.loadPasswordsActions()
	}
	return laucher.FilterAction(input, p.passwords)
}

func (p *Pass) loadPasswordsActions() {
	dirname := filepath.Join(system.HomeDir(), PASS_DIR)
	p.passwords = loadPassFromDir(dirname, true)
}

func loadPassFromDir(dirname string, isRootDir bool) laucher.Actions {
	var entries laucher.Actions
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
				entries = append(entries, laucher.Action{
					Details: laucher.ActionMeta{
						Name:        pass,
						Description: fmt.Sprintf("Password for %s", pass),
						Category:    laucher.Password,
					},
					Exec: wrapPass(pass),
				})
			}
		}
	}
	return entries
}

func wrapPass(pass string) (func ()) {
	return func ()  {
		err := proc.NewProcess(PASS_SCRIPT, pass).FireAndForget()
		if err != nil {
			logger.Errorf("Error launching passtype")
		}
	}
}