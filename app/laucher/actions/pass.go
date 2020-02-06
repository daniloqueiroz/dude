package actions

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
	PASS_PREFIX = "@"
	PASS_SCRIPT = "/usr/share/dude/scripts/passtype.sh"
)

type Pass struct {
	name        string
	description string
}

func (p Pass) Input() string {
	return fmt.Sprintf("%s%s", PASS_PREFIX, p.name)
}

func (p Pass) Description() string {
	return p.description
}

func (p Pass) Exec() {
	err := proc.NewProcess(PASS_SCRIPT, p.name).FireAndForget()
	if err != nil {
		logger.Errorf("Error launching passtype")
	}
}

func (p Pass) String() string {
	return p.Input()
}

func loadPasswordsActions(actions map[string]Action) {
	var entries []Pass

	dirname := filepath.Join(system.HomeDir(), PASS_DIR)
	entries = append(entries, loadPassFromDir(dirname, true)...)
	for _, pass := range entries {
		actions[pass.Input()] = pass
	}
}

func loadPassFromDir(dirname string, isRootDir bool) []Pass {
	var entries []Pass
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
				entries = append(entries, Pass{
					name:        pass,
					description: fmt.Sprintf("Password for %s", pass),
				})
			}
		}
	}
	return entries
}
