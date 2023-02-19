// This creates a simple hello world window
package main

import 	(
	"git.wit.org/wit/gui"
)

type LogOptions struct {
	LogFile string `help:"write all output to a file"`
	Verbose bool
	VerboseNet bool  `arg:"--verbose-net" help:"debug network settings"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug dns settings"`
	// GuiDebug bool `help:"open up the wit/gui Debugging Window"`
	// GuiDemo bool `help:"open the wit/gui Demo Window"`
	User string `arg:"env:USER"`
}

var args struct {
	LogOptions
	gui.GuiArgs
}
