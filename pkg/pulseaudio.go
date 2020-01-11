package pkg

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
)

func VolumeUp() error {
	return proc.NewProcess(internal.Config.AppPactl, "set-sink-volume", "@DEFAULT_SINK@", "+2%").FireAndWait()
}

func VolumeDown() error {
	return proc.NewProcess(internal.Config.AppPactl, "set-sink-volume", "@DEFAULT_SINK@", "-2%").FireAndWait()
}

func VolumeMuteToggle() error {
	return proc.NewProcess(internal.Config.AppPactl, "set-sink-mute", "@DEFAULT_SINK@", "toggle").FireAndWait()
}

func MicMuteToggle() error {
	return proc.NewProcess(internal.Config.AppPactl, "set-source-mute", "@DEFAULT_SOURCE@", "toggle").FireAndWait()
}


