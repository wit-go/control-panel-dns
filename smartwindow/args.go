package smartwindow

/*
	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"go.wit.com/log"
)

var NOW log.LogFlag
var INFO log.LogFlag
var SPEW log.LogFlag
var WARN log.LogFlag

func myreg(f *log.LogFlag, b bool, name string, desc string) {
	f.B = b
	f.Subsystem = "go.wit.com/gadgets/smartwindow"
	f.Short = "smartWin"
	f.Desc = desc
	f.Name = name
	f.Register()
}

func init() {
	myreg(&NOW,    true,  "NOW",    "temp debugging stuff")
	myreg(&INFO,   false, "INFO",   "normal debugging stuff")
	myreg(&SPEW,   false, "SPEW",   "spew stuff")
	myreg(&WARN,   true,  "WARN",   "bad things")
}
