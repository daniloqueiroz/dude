package cmd

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"syscall"
	"time"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "initialize DE session",
	Run: func(cmd *cobra.Command, args []string) {
		wd := proc.NewWatchdog()
		internal.StartCompositor(wd)
		internal.StartScreensaver(wd)
		internal.StartPolkit(wd)
		internal.SetWallpaper()
		internal.AutostartApps()
		proc.DaemonExec(wd, "powerd")
		proc.DaemonExec(wd, "trackerd")
		system.SimpleNotification("Session started").Show()

		wd.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		time.Sleep(1 * time.Second)
		logger.Info("Session ended")
		// TODO IPC -> processes status
	},
}
