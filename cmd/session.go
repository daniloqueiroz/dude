package cmd

import (
	"github.com/daniloqueiroz/dude/app/appusage"
	"github.com/daniloqueiroz/dude/app/session"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"os"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Session daemon",
	Run: func(cmd *cobra.Command, args []string) {
		proc.CreatePidFile("dude-session")
		session := session.NewSession(getRecorder())
		session.Start()
		logger.Info("Session ended")
	},
}

func getRecorder() *appusage.Recorder {
	display := os.Getenv("DISPLAY")
	recorder, err := appusage.NewRecorder(display, journalStore())
	if err != nil {
		logger.Fatalf("Error starting time trackerd: %v", err)
	}
	return recorder
}