package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
)

var selectedProfile string
var displayCmd = &cobra.Command{
	Use:   "display [--profile=<Display Profile>]",
	Short: "Adjust displays - if no profile is specified auto detect output(s) and tries to match a profile",
	Run: func(cmd *cobra.Command, args []string) {
		var profile string
		if selectedProfile != "" {
			err := display.ApplyProfile(selectedProfile)
			if err != nil {
				logger.Fatalf("Unable to apply profile %s", selectedProfile)
			}
			profile = selectedProfile
		} else {
			profile = display.AutoConfigureDisplay()
		}
		system.TitleNotification("displayd", fmt.Sprintf("Profile %s applied", profile)).Show()
	},
}
