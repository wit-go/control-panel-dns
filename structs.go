// This creates a simple hello world window
package main

import 	(
	"net"
	"time"
	"go.wit.com/gui"
	"go.wit.com/control-panel-dns/cloudflare"
	"github.com/miekg/dns"
)

// It's probably a terrible idea to call this 'me'
var me Host

type Host struct {
	hostname	string			// mirrors
	domainname	*gui.Node		// kernel.org
	hostshort	*gui.Node		// hostname -s
	hostnameStatus	*gui.Node		// is the hostname configured correctly in the OS?
	// fqdn		string			// mirrors.kernel.org

//	dnsTTL		int	`default:"3"`	// Recheck DNS is working every TTL (in seconds)
//	dnsTTLsleep	float64			// sleep between loops
	artificialSleep float64	`default:"0.7"`	// artificial sleep on startup
	artificialS     string 	`default:"abc"`	// artificial sleep on startup

	ttl		*cloudflare.Duration
	dnsTtl		*cloudflare.Duration
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

	window		*gui.Node		// the main window
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
	DnsAAAA		*gui.Node		// the actual DNS AAAA results
	workingIPv6	*gui.Node		// currently working AAAA
	DnsA		*gui.Node		// the actual DNS A results (ignore for status since mostly never happens?)
	DnsStatus	*gui.Node		// the current state of DNS
	DnsSpeed	*gui.Node		// 'FAST', 'OK', 'SLOW', etc
	DnsSpeedActual	*gui.Node		// the last actual duration
	DnsSpeedLast	string			// the last state 'FAST', 'OK', etc

	fix		*gui.Node		// button for the user to click
	fixProc		*gui.Node		// button for the user to click

	mainStatus	*gui.Node		// group for the main display of stuff
	cloudflareB	*gui.Node		// cloudflare button

	dbOn		*gui.Node		// button for setting debugging on
	dbNet		*gui.Node		// button for setting network debugging on
	dbProc		*gui.Node		// button for setting proc debugging on

	digStatus	*digStatus
	statusIPv6	*cloudflare.OneLiner
	digStatusButton *gui.Node
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
