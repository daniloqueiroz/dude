package display

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/google/logger"
	"log"
)

type Screen struct {
	Name        string
	BestMode    string
	IsConnected bool
}

func getConnAndWindow() (*xgb.Conn, xproto.Window) {
	X, _ := xgb.NewConn()
	err := randr.Init(X)
	if err != nil {
		log.Fatal(err)
	}
	return X, xproto.Setup(X).DefaultScreen(X).Root
}

func DetectOutputs() []*Screen {
	var screens []*Screen
	X, root := getConnAndWindow()

	resources, err := randr.GetScreenResources(X, root).Reply()
	if err != nil {
		log.Fatal(err)
	}

	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(X, output, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}

		if info.Connection == randr.ConnectionConnected {
			bestMode := info.Modes[0]
			for _, mode := range resources.Modes {
				if mode.Id == uint32(bestMode) {
					screens = append(screens, &Screen{
						Name:     string(info.Name),
						BestMode: fmt.Sprintf("%dx%d", mode.Width, mode.Height),
						IsConnected: true,
					})
				}
			}
		} else {
			screens = append(screens, &Screen{
				Name:     string(info.Name),
				BestMode: "",
				IsConnected: false,
			})
		}
	}
	return screens
}

func ListenOutputEvents(notifyChn chan xgb.Event) {
	X, root := getConnAndWindow()
	err := randr.SelectInputChecked(X, root,
		randr.NotifyMaskScreenChange|
			randr.NotifyMaskCrtcChange|
			randr.NotifyMaskOutputChange|
			randr.NotifyMaskOutputProperty).Check()
	if err != nil {
		logger.Fatalf("Unable to register for randr events: %v", err)
	}

	for {
		ev, err := X.WaitForEvent()
		if err != nil {
			logger.Errorf("Error on Event from X11 %v", err)
		}
		notifyChn <- ev
	}
}
