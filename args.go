// This creates a simple hello world window
package main

import 	(
	"git.wit.org/wit/gui"
)

type LogOptions struct {
	LogFile string
	Verbose bool
	// GuiDebug bool `help:"open up the wit/gui Debugging Window"`
	// GuiDemo bool `help:"open the wit/gui Demo Window"`
	User string `arg:"env:USER"`
}

var args struct {
	LogOptions
	gui.GuiArgs
}
