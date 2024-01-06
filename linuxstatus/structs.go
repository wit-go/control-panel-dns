/* 
	figures out if your hostname is valid
	then checks if your DNS is setup correctly
*/

package linuxstatus

import (
	"net"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

var me *LinuxStatus

type LinuxStatus struct {
	ready	bool
	hidden	bool
	changed	bool

	parent	*gui.Node

	ifmap	map[int]*IFtype		// the current interfaces
	ipmap	map[string]*IPtype	// the current ip addresses

	window	*gadgets.BasicWindow
	group	*gui.Node
	grid	*gui.Node

	hostnameStatus	*gadgets.OneLiner
	hostname	*gadgets.OneLiner
	hostshort	*gadgets.OneLiner
	domainname	*gadgets.OneLiner
	resolver	*gadgets.OneLiner
	uid		*gadgets.OneLiner
	IPv4		*gadgets.OneLiner
	IPv6		*gadgets.OneLiner
	workingIPv6	*gadgets.OneLiner
	Interfaces	*gui.Node
	speed		*gadgets.OneLiner
	speedActual	*gadgets.OneLiner

}

type IPtype struct {
	gone		bool		// used to track if the ip exists
	ipv6		bool		// the future
	ipv4		bool		// the past
	LinkLocal	bool
	iface		*net.Interface
	ip		net.IP
	ipnet		*net.IPNet
}

type IFtype struct {
	gone		bool		// used to track if the interface exists
	name		string		// just a shortcut to the name. maybe this is dumb
	// up		bool		// could be used to track ifup/ifdown
	iface		*net.Interface
}
