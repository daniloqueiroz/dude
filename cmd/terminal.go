package cmd

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/spf13/cobra"
)

var terminalCmd = &cobra.Command{
	Use:   "terminal",
	Short: "Launch a Terminal",
	Run: func(cmd *cobra.Command, args []string) {
		internal.NewTmuxTerminal()
	},
}
