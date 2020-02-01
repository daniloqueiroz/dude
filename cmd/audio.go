package cmd

import (
	"errors"
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"strings"
)

var audioCmd = &cobra.Command{
	Use:   "audio [vol-up, vol-down, vol-mute, mic-mute]",
	Short: "Adjust audio",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var notification system.NotificationEvent
		var err error
		switch op :=  strings.ToLower(args[0]); op {
		case "vol-up":
			notification = system.SimpleNotification("Volume Up")
			err = app.VolumeUp()
		case "vol-down":
			notification = system.SimpleNotification("Volume Down")
			err = app.VolumeDown()
		case "vol-mute":
			notification = system.SimpleNotification("Volume mute toggled")
			err = app.VolumeMuteToggle()
		case "mic-mute":
			notification = system.SimpleNotification("Mic mute toggled")
			err = app.MicMuteToggle()
		default:
			err =  errors.New(fmt.Sprintf("invalid operation %s", op))
		}

		if err != nil {
			logger.Fatalf("Unable to perform audio operation: %v", err)
		} else {
			notification.Show()
		}
	},
}
