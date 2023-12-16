package main

/*
	this parses the command line arguements
*/

import 	(
	"fmt"
	arg "github.com/alexflint/go-arg"
	"git.wit.org/wit/gui"
	log "git.wit.org/wit/gui/log"
)


var args struct {
	Verbose bool
	VerboseNet bool  `arg:"--verbose-net" help:"debug your local OS network settings"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug your dns settings"`
	LogFile string `help:"write all output to a file"`
	// User string `arg:"env:USER"`
	Display string `arg:"env:DISPLAY"`

	Foo string
	Bar bool
	User string `arg:"env:USER"`
	Demo bool `help:"run a demo"`
	gui.GuiArgs
	log.LogArgs
}

func init() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar, args.User)

	if (args.Gui != "") {
		gui.GuiArg.Gui = args.Gui
	}
	log.Log(true, "INIT() args.GuiArg.Gui =", gui.GuiArg.Gui)

}
