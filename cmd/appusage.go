package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/appusage"
	"github.com/google/logger"
	"github.com/hako/durafmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var appUsageCmd = &cobra.Command{
	Use:   "app-usage",
	Short: "Show app usage report",
	Run: func(cmd *cobra.Command, args []string) {
		report, err := appusage.NewReport(journalStore())
		if err != nil {
			logger.Fatalf("Unable to load report: %v", err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "App", "Time Spent", "Percentage"})
		for idx, rec := range report.ClassRecords {
			table.Append([]string{
				strconv.Itoa(idx + 1),
				rec.Class,

				fmt.Sprintf("%s", durafmt.Parse(rec.Spent).LimitFirstN(2)),
				fmt.Sprintf("%.2f%%", rec.Percent),
			})
		}
		table.Render()
	},
}
