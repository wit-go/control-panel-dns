// This creates a simple hello world window
package main

import 	(
	"net"
)

// It's probably a terrible idea to call this 'me'
var me Host

type Host struct {
	hostname	string			// mirrors
	domainname	string			// kernel.org
	fqdn		string			// mirrors.kernel.org
	ipmap		map[string]*IPtype	// the current ip addresses
	ifmap		map[int]*IFtype		// the current interfaces
	ipchange	bool			// set to true if things change
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
