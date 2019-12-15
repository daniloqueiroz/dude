package internal

import (
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"math"
)

type PowerState string

const (
	AC_ONLINE   PowerState = "ac_online"
	DISCHARGING PowerState = "discharging"
	LOW         PowerState = "low"
	VERY_LOW    PowerState = "very_low"
	CRITICAL    PowerState = "critical"
)

func CheckBattery() PowerState {
	if !system.IsOnAC() {
		var level float64 = 0
		bats := system.GetBatteries()
		for _, bat := range bats {
			level = math.Max(level, float64(bat.Level))
		}
		intLevel := int(level)
		if intLevel < 20 {
			return LOW
		} else if intLevel < 15 {
			return VERY_LOW
		} else if intLevel < 10 {
			return CRITICAL
		} else {
			return DISCHARGING
		}
	} else {
		return AC_ONLINE
	}
}
