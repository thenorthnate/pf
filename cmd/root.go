package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pf",
		Short: "pf is a simple command line tool to manage your personal finances",
		Long:  `TBD`,
	}
	verbose    = false
	configPath = ""
	config     = &Config{}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Set the verbosity of logs (default is not verbose)")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Specify the path to the config file")
	registerCommands()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func registerCommands() {
	rootCmd.AddCommand(
		makeInvestCommand(),
	)
}
