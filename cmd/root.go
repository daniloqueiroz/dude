package cmd

import (
	"github.com/daniloqueiroz/dude/app/system"
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
	logger := logger.Init("dude", verbose, true, ioutil.Discard)
	defer logger.Close()

	system.InitConfig()

	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(lockCmd)
	rootCmd.AddCommand(suspendCmd)
	rootCmd.AddCommand(terminalCmd)
	rootCmd.AddCommand(backlightCmd)
	rootCmd.AddCommand(launcherCmd)
	rootCmd.AddCommand(appUsageCmd)

	displayCmd.Flags().StringVarP(&selectedProfile, "profile", "p", "", "Display profile to activate")
	rootCmd.AddCommand(displayCmd)

	inputCmd.Flags().StringVarP(&selectedKeyboard, "keyboard", "k", "", "Display profile to activate")
	rootCmd.AddCommand(inputCmd)

	rootCmd.AddCommand(audioCmd)
}
