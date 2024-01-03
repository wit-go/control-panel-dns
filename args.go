package main

/*
	this parses the command line arguements

	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"time"
	arg "github.com/alexflint/go-arg"
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

var args struct {
	Display string `arg:"env:DISPLAY"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug your dns settings"`
}

var NET log.LogFlag
var SPEW log.LogFlag

func init() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)

	NET.B = false
	NET.Name = "NET"
	NET.Subsystem = "cpdns"
	NET.Desc = "Network logging"
	NET.Register()

	SPEW.B = false
	SPEW.Name = "SPEW"
	SPEW.Subsystem = "cpdns"
	SPEW.Desc = "spew logging"
	SPEW.Register()

	if gui.ArgDebug() {
		log.Log(true, "INIT() gui debug == true")
	} else {
		log.Log(true, "INIT() gui debug == false")
	}

	me.dnsSleep = 500 * time.Millisecond
	me.localSleep = 100 * time.Millisecond

	me.artificialSleep = 0.4	// seems to need to exist or GTK crashes. TODO: fix andlabs plugin
	me.artificialS = "blah"
	log.Log(true, "init() me.artificialSleep =", me.artificialSleep)
	log.Log(true, "init() me.artificialS =", me.artificialS)
	sleep(me.artificialSleep)
}
