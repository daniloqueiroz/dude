package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/pkg"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var timereportCmd = &cobra.Command{
	Use:   "time-report",
	Short: "Show time report",
	Run: func(cmd *cobra.Command, args []string) {
		time_now := time.Now()
		year, wk_num := time_now.ISOWeek()
		currentWeek := fmt.Sprintf("%d-w%d", year, wk_num)

		report := pkg.LoadReport(currentWeek)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "App", "Time Spent", "Percentage"})
		for idx, rec := range report.ClassRecords {
			table.Append([]string{
				strconv.Itoa(idx + 1),
				rec.Class,
				fmt.Sprintf("%s", rec.Spent),
				fmt.Sprintf("%.2f%%", rec.Percent),
			})
		}
		table.Render()
	},
}
