package cmd

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "Suspend computer",
	Run: func(cmd *cobra.Command, args []string) {
		system.Suspend()
	},
}
