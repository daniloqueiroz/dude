package internal

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"strconv"
)

func GetBacklight() (string, error) {
	process := proc.NewProcess(commons.Config.AppBacklight, "-get")
	return process.FireAndWaitForOutput()
}

func SetBacklight(value int) error {
	process := proc.NewProcess(commons.Config.AppBacklight, "-set", strconv.Itoa(value))
	return process.FireAndWait()
}

func AdjustBacklight(delta int, inc bool) error {
	var param string
	if inc {
		param = "+"
	} else {
		param = "-"
	}
	return proc.NewProcess(commons.Config.AppBacklight, param, strconv.Itoa(delta)).FireAndWait()
}