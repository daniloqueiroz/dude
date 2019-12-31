package daemons

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/internal/system"
	"github.com/daniloqueiroz/dude/pkg"
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
	var state = pkg.CheckBattery()

	for {
		newstate := pkg.CheckBattery()
		if state != pkg.AC_ONLINE && state != pkg.DISCHARGING {
			notify(state, &notifiedLow)
		}
		stateChanged = newstate != state
		if stateChanged {
			adjustBacklight(newstate)
		}
		state = pkg.CheckBattery()
		time.Sleep(nextCheckDelay(state))
	}
}

func adjustBacklight(newState pkg.PowerState) {
	if newState == pkg.AC_ONLINE {
		pkg.SetBacklight(internal.Config.BackLightAC)
	} else {
		pkg.SetBacklight(internal.Config.BackLightBattery)
	}
}

func notify(level pkg.PowerState, notifiedLow *bool) {
	switch level {
	case pkg.LOW:
		err := system.TitleNotification("powerd", "Battery level low").Show()
		if err == nil {
			*notifiedLow = true
		}
	case pkg.VERY_LOW:
		system.TitleNotification("powerd", "Battery level very low").Show()
	case pkg.CRITICAL:
		err := system.TitleNotification("powerd", "Computer is going to be suspended").Show()
		if err == nil {
			time.Sleep(5 * time.Second)
		}
		system.Suspend()
	}
}

func nextCheckDelay(state pkg.PowerState) time.Duration {
	if state != pkg.AC_ONLINE {
		return 20 * time.Second
	} else {
		return 5 * time.Second
	}
}
