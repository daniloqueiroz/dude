package cmd

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"io/ioutil"
	//"github.com/spf13/viper"
)

var (
	// Used for flags.
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "dude",
		Short: "danilo's unique desktop environment",
		Long:  `Dude is a Desktop Environment for X11 Window Managers`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose mode")

	defer logger.Init("dude", verbose, true, ioutil.Discard).Close()
	commons.InitConfig()

	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(sessionCmd)
	rootCmd.AddCommand(lockCmd)
	rootCmd.AddCommand(suspendCmd)
	rootCmd.AddCommand(terminalCmd)
	rootCmd.AddCommand(backlightCmd)
	rootCmd.AddCommand(launcherCmd)
}
