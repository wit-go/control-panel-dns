package main

import 	(
	"log"
	"reflect"
	witlog "go.wit.com/log"
)

var LogPrefix = "ipv6cp" // ipv6 control panel debugging line

// various debugging flags
var DEBUGON bool = true
var LogNow bool = true	// useful for active development
var LogError bool = true // probably always leave this one
var LogChange bool = true // turn on /proc debugging output

var LogInfo bool = false // general info
var LogNet bool = false // general network debugging
var LogProc bool = false // turn on /proc debugging output
var LogExec bool = false // turn on os.Exec() debugging

// var SPEW witlog.Spewt

// var log interface{}

/*
func log(a ...any) {
	witlog.Where = "wit/gui"
	witlog.Log(a...)
}
*/

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	debug(LogError, "got to log() exit")
	witlog.Exit(a...)
}

func debug(a ...any) {
	if (! DEBUGON) {
		return
	}

	if (a == nil) {
		return
	}
	var tbool bool
	if (reflect.TypeOf(a[0]) == reflect.TypeOf(tbool)) {
		if (a[0] == false) {
			return
		}
		a[0] = LogPrefix // ipv6 control panel debugging line
	}

	log.Println(a...)
}
