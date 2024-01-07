// This creates a simple hello world window
package main

import 	(
	"net"
	"time"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	// "go.wit.com/gui/cloudflare"
	"go.wit.com/apps/control-panel-dns/linuxstatus"
	"go.wit.com/apps/control-panel-dns/smartwindow"

	"github.com/miekg/dns"
)

// It's probably a terrible idea to call this 'me'
var me Host

type Host struct {
	myGui		*gui.Node	// the 'gui' binary tree root node

	window		*gadgets.BasicWindow	// the main window
	debug		*gadgets.BasicWindow	// the debug window

	statusDNS	*hostnameStatus		// keeps track of the hostname and it's status
	statusOS	*linuxstatus.LinuxStatus		// what the Linux OS sees
	digStatus	*digStatus		// window of the results of DNS lookups

	// WHEN THESE ARE ALL "WORKING", then everything is good
	hostnameStatus	*gui.Node		// a summary for the user of where things are
	DnsAPIstatus	*gui.Node		// does your DNS API work?
	APIprovider	string
	apiButton	*gui.Node		// the button you click for the API config page

	artificialSleep float64	`default:"0.7"`	// artificial sleep on startup
	artificialS     string 	`default:"abc"`	// artificial sleep on startup

	ttl		*gadgets.Duration
	dnsTtl		*gadgets.Duration
	dnsSleep	time.Duration
	localSleep	time.Duration

	changed		bool			// set to true if things changed

	ipmap		map[string]*IPtype	// the current ip addresses
	dnsmap		map[string]*IPtype	// the current dns addresses
	ifmap		map[int]*IFtype		// the current interfaces
	nsmap		map[string]string	// the NS records

	// DNS A and AAAA results
	ipv4s		map[string]dns.RR
	ipv6s		map[string]dns.RR

	// DNS stuff
	DnsStatus	*gui.Node		// the current state of DNS
	DnsSpeed	*gui.Node		// 'FAST', 'OK', 'SLOW', etc
	DnsSpeedActual	*gui.Node		// the last actual duration
	DnsSpeedLast	string			// the last state 'FAST', 'OK', etc

	statusIPv6	*gadgets.OneLiner
	digStatusButton *gui.Node
	statusDNSbutton *gui.Node
	witcom		*gadgets.BasicWindow
	fixButton	*gui.Node
	fixWindow	*smartwindow.SmartWindow
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
