package display

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"strconv"
)

func GetBrightness() (string, error) {
	process := proc.NewProcess(system.Config.AppBacklight, "-get")
	return process.FireAndWaitForOutput()
}

func SetBrightness(value int) error {
	process := proc.NewProcess(system.Config.AppBacklight, "-set", strconv.Itoa(value))
	return process.FireAndWait()
}

func AdjustBrightness(delta int, inc bool) error {
	var param string
	if inc {
		param = "+"
	} else {
		param = "-"
	}
	return proc.NewProcess(system.Config.AppBacklight, param, strconv.Itoa(delta)).FireAndWait()
}