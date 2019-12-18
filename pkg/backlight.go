package pkg

import (
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"strconv"
)

func GetBacklight() (string, error) {
	process := proc.NewProcess(internal.Config.AppBacklight, "-get")
	return process.FireAndWaitForOutput()
}

func SetBacklight(value int) error {
	process := proc.NewProcess(internal.Config.AppBacklight, "-set", strconv.Itoa(value))
	return process.FireAndWait()
}

func AdjustBacklight(delta int, inc bool) error {
	var param string
	if inc {
		param = "+"
	} else {
		param = "-"
	}
	return proc.NewProcess(internal.Config.AppBacklight, param, strconv.Itoa(delta)).FireAndWait()
}