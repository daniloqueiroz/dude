package pkg

import (
	"encoding/json"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"io/ioutil"
	"os"
)

type Tracker struct {
	trackingDir    string
	display        string
	currentTracker *gone.Recorder
}

func NewTracker(trackingDir string) *Tracker {
	return &Tracker{
		trackingDir: trackingDir,
		display:     os.Getenv("DISPLAY"),
	}
}

func (t *Tracker) Track() {
	tracker, err := gone.NewTracker(t.display, t.trackingDir)
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	t.currentTracker = tracker
	t.currentTracker.Start()
}

func LoadReport(week string) *gone.Report {
	reportFile := gone.ReportFileName(internal.Config.ScreenTimeDataDir, week)

	byteValue, err := ioutil.ReadFile(reportFile)
	if err != nil {
		logger.Fatalf("Unable to parse report file %s: %v", reportFile, err)
	}
	var report gone.Report
	json.Unmarshal(byteValue, &report)
	return &report
}