package cmd

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/spf13/cobra"
	"log"
)

var displayCmd = &cobra.Command{
	Use:   "display [auto, single, mirror]",
	Short: "Adjust monitors",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO remove log
		X, _ := xgb.NewConn()

		// Every extension must be initialized before it can be used.
		err := randr.Init(X)
		if err != nil {
			log.Fatal(err)
		}

		// Get the root window on the default screen.
		root := xproto.Setup(X).DefaultScreen(X).Root

		// Gets the current screen resources. Screen resources contains a list
		// of names, crtcs, outputs and modes, among other things.
		resources, err := randr.GetScreenResources(X, root).Reply()
		if err != nil {
			log.Fatal(err)
		}

		// Iterate through all of the outputs and show some of their info.
		for _, output := range resources.Outputs {
			info, err := randr.GetOutputInfo(X, output, 0).Reply()
			if err != nil {
				log.Fatal(err)
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

		fmt.Println("\n")

		// Iterate through all of the crtcs and show some of their info.
		for _, crtc := range resources.Crtcs {
			info, err := randr.GetCrtcInfo(X, crtc, 0).Reply()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s -> X: %d, Y: %d, Width: %d, Height: %d\n",
				info.X, info.Y, info.Width, info.Height)
		}

		// Tell RandR to send us events. (I think these are all of them, as of 1.3.)
		err = randr.SelectInputChecked(X, root,
			randr.NotifyMaskScreenChange|
				randr.NotifyMaskCrtcChange|
				randr.NotifyMaskOutputChange|
				randr.NotifyMaskOutputProperty).Check()
		if err != nil {
			log.Fatal(err)
		}

		// Listen to events and just dump them to standard out.
		// A more involved approach will have to read the 'U' field of
		// RandrNotifyEvent, which is a union (really a struct) of type
		// RanrNotifyDataUnion.
		for {
			ev, err := X.WaitForEvent()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ev)
		}
	},
}
