package cmd

import (
	"github.com/daniloqueiroz/dude/app/session"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Session daemon",
	Run: func(cmd *cobra.Command, args []string) {
		proc.CreatePidFile("dude-session")
		session := session.NewSession()
		session.Start()
		logger.Info("Session ended")
	},
}
