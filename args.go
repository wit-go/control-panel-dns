package main

/*
	this parses the command line arguements
*/

import 	(
	"log"
	"fmt"
	"time"
	arg "github.com/alexflint/go-arg"
	"go.wit.com/gui"
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
}

func init() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar, args.User)

	if (args.Gui != "") {
		gui.GuiArg.Gui = args.Gui
	}
	log.Println(true, "INIT() args.GuiArg.Gui =", gui.GuiArg.Gui)

//	me.dnsTTL = 2		// how often to recheck DNS
//	me.dnsTTLsleep = 0.4	// sleep between loops

	me.dnsSleep = 500 * time.Millisecond
	me.localSleep = 100 * time.Millisecond

	me.artificialSleep = 0.4	// seems to need to exist or GTK crashes. TODO: fix andlabs plugin
	me.artificialS = "blah"
	log.Println("init() me.artificialSleep =", me.artificialSleep)
	log.Println("init() me.artificialS =", me.artificialS)
	sleep(me.artificialSleep)
}
