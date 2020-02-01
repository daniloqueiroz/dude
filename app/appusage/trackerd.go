package appusage

import (
	"encoding/json"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"os"
	"path"
)

func TrackAppUsage() {
	trackingDir := path.Join(basedir.DataHome, system.Config.AppUsageDataDir)
	display := os.Getenv("DISPLAY")
	recorder, err := NewRecorder(display, trackingDir)
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	recorder.Start()
}

func LoadReport(week string) *Report {
	reportFile := ReportFileName(path.Join(basedir.DataHome, system.Config.AppUsageDataDir), week)

	byteValue, err := ioutil.ReadFile(reportFile)
	if err != nil {
		logger.Fatalf("Unable to parse report file %s: %v", reportFile, err)
	}
	var report Report
	json.Unmarshal(byteValue, &report)
	return &report
}