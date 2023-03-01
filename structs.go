// This creates a simple hello world window
package main

import 	(
	"net"
	"git.wit.org/wit/gui"
)

// It's probably a terrible idea to call this 'me'
var me Host

type Host struct {
	hostname	string			// mirrors
	domainname	string			// kernel.org
	fqdn		string			// mirrors.kernel.org
	dnsTTL		int			// Recheck DNS is working every TTL (in seconds)
	user		string			// name of the user
	ipmap		map[string]*IPtype	// the current ip addresses
	dnsmap		map[string]*IPtype	// the current dns addresses
	ifmap		map[int]*IFtype		// the current interfaces
	ipchange	bool			// set to true if things change
	window		*gui.Node		// the main window
	tab		*gui.Node		// the main dns tab
	notes		*gui.Node		// using this to put notes here
	output		*gui.Node		// Textbox for dumping output
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
