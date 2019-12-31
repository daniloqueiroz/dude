package cmd

import (
	"errors"
	"github.com/daniloqueiroz/dude/cmd/daemons"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Control daemons",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires the daemon name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if name == "powerd" {
			daemons.Powerd()
		} else if name == "trackerd" {
			daemons.TimeTrackd()
		} else if name == "sessiond" {
			daemons.Sessiond()
		} else if name == "displayd" {
			daemons.Displayd()
		}
	},
}
