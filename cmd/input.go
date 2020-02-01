package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/input"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
)

var selectedKeyboard string
var inputCmd = &cobra.Command{
	Use:   "input [--keyboard=<Keyboard Name>]",
	Short: "Adjust input",
	Run: func(cmd *cobra.Command, args []string) {
		var kb string
		if selectedKeyboard != "" {
			err := input.SetKeyboard(selectedKeyboard)
			if err != nil {
				logger.Fatalf("Error configuring keyboard %s: %v", selectedKeyboard, err)
			}
			kb = selectedKeyboard
		} else {
			defaultKb, err := input.SetDefaultKeyboard()
			if err != nil {
				logger.Fatalf("Error configuring default keyboard: %v", err)
			}
			kb = defaultKb.Name
		}
		system.SimpleNotification(fmt.Sprintf("Keyboard switched to %s", kb)).Show()
	},
}
