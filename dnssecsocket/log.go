package dnssecsocket

import 	(
	witlog "go.wit.com/gui/log"
)

// various debugging flags
var logNow bool = true	// useful for active development
var logError bool = true
var logWarn bool = false
var logInfo bool = false
var logVerbose bool = false

var SPEW witlog.Spewt

// var log interface{}

func log(a ...any) {
	witlog.Where = "wit/gui"
	witlog.Log(a...)
}

func sleep(a ...any) {
	witlog.Sleep(a...)
}

func exit(a ...any) {
	log(logError, "got to log() exit")
	witlog.Exit(a...)
}
