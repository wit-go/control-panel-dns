package linuxstatus

/*
	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"go.wit.com/log"
)

var NOW *log.LogFlag
var INFO *log.LogFlag
var NET *log.LogFlag
var DNS *log.LogFlag

var PROC *log.LogFlag
var SPEW *log.LogFlag
var WARN *log.LogFlag

var CHANGE *log.LogFlag
var STATUS *log.LogFlag

func init() {
	full := "go.wit.com/control-panels/dns/linuxstatus"
	short := "linux"

	NOW = log.NewFlag( "NOW",  true,  full, short, "temp debugging stuff")
	INFO = log.NewFlag("INFO", false, full, short, "normal debugging stuff")
	NET = log.NewFlag( "NET",  false, full, short, "Network logging")
	DNS = log.NewFlag( "DNS",  false, full, short, "dnsStatus.update()")

	PROC = log.NewFlag("PROC", false, full, short, "/proc loggging")
	WARN = log.NewFlag("WARN", true,  full, short, "bad things")
	SPEW = log.NewFlag("SPEW", false, full, short, "spew stuff")

	CHANGE = log.NewFlag("CHANGE", true,  full, short, "when host or dns change")
	STATUS = log.NewFlag("STATUS", false, full, short, "Update() details")
}
