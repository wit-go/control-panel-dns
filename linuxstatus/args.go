package linuxstatus

/*
	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	"go.wit.com/log"
)

var NOW log.LogFlag
var NET log.LogFlag
var DNS log.LogFlag
var PROC log.LogFlag
var SPEW log.LogFlag
var CHANGE log.LogFlag
var STATUS log.LogFlag

func init() {
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

	DNS.B = false
	DNS.Name = "DNS"
	DNS.Subsystem = "cpdns"
	DNS.Desc = "dnsStatus.update()"
	DNS.Register()

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

	STATUS.B = false
	STATUS.Name = "STATUS"
	STATUS.Subsystem = "cpdns"
	STATUS.Desc = "updateStatus()"
	STATUS.Register()
}
