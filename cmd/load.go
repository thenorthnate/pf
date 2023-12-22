package cmd

import "github.com/spf13/cobra"

func makeLoadCommand() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "load transactions from a csv file",
		Long:  `TBD`,
		Run:   load,
	}
	return loadCmd
}

func load(cmd *cobra.Command, args []string) {

}
