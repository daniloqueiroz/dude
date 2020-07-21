package cmd

import (
	"github.com/daniloqueiroz/dude/app/re_launcher"
	"github.com/daniloqueiroz/dude/app/re_launcher/view"
	"github.com/spf13/cobra"
)

var launcherCmd = &cobra.Command{
	Use:   "launcher",
	Short: "dude Launcher",
	Run: func(cmd *cobra.Command, args []string) {
		p := re_launcher.PresenterNew(view.ViewNew())
		p.Init()
	},
}
