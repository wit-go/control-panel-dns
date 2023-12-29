package main

/*
	this parses the command line arguements

	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"log"
	"time"
	arg "github.com/alexflint/go-arg"
	"go.wit.com/gui"
)

var args struct {
	Display string `arg:"env:DISPLAY"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug your dns settings"`
}

func init() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)

	if gui.ArgDebug() {
		log.Println(true, "INIT() gui debug == true")
	} else {
		log.Println(true, "INIT() gui debug == false")
	}

	me.dnsSleep = 500 * time.Millisecond
	me.localSleep = 100 * time.Millisecond

	me.artificialSleep = 0.4	// seems to need to exist or GTK crashes. TODO: fix andlabs plugin
	me.artificialS = "blah"
	log.Println("init() me.artificialSleep =", me.artificialSleep)
	log.Println("init() me.artificialS =", me.artificialS)
	sleep(me.artificialSleep)
}
