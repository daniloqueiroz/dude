package system

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

type Xorg struct {
	conn *xgb.Conn
	root xproto.Window
}

func NewXorg(display *string) *Xorg {
	var conn *xgb.Conn
	var err error
	if display != nil {
		conn, err = xgb.NewConnDisplay(*display)
		if err != nil {
			log.Fatal("xgb: ", err)
		}
	} else {
		conn, err = xgb.NewConn()
		if err != nil {
			log.Fatal("xgb: ", err)
		}
	}
	err = randr.Init(conn)
	if err != nil {
		logger.Fatalf("Unable to Initialize randr: %v", err)
	}
	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root
	return &Xorg{
		conn: conn,
		root: root,
	}
}

func (x *Xorg) Close() {
	x.conn.Close()
}

func (x *Xorg) Subscribe(chn chan xgb.Event, eventMask int) {
	if eventMask > 0 {
		err := randr.SelectInputChecked(x.conn, x.root, uint16(eventMask)).Check()
		if err != nil {
			logger.Fatalf("Unable to register for randr events: %v", err)
		}
	}

	for {
		ev, err := x.conn.WaitForEvent()
		if err != nil {
			logger.Errorf("Error on Event from X11 %v", err)
		}
		chn <- ev
	}
}

func (x *Xorg) DetectOutputs() []*Screen {
	var screens []*Screen

	resources, err := randr.GetScreenResources(x.conn, x.root).Reply()
	if err != nil {
		log.Fatal(err)
	}

	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(x.conn, output, 0).Reply()
		if err != nil {
			log.Fatal(err)
		}

		if info.Connection == randr.ConnectionConnected {
			bestMode := info.Modes[0]
			for _, mode := range resources.Modes {
				if mode.Id == uint32(bestMode) {
					screens = append(screens, &Screen{
						Name:        string(info.Name),
						BestMode:    fmt.Sprintf("%dx%d", mode.Width, mode.Height),
						IsConnected: true,
					})
				}
			}
		} else {
			screens = append(screens, &Screen{
				Name:        string(info.Name),
				BestMode:    "",
				IsConnected: false,
			})
		}
	}
	return screens
}
