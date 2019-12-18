package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/gone"
	"github.com/google/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/rkoesters/xdg/basedir"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strconv"
)

var timereportCmd = &cobra.Command{
	Use:   "time-report",
	Short: "Show time report",
	Run: func(cmd *cobra.Command, args []string) {

		display := os.Getenv("DISPLAY")
		timeFile := path.Join(basedir.CacheHome, commons.Config.TimeTrackingFile)
		tracker, err := gone.NewTracker(display, timeFile)
		if err != nil {
			logger.Fatalf("Error loading tracking data: %v", err)
		}
		report := gone.GenerateReport(tracker)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "App", "Time Spent", "Percentage"})

		for idx, rec := range report.Classes{
			table.Append([]string{
				strconv.Itoa(idx+1),
				rec.Class,
				fmt.Sprintf("%s", rec.Spent),
				fmt.Sprintf("%.2f%%", rec.Percent),
			})
		}
		table.Render()
	},
}
