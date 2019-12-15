package cmd

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "initialize DE session",
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartCompositor()
		internal.SetWallpaper()
		internal.StartScreensaver()
		internal.StartPolkit()
		internal.AutostartApps()
		proc.DaemonExec("powerd")
		proc.DaemonExec("trackerd")
		system.SimpleNotification("Session started").Show()
		select { }
	},
}
