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

func init() {
	full := "go.wit.com/gui/gadgets/smartwindow"
	short := "smartWin"

	NOW.NewFlag( "NOW",  true,  full, short, "temp debugging stuff")
	INFO.NewFlag("INFO", false, full, short, "General Info")
	SPEW.NewFlag("SPEW", false, full, short, "spew stuff")
	WARN.NewFlag("WARN", false, full, short, "bad things")
}
