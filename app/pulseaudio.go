package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

func VolumeUp() error {
	return proc.NewProcess(system.Config.AppPactl, "set-sink-volume", "@DEFAULT_SINK@", "+2%").FireAndWait()
}

func VolumeDown() error {
	return proc.NewProcess(system.Config.AppPactl, "set-sink-volume", "@DEFAULT_SINK@", "-2%").FireAndWait()
}

func VolumeMuteToggle() error {
	return proc.NewProcess(system.Config.AppPactl, "set-sink-mute", "@DEFAULT_SINK@", "toggle").FireAndWait()
}

func MicMuteToggle() error {
	return proc.NewProcess(system.Config.AppPactl, "set-source-mute", "@DEFAULT_SOURCE@", "toggle").FireAndWait()
}


