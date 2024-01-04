package main

/*
	this parses the command line arguements

	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"time"
	arg "github.com/alexflint/go-arg"
	"go.wit.com/log"
	"go.wit.com/gui/debugger"
)

var args struct {
	Display string `arg:"env:DISPLAY"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug your dns settings"`
}

var NET log.LogFlag
var NOW log.LogFlag
var PROC log.LogFlag
var SPEW log.LogFlag
var CHANGE log.LogFlag

func init() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)

	NOW.B = false
	NOW.Name = "NOW"
	NOW.Subsystem = "cpdns"
	NOW.Desc = "temp debugging stuff"
	NOW.Register()

	NET.B = false
	NET.Name = "NET"
	NET.Subsystem = "cpdns"
	NET.Desc = "Network logging"
	NET.Register()

	PROC.B = false
	PROC.Name = "PROC"
	PROC.Subsystem = "cpdns"
	PROC.Desc = "/proc logging"
	PROC.Register()

	SPEW.B = false
	SPEW.Name = "SPEW"
	SPEW.Subsystem = "cpdns"
	SPEW.Desc = "spew logging"
	SPEW.Register()

	CHANGE.B = false
	CHANGE.Name = "CHANGE"
	CHANGE.Subsystem = "cpdns"
	CHANGE.Desc = "show droplet state changes"
	CHANGE.Register()

	if debugger.ArgDebug() {
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
	log.Sleep(me.artificialSleep)
}
