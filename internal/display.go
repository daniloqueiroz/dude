package internal

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/google/logger"
)

func SetWallpaper() {
	logger.Info("Starting feh wallpaper")
	cmd := proc.NewProcess(commons.Config.AppFeh, "--bg-fill", "--randomize", "/home/danilo/.config/i3/wallpapers")
	cmd.FireAndForget()
}

func StartCompositor(wd *proc.Watchdog) {
	logger.Info("Starting compton compositor")
	cmd := proc.NewProcess(commons.Config.AppCompton, "-d", ":0")
	wd.Supervise(cmd)
}

func ConnectedOutputs() {
	X, _ := xgb.NewConn()

	// Every extension must be initialized before it can be used.
	err := randr.Init(X)
	if err != nil {
		logger.Fatal(err)
	}

	// Get the root window on the default screen.
	root := xproto.Setup(X).DefaultScreen(X).Root

	// Gets the current screen resources. Screen resources contains a list
	// of names, crtcs, outputs and modes, among other things.
	resources, err := randr.GetScreenResources(X, root).Reply()
	if err != nil {
		logger.Fatal(err)
	}

	// Iterate through all of the outputs and show some of their info.
	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(X, output, 0).Reply()
		if err != nil {
			logger.Fatal(err)
		}

		if info.Connection == randr.ConnectionConnected {
			bestMode := info.Modes[0]
			for _, mode := range resources.Modes {
				if mode.Id == uint32(bestMode) {
					fmt.Printf("Output: %s, Width: %d, Height: %d\n",
						info.Name, mode.Width, mode.Height)
				}
			}
		}
	}
}
