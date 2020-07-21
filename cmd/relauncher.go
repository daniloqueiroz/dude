package cmd

import (
	"github.com/daniloqueiroz/dude/app/re_launcher"
	"github.com/daniloqueiroz/dude/app/re_launcher/view"
	"github.com/spf13/cobra"
)

var relauncherCmd = &cobra.Command{
	Use:   "relauncher",
	Short: "dude Launcher",
	Run: func(cmd *cobra.Command, args []string) {

		p := re_launcher.PresenterNew(view.ViewNew())
		p.Init()

	},
}
