package cmd

import (
	"os"

	"github.com/pterm/pterm"
)

func Fatal(message string) {
	pterm.FgLightRed.Println(message)
	os.Exit(1)
}

func Fatalf(message string, args ...any) {
	pterm.FgLightRed.Printfln(message, args...)
	os.Exit(1)
}
