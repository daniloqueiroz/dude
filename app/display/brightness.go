package display

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"strconv"
	"strings"
)

func GetBrightness() (string, error) {
	process := proc.NewProcess(system.Config.AppBrightness, "get")
	currentBrightness, err := process.FireAndWaitForOutput()
	if err != nil {
		return "", err
	}
	process = proc.NewProcess(system.Config.AppBrightness, "max")
	maxBrightness, err := process.FireAndWaitForOutput()
	if err != nil {
		return "", err
	}
	currentBrightnessInt, err := strconv.ParseFloat(strings.TrimSpace(currentBrightness), 64)
	if err != nil {
		return "", err
	}
	maxBrightnessInt, err := strconv.ParseFloat(strings.TrimSpace(maxBrightness), 64)
	if err != nil {
		return "", err
	}
	// [VALUE]^[K] * [MAX] * 100^-[K]
	// where default K is 4
	return fmt.Sprintf("%.2f", currentBrightnessInt * 100 / maxBrightnessInt), nil
}

func SetBrightness(value int) error {
	percentage := strconv.Itoa(value) + "%"
	process := proc.NewProcess(system.Config.AppBrightness, "set", percentage)
	return process.FireAndWait()
}

func AdjustBrightness(delta int, inc bool) error {
	var param string
	if inc {
		param = "+" + strconv.Itoa(delta) + "%"
	} else {
		param = strconv.Itoa(delta) + "-%"
	}
	process := proc.NewProcess(system.Config.AppBrightness, "set", param)
	return process.FireAndWait()
}
