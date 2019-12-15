package internal

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/desktop"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadDesktopEntries(dirname string) []desktop.Entry {
	var entries []desktop.Entry // an empty list

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		logger.Info("Unable to load desktop files")
		return entries
	}
	for _, f := range files {
		fullFilepath := filepath.Join(dirname, f.Name())
		file, err := os.Open(fullFilepath)
		if err != nil {
			logger.Error(fmt.Sprintf("Unable to load file %s", fullFilepath))
			continue
		}
		entry, err := desktop.New(file)
		if err != nil {
			continue
		}
		entries = append(entries, *entry)
	}
	return entries
}

func AutostartApps() {
	entries := LoadDesktopEntries(filepath.Join(system.HomeDir(), "/.config/autostart"))
	for _, entry := range entries {
		logger.Info(fmt.Sprintf("Autostarting %s", entry.Exec))
		proc.NewProcess(entry.Exec).FireAndForget()
	}
}
