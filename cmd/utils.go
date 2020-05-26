package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/appusage"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/rkoesters/xdg/basedir"
	"os"
	"path"
	"time"
)

func journalStore() *appusage.Journal {
	dataDir := path.Join(basedir.DataHome, system.Config.AppUsageDataDir)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	timeNow := time.Now()
	year, wkNum := timeNow.ISOWeek()
	currentWeekFile := fmt.Sprintf("%d-w%d.log", year, wkNum)
	return appusage.NewJornal(path.Join(dataDir, currentWeekFile), appusage.EventSerializer{})
}
