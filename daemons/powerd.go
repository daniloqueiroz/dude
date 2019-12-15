package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/google/logger"
	"time"
)

func Powerd() {
	logger.Info("powerd is running")
	var notifiedLow = false
	var stateChanged = false
	var state = internal.CheckBattery()

	for {
		newstate := internal.CheckBattery()
		if state != internal.AC_ONLINE && state != internal.DISCHARGING {
			notify(state, &notifiedLow)
		}
		stateChanged = newstate != state
		if stateChanged {
			adjustBacklight(newstate)
		}
		state = internal.CheckBattery()
		time.Sleep(nextCheckDelay(state))
	}
}

func adjustBacklight(newState internal.PowerState) {
	if newState == internal.AC_ONLINE {
		internal.SetBacklight(commons.Config.BackLightAC)
	} else {
		internal.SetBacklight(commons.Config.BackLightBattery)
	}
}

func notify(level internal.PowerState, notifiedLow *bool) {
	switch level {
	case internal.LOW:
		err := system.TitleNotification("powerd", "Battery level low").Show()
		if err == nil {
			*notifiedLow = true
		}
	case internal.VERY_LOW:
		system.TitleNotification("powerd", "Battery level very low").Show()
	case internal.CRITICAL:
		err := system.TitleNotification("powerd", "Computer is going to be suspended").Show()
		if err == nil {
			time.Sleep(5 * time.Second)
		}
		system.Suspend()
	}
}

func nextCheckDelay(state internal.PowerState) time.Duration {
	if state != internal.AC_ONLINE {
		return 20 * time.Second
	} else {
		return 5 * time.Second
	}
}
