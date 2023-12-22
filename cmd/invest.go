package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func makeInvestCommand() *cobra.Command {
	investCmd := &cobra.Command{
		Use:   "invest",
		Short: "evaulate investments",
		Long:  `TBD`,
		Run:   invest,
	}
	return investCmd
}

func invest(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			Fatalf("could not parse amount: %v", err)
		}
		for fund, percentage := range config.Investments.Blend {
			investAmount := math.Round(amount * percentage)
			amountStr := pterm.FgLightBlue.Sprintf("$%.2f", investAmount)
			fmt.Printf("%-8s%v\n", fund+":", amountStr)
		}
	}
}
