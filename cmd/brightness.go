package cmd

import (
	"errors"
	"fmt"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var brightnessCmd = &cobra.Command{
	Use:   "brightness",
	Short: "Change brightness",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("[up|down]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		op := args[0]
		if op == "up" {
			err = adjust(true)
		} else if op == "down" {
			err = adjust(false)
		} else {
			err = setValue(op)
		}

		if err != nil {
			logger.Fatalf("Unable to adjust brightness %v", err)
		} else {
			brightness, err := display.GetBrightness()
			if err != nil {
				logger.Errorf("Unable to get brightness info", err)
			} else {
				value, _ := strconv.ParseFloat(strings.TrimSpace(brightness), 64)
				system.SimpleNotification(fmt.Sprintf("Brightness set to %.0f%%", value)).Show()
			}
		}
	},
}

func adjust(inc bool) error {
	return display.AdjustBrightness(10, inc)
}

func setValue(param string) error {
	value, err := strconv.Atoi(param)
	if err != nil {
		return err
	} else {
		return display.SetBrightness(value)
	}
}