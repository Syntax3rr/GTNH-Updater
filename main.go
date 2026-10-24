package main

import (
	"github.com/Syntax3rr/GTNH-Updater/internal/cli"
	// "github.com/Syntax3rr/GTNH-Updater/internal/gui"

	"os"
)

func main() {
	if len(os.Args) == 1 {
		// os.Exit(gui.Run())
		os.Exit(0)
	}
	os.Exit(cli.Run(os.Args[1:]))
}
