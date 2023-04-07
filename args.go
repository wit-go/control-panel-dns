// This creates a simple hello world window
package main

import 	(
	"git.wit.org/wit/gui"
	"git.wit.org/jcarr/dnssecsocket"
)

type LogOptions struct {
	Verbose bool
	VerboseNet bool  `arg:"--verbose-net" help:"debug your local OS network settings"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug your dns settings"`
	LogFile string `help:"write all output to a file"`
	// User string `arg:"env:USER"`
	Display string `arg:"env:DISPLAY"`
}

var args struct {
	LogOptions
	dnssecsocket.Args
	gui.GuiArgs
}

func parsedown () {
	dnssecsocket.Parse(args.VerboseDnssec)
}
