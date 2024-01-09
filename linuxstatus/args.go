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

func init() {
	full := "go.wit.com/control-panels/dns/linuxstatus"
	short := "linux"

	NOW.NewFlag( "NOW",  true,  full, short, "temp debugging stuff")
	INFO.NewFlag("INFO", false, full, short, "normal debugging stuff")
	NET.NewFlag( "NET",  false, full, short, "Network logging")
	DNS.NewFlag( "DNS",  false, full, short, "dnsStatus.update()")

	PROC.NewFlag("PROC", false, full, short, "/proc loggging")
	WARN.NewFlag("WARN", true,  full, short, "bad things")
	SPEW.NewFlag("SPEW", false, full, short, "spew stuff")

	CHANGE.NewFlag("CHANGE", true,  full, short, "when host or dns change")
	STATUS.NewFlag("STATUS", false, full, short, "Update() details")
}
