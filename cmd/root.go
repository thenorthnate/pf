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
	verbose  = false
	dataDir  = ""
	config   = &Config{}
	accounts = &AccountFile{}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Set the verbosity of logs (default is not verbose)")
	rootCmd.PersistentFlags().StringVarP(&dataDir, "dir", "d", "", "Specify the path to the data directory")
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
		makeAddCommand(),
	)
}
