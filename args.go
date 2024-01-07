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

var NOW log.LogFlag
var INFO log.LogFlag
var NET log.LogFlag
var DNS log.LogFlag
var WARN log.LogFlag
var SPEW log.LogFlag
var CHANGE log.LogFlag
var STATUS log.LogFlag

func myreg(f *log.LogFlag, b bool, name string, desc string) {
	f.B = b
	f.Subsystem = "go.wit.com/control-panels/dns"
	f.Short = "cpdns"
	f.Desc = desc
	f.Name = name
	f.Register()
}

func init() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)

	myreg(&NOW,    true,  "NOW",    "temp debugging stuff")
	myreg(&INFO,   false, "INFO",   "normal debugging stuff")
	myreg(&NET,    false, "NET",    "Network logging")
	myreg(&DNS,    false, "DNS",    "dnsStatus.update()")
	myreg(&WARN,   true,  "WARN",   "bad things")
	myreg(&SPEW,   false, "SPEW",   "spew stuff")
	myreg(&CHANGE, true,  "CHANGE", "when host or dns change")
	myreg(&STATUS, false, "STATUS", "updateStatus()")

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
