package cmd

import (
	"errors"
	"fmt"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var backlightCmd = &cobra.Command{
	Use:   "backlight",
	Short: "Change backlight",
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
			logger.Fatalf("Unable to adjust backlight %v", err)
		} else {
			backlight, err := internal.GetBacklight()
			if err != nil {
				logger.Errorf("Unable to get backlight info")
			} else {
				value, _ := strconv.ParseFloat(strings.TrimSpace(backlight), 64)
				system.SimpleNotification(fmt.Sprintf("Backlight set to %.0f%%", value)).Show()
			}
		}
	},
}

func adjust(inc bool) error {
	return internal.AdjustBacklight(10, inc)
}

func setValue(param string) error {
	value, err := strconv.Atoi(param)
	if err != nil {
		return err
	} else {
		return internal.SetBacklight(value)
	}
}