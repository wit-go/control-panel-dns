// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/gui/gadgets"
)

// creates the actual widgets.
// it's assumed you are always passing in a box
func draw(ls *LinuxStatus) {
	if ! ls.Ready() {return}
	ls.group = ls.window.Box().NewGroup("Real Stuff")

	ls.grid = ls.group.NewGrid("gridnuts", 2, 2)

	ls.grid.SetNext(1,1)

	ls.hostshort	= gadgets.NewOneLiner(ls.grid, "hostname -s")
	ls.domainname	= gadgets.NewOneLiner(ls.grid, "domain name")
	ls.NSrr		= gadgets.NewOneLiner(ls.grid, "NS records =")
	ls.uid		= gadgets.NewOneLiner(ls.grid, "UID =")
	ls.IPv4		= gadgets.NewOneLiner(ls.grid, "Current IPv4 =")
	ls.IPv6		= gadgets.NewOneLiner(ls.grid, "Current IPv6 =")
	ls.workingIPv6	= gadgets.NewOneLiner(ls.grid, "Real IPv6 =")
	// ls.nics		= gadgets.NewOneLiner(ls.grid, "network intefaces =")

	ls.grid.NewLabel("interfaces =")
	ls.Interfaces = ls.grid.NewCombobox("Interfaces")

	ls.speedActual	= gadgets.NewOneLiner(ls.grid, "refresh speed =")

	ls.grid.Margin()
	ls.grid.Pad()
}