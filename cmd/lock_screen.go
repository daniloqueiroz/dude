package cmd

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"os"
)

var lockCmd = &cobra.Command{
	Use:   "lock-screen",
	Short: "Lock Screen",
	Run: func(cmd *cobra.Command, args []string) {
		err := system.LockScreen()
		if err != nil {
			logger.Error("Unable to lock screen", err)
			fmt.Println("Unable to lock screen")
			os.Exit(-1)
		}
	},
}
