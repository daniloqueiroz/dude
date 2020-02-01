package system

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
	"regexp"
	"strconv"
	"strings"
)

type BatteryState string

const (
	Full        BatteryState = "Full"
	Discharging BatteryState = "Discharging"
	Charging    BatteryState = "Charging"
	Unknown     BatteryState = "Unknown"
)

type Battery struct {
	Name  string
	State BatteryState
	Level int
}

func (b Battery) String() string {
	return fmt.Sprintf("%s: %d%% - %s", b.Name, b.Level, b.State)
}

func GetBatteries() []Battery {
	output, err := proc.NewProcess(Config.AppAcpi, "-b").FireAndWaitForOutput()
	if err != nil {
		logger.Errorf("Unable get battery status using acpi: %v", err)
		return nil
	}

	var entries []Battery
	scanner := bufio.NewScanner(bytes.NewReader([]byte(output)))
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`(?P<Battery>.+): (?P<Status>\w+), (?P<Level>\d+)%.*`)
		details := r.FindStringSubmatch(line)
		level, _ := strconv.Atoi(details[3])
		entry := &Battery{
			Name:  details[1],
			State: BatteryState(details[2]),
			Level: level,
		}
		entries = append(entries, *entry)

	}
	return entries
}

func IsOnAC() bool {
	output, err := proc.NewProcess(Config.AppAcpi, "-a").FireAndWaitForOutput()
	if err != nil {
		logger.Error("Unable to get ac-adapter status", err)
		return false
	}
	return strings.TrimSpace(output) == "Adapter 0: on-line"

}
