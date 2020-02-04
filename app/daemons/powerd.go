package daemons

import (
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"syscall"
	"time"
)

func Powerd() {
	proc.CreatePidFile("powerd")
	daemon := proc.NewDaemon(monitorBattery)
	daemon.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
}

func monitorBattery() {
	logger.Info("Powerd is running")
	var notifiedLow = false
	var stateChanged = false
	var state = app.CheckBattery()

	for {
		newstate := app.CheckBattery()
		if state != app.AC_ONLINE && state != app.DISCHARGING {
			notify(state, &notifiedLow)
		}
		stateChanged = newstate != state
		if stateChanged {
			logger.Infof("Power state changed to %#v", newstate)
			adjustBacklight(newstate)
		}
		state = app.CheckBattery()
		time.Sleep(nextCheckDelay(state))
	}
}

func adjustBacklight(newState app.PowerState) {
	if newState == app.AC_ONLINE {
		display.SetBrightness(system.Config.BackLightAC)
	} else {
		display.SetBrightness(system.Config.BackLightBattery)
	}
}

func notify(level app.PowerState, notifiedLow *bool) {
	switch level {
	case app.LOW:
		err := system.TitleNotification("powerd", "Battery level low").Show()
		if err == nil {
			*notifiedLow = true
		}
	case app.VERY_LOW:
		system.TitleNotification("powerd", "Battery level very low").Show()
	case app.CRITICAL:
		err := system.TitleNotification("powerd", "Computer is going to be suspended").Show()
		if err == nil {
			time.Sleep(5 * time.Second)
		}
		system.Suspend()
	}
}

func nextCheckDelay(state app.PowerState) time.Duration {
	if state != app.AC_ONLINE {
		return 20 * time.Second
	} else {
		return 5 * time.Second
	}
}
