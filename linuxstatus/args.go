package linuxstatus

/*
	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"go.wit.com/log"
)

var NOW log.LogFlag
var INFO log.LogFlag
var NET log.LogFlag
var DNS log.LogFlag
var PROC log.LogFlag
var SPEW log.LogFlag
var WARN log.LogFlag
var CHANGE log.LogFlag
var STATUS log.LogFlag

func myreg(f *log.LogFlag, b bool, name string, desc string) {
	f.B = b
	f.Subsystem = "go.wit.com/control-panels/dns/linuxstatus"
	f.Short = "linux"
	f.Desc = desc
	f.Name = name
	f.Register()
}

func init() {
	myreg(&NOW,    true,  "NOW",    "temp debugging stuff")
	myreg(&INFO,   false, "INFO",   "normal debugging stuff")
	myreg(&NET,    false, "NET",    "Network Logging")
	myreg(&DNS,    false, "DNS",    "dnsStatus.update()")
	myreg(&PROC,   false, "PROC",   "/proc logging")
	myreg(&SPEW,   false, "SPEW",   "spew stuff")
	myreg(&WARN,   true,  "WARN",   "bad things")
	myreg(&CHANGE, true,  "CHANGE", "show droplet state changes")
	myreg(&STATUS, false, "STATUS", "Update() details")
}
