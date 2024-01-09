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
	TmpLog bool  `arg:"--tmp-log" help:"automatically send STDOUT to /tmp"`
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

func init() {
	arg.MustParse(&args)
	full := "go.wit.com/control-panels/dns"
	short := "cpdns"

	NOW.NewFlag( "NOW",  true,  full, short, "temp debugging stuff")
	INFO.NewFlag("INFO", false, full, short, "normal debugging stuff")
	NET.NewFlag( "NET",  false, full, short, "Network logging")
	DNS.NewFlag( "DNS",  false, full, short, "dnsStatus.update()")

	WARN.NewFlag("WARN", true,  full, short, "bad things")
	SPEW.NewFlag("SPEW", false, full, short, "spew stuff")

	CHANGE.NewFlag("CHANGE", true,  full, short, "when host or dns change")
	STATUS.NewFlag("STATUS", false, full, short, "updateStatus() polling")

	if debugger.ArgDebug() {
		log.Log(NOW, "INIT() gui debug == true")
	} else {
		log.Log(NOW, "INIT() gui debug == false")
	}

	me.dnsSleep = 500 * time.Millisecond
	me.localSleep = 100 * time.Millisecond

	me.artificialSleep = 0.4	// seems to need to exist or GTK crashes. TODO: fix andlabs plugin
	me.artificialS = "blah"
	log.Log(INFO, "init() me.artificialSleep =", me.artificialSleep)
	log.Log(INFO, "init() me.artificialS =", me.artificialS)
	log.Sleep(me.artificialSleep)
}
