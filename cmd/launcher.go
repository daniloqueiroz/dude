package cmd

import (
	"github.com/daniloqueiroz/dude/app/laucher/controller"
	"github.com/daniloqueiroz/dude/app/laucher/gtk"
	"github.com/spf13/cobra"
)

var launcherCmd = &cobra.Command{
	Use:   "launcher",
	Short: "dude Launcher",
	Run: func(cmd *cobra.Command, args []string) {
		launcher := controller.Launcher{}
		view := gtk.NewGtkView(&launcher)
		launcher.Start(view)
	},
}
