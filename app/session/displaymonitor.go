package session

import (
	"context"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/supervisor"
	"github.com/google/logger"
	"strings"
	"time"
)

func displayMonitorSupervisor(supervisor *supervisor.Supervisor) {
	chn := make(chan xgb.Event)
	xorg := system.NewXorg(nil)
	supervisor.AddTask("DisplayMonitor", func(ctx context.Context) error {
		time.Sleep(2 * time.Second)
		display.AutoConfigureDisplay()

		var lastEvent int64
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-chn:
				logger.Infof("Xrandr event received: %#v", event.String())
				now := time.Now().Unix()
				timeSinceLast := now - lastEvent
				if !strings.HasPrefix(event.String(), "MappingNotify") && timeSinceLast > 5 {
					display.AutoConfigureDisplay()
				}
				lastEvent = now
			}
		}
	})
	supervisor.AddTask("DisplayMonitorSubscriber", func(ctx context.Context) error {
		eventMask := randr.NotifyMaskScreenChange |
			randr.NotifyMaskCrtcChange |
			randr.NotifyMaskOutputChange |
			randr.NotifyMaskOutputProperty
		xorg.Subscribe(chn, eventMask)
		return nil
	})
}
