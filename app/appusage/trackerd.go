package appusage

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"os"
	"path"
	"time"
)

func TrackAppUsage() {
	display := os.Getenv("DISPLAY")
	recorder, err := NewRecorder(display, journalStore())
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	recorder.Start()
}

func LoadReport() (*Report, error) {
	return NewReport(journalStore())
}

func journalStore() *Journal {
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
	return NewJornal(path.Join(dataDir, currentWeekFile), EventSerializer{})
}
