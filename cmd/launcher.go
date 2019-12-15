package cmd

import (
	"github.com/daniloqueiroz/dude/internal/laucher/ui"
	"github.com/spf13/cobra"
)

var launcherCmd = &cobra.Command{
	Use:   "launcher",
	Short: "dude Launcher",
	Run: func(cmd *cobra.Command, args []string) {
		launcherUI := ui.NewUI()
		launcherUI.Show()
	},
}
