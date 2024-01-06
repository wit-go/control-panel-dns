// This creates a simple hello world window
package main

import 	(
	"net"
	"time"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/control-panels/dns/linuxstatus"

	"github.com/miekg/dns"
)

// It's probably a terrible idea to call this 'me'
var me Host

type Host struct {
	status		*hostnameStatus		// keeps track of the hostname and it's status
	statusOS	*linuxstatus.LinuxStatus		// what the Linux OS sees

	hostnameStatus	*gui.Node		// a summary for the user of where things are

	// domainname	*gui.Node		// kernel.org
	hostshort	*gui.Node		// hostname -s

	artificialSleep float64	`default:"0.7"`	// artificial sleep on startup
	artificialS     string 	`default:"abc"`	// artificial sleep on startup

	ttl		*gadgets.Duration
	dnsTtl		*gadgets.Duration
	dnsSleep	time.Duration
	localSleep	time.Duration

	changed		bool			// set to true if things changed
	user		string			// name of the user

	ipmap		map[string]*IPtype	// the current ip addresses
	dnsmap		map[string]*IPtype	// the current dns addresses
	ifmap		map[int]*IFtype		// the current interfaces
	nsmap		map[string]string	// the NS records

	// DNS A and AAAA results
	ipv4s		map[string]dns.RR
	ipv6s		map[string]dns.RR

	window		*gadgets.BasicWindow	// the main window
	details		*gadgets.BasicWindow	// more details of the DNS state
	debug		*gadgets.BasicWindow	// more attempts to debug the DNS state

	tab		*gui.Node		// the main dns tab
	notes		*gui.Node		// using this to put notes here

	// local OS settings, network interfaces, etc
	uid		*gui.Node		// user
	fqdn		*gui.Node		// display the full hostname
	IPv4		*gui.Node		// show valid IPv4 addresses
	IPv6		*gui.Node		// show valid IPv6 addresses
	Interfaces	*gui.Node		// Interfaces
	LocalSpeedActual *gui.Node		// the time it takes to check each network interface

	// DNS stuff
	NSrr		*gui.Node		// NS resource records for the domain name
	DnsAPI		*gui.Node		// what DNS API to use?
	DnsAAAA		*gadgets.OneLiner	// the actual DNS AAAA results
	workingIPv6	*gui.Node		// currently working AAAA
	DnsA		*gui.Node		// the actual DNS A results (ignore for status since mostly never happens?)
	DnsStatus	*gui.Node		// the current state of DNS
	DnsSpeed	*gui.Node		// 'FAST', 'OK', 'SLOW', etc
	DnsSpeedActual	*gui.Node		// the last actual duration
	DnsSpeedLast	string			// the last state 'FAST', 'OK', etc

	// fix		*gui.Node		// button for the user to click
	// fixProc		*gui.Node		// button for the user to click

	// mainStatus	*gui.Node		// group for the main display of stuff
	// cloudflareB	*gui.Node		// cloudflare button

	digStatus	*digStatus
	statusIPv6	*gadgets.OneLiner
	digStatusButton *gui.Node

	myDebug			*gui.Node
	myGui			*gui.Node
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
