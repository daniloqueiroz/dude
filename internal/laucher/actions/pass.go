package actions

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/daniloqueiroz/dude/internal/laucher"
	"github.com/google/logger"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

const (
	PASS_DIR    = ".password-store"
	PASS_PREFIX = "@"
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
	output, err := proc.NewProcess(commons.Config.AppPass, "show", p.name).FireAndWaitForOutput()
	if err != nil {
		logger.Errorf("Failed to run pass to get password for %s", p.name, err)
	} else {
		password := strings.Fields(output)[0]
		proc.NewProcess(commons.Config.AppXdotool, "type", password).FireAndWait()
	}
}

func (p Pass) String() string {
	return p.Input()
}

func loadPasswordsActions(actions map[string]laucher.Action) {
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
