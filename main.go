package main

import (
	"github.com/pterm/pterm"
	"github.com/thenorthnate/pf/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			pterm.FgLightRed.Printfln("%v", r)
		}
	}()
	cmd.Execute()
}
