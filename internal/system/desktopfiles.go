package system

import (
	"fmt"
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
