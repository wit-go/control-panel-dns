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
	// fqdn		string			// mirrors.kernel.org
	dnsTTL		int			// Recheck DNS is working every TTL (in seconds)
	changed		bool			// set to true if things changed
	user		string			// name of the user
	ipmap		map[string]*IPtype	// the current ip addresses
	dnsmap		map[string]*IPtype	// the current dns addresses
	ifmap		map[int]*IFtype		// the current interfaces
	window		*gui.Node		// the main window
	tab		*gui.Node		// the main dns tab
	notes		*gui.Node		// using this to put notes here
	uid		*gui.Node		// user
	fqdn		*gui.Node		// display the full hostname
	IPv4		*gui.Node		// show valid IPv4 addresses
	IPv6		*gui.Node		// show valid IPv6 addresses
	Interfaces	*gui.Node		// Interfaces
	DnsAAAA		*gui.Node		// the actual DNS AAAA results
	DnsA		*gui.Node		// the actual DNS A results (ignore for status since mostly never happens?)
	DnsStatus	*gui.Node		// the current state of DNS
	fix		*gui.Node		// button for the user to click
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
