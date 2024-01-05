/* 
	'dig'

	This is essentially doing what the command 'dig' does
	It performing DNS queries on TCP and UDP
	against localhost, cloudflare & google
	
	IPv4() and IPv6() return true if they are working
	
	with the 'gui' package, it can also display the results
*/

package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"reflect"
	"errors"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/shell"

	"github.com/miekg/dns"
)

type digStatus struct {
	ready		bool
	hidden		bool
	statusIPv4	string
	statusIPv6	string

	parent	*gui.Node
	window	*gadgets.BasicWindow
	group	*gui.Node
	grid	*gui.Node

	summary		*gui.Node
	status		*gadgets.OneLiner
	statusAAAA	*gadgets.OneLiner
	speed		*gadgets.OneLiner
	speedActual	*gadgets.OneLiner

	details		*gui.Node
	dsLocalhost	*dnsStatus
	dsLocalNetwork	*dnsStatus
	dsCloudflare	*dnsStatus
	dsGoogle	*dnsStatus
	DnsDigUDP	*gui.Node
	DnsDigTCP	*gui.Node

	httpGoWitCom	*gadgets.OneLiner
	statusHTTP	*gadgets.OneLiner
}

type dnsStatus struct {
	title	string
	server	string // The DNS server. Example: "127.0.0.1:53" or "1.1.1.1:53"
	hostname	string // the hostname to lookup. Example: "www.google.com" or "go.wit.com"

	parent	*gui.Node
	group	*gui.Node
	grid	*gui.Node

	// DNS setup options
	udpA	*gui.Node
	tcpA	*gui.Node
	udpAAAA	*gui.Node
	tcpAAAA	*gui.Node

	// show the display
	aFail		*gui.Node
	aSuccess	*gui.Node
	aaaaFail	*gui.Node
	aaaaSuccess	*gui.Node

	// interger counters
	aFailc		int
	aSuccessc	int
	aaaaFailc	int
	aaaaSuccessc	int
}

func NewDigStatusWindow(p *gui.Node) *digStatus {
	var ds *digStatus
	ds = new(digStatus)

	ds.ready = false
	ds.hidden = true

	ds.window = gadgets.NewBasicWindow(p, "DNS Resolver Status")
	ds.window.Hide()

	// summary of the current state of things
	ds.summary = ds.window.Box().NewGroup("Summary")
	g := ds.summary.NewGrid("LookupStatus", 2, 2)
	g.Pad()

	ds.status	= gadgets.NewOneLiner(g, "status").Set("unknown")
	ds.statusAAAA	= gadgets.NewOneLiner(g, "IPv6 status").Set("unknown")
	ds.statusHTTP	= gadgets.NewOneLiner(g, "IPv6 via HTTP").Set("unknown")
	ds.speed	= gadgets.NewOneLiner(g, "speed").Set("unknown")
	ds.speedActual	= gadgets.NewOneLiner(g, "actual").Set("unknown")

	// make the area to store the raw details
	ds.details = ds.window.Box().NewGroup("Details")
	ds.dsLocalhost		= NewDnsStatus(ds.details, "(localhost)", "127.0.0.1:53", "go.wit.com")
	ds.dsLocalNetwork	= NewDnsStatus(ds.details, "(Local Network)", "172.22.0.1:53", "go.wit.com")
	ds.dsCloudflare		= NewDnsStatus(ds.details, "(cloudflare)", "1.1.1.1:53", "go.wit.com")
	ds.dsGoogle		= NewDnsStatus(ds.details, "(google)", "8.8.8.8:53", "go.wit.com")
	ds.makeDnsStatusGrid()
	ds.makeHttpStatusGrid()

	ds.hidden = false
	ds.ready = true
	return ds
}

func (ds *digStatus) Update() {
	log.Info("digStatus() Update() START")
	if ds == nil {
		log.Error(errors.New("digStatus() Update() ds == nil"))
		return
	}
	duration := timeFunction(func () {
		ds.updateDnsStatus()
	})
	s := fmt.Sprint(duration)
	// ds.speedActual.Set(s)
	me.digStatus.set(ds.speedActual, s)

	if (duration > 500 * time.Millisecond ) {
		me.digStatus.set(ds.speed, "SLOW")
	} else if (duration > 100 * time.Millisecond ) {
		me.digStatus.set(ds.speed, "OK")
	} else {
		me.digStatus.set(ds.speed, "FAST")
	}
	log.Info("digStatus() Update() END")
}

// Returns true if the status is valid
func (ds *digStatus) Ready() bool {
	if ds == nil {return false}
	return ds.ready
}

// Returns true if IPv4 is working
func (ds *digStatus) IPv4() bool {
	if ! ds.Ready() {return false}
	if (ds.statusIPv4 == "OK") {
		return true
	}
	if (ds.statusIPv4 == "GOOD") {
		return true
	}
	return false
}

// Returns true if IPv6 is working
func (ds *digStatus) IPv6() bool {
	if ! ds.Ready() {return false}
	if (ds.statusIPv6 == "GOOD") {
		return true
	}
	return false
}

func (ds *digStatus) setIPv4(s string) {
	ds.statusIPv4 = s
	if ! ds.Ready() {return}
	me.digStatus.set(ds.status, s)
}

func (ds *digStatus) setIPv6(s string) {
	ds.statusIPv6 = s
	if ! ds.Ready() {return}
	me.digStatus.set(ds.statusAAAA, s)
}

func (ds *digStatus) set(a any, s string) {
	if ! ds.Ready() {return}
	if ds.hidden {
		return
	}
	if a == nil {
		return
	}
	var n *gui.Node
	if reflect.TypeOf(a) == reflect.TypeOf(n) {
		n = a.(*gui.Node)
		n.SetText(s)
		return
	}
	var ol *gadgets.OneLiner
	if reflect.TypeOf(a) == reflect.TypeOf(ol) {
		ol = a.(*gadgets.OneLiner)
		ol.Set(s)
		return
	}
	log.Warn("unknown type TypeOf(a) =", reflect.TypeOf(a), "a =", a)
	os.Exit(0)
}

func (ds *digStatus) updateDnsStatus() {
	var cmd, out string
	var ipv4, ipv6 bool

	log.Info("updateDnsStatus() START")
	if (ds == nil) {
		log.Error(errors.New("updateDnsStatus() not initialized yet. ds == nil"))
		return
	}

	if (! ds.ready) {
		log.Error(errors.New("updateDnsStatus() not ready yet"))
		return
	}

	ipv4, ipv6 = ds.dsLocalhost.update()
	ipv4, ipv6 = ds.dsLocalNetwork.update()
	ipv4, ipv6 = ds.dsCloudflare.update()
	ipv4, ipv6 = ds.dsGoogle.update()

	if ds.checkLookupDoH("go.wit.com") {
		log.Println("updateDnsStatus() HTTP DNS lookups working")
		me.digStatus.set(ds.statusHTTP, "WORKING")
	} else {
		log.Println("updateDnsStatus() HTTP DNS lookups not working")
		log.Println("updateDnsStatus() It's really unlikely you are on the internet")
		me.digStatus.set(ds.statusHTTP, "BROKEN")
	}

	if (ipv4) {
		log.Println("updateDnsStatus() IPv4 A lookups working")
		ds.setIPv4("OK")
	} else {
		log.Println("updateDnsStatus() IPv4 A lookups not working. No internet?")
		ds.setIPv4("No Internet?")
	}
	if (ipv6) {
		log.Println("updateDnsStatus() IPv6 AAAA lookups working")
		ds.setIPv4("GOOD")
		ds.setIPv6("GOOD")
	} else {
		log.Println("updateDnsStatus() IPv6 AAAA lookups are not working")
		ds.setIPv6("Need VPN")
	}

	cmd = "dig +noall +answer www.wit.com A"
	out = shell.Run(cmd)
	log.Println("makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigUDP, out)

	cmd = "dig +noall +answer www.wit.com AAAA"
	out = shell.Run(cmd)
	log.Println("makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigTCP, out)
}

// Makes a DNS Status Grid
func NewDnsStatus(p *gui.Node, title string, server string, hostname string) *dnsStatus {
	var ds *dnsStatus
	ds = new(dnsStatus)
	ds.parent = p
	ds.group = p.NewGroup(server + " " + title + " lookup")
	ds.grid = ds.group.NewGrid("LookupStatus", 5, 2)

	ds.server = server
	ds.hostname = hostname

	ds.grid.NewLabel("")
	ds.grid.NewLabel("UDP")
	ds.grid.NewLabel("TCP")
	ds.grid.NewLabel("Success")
	ds.grid.NewLabel("Fail")

	ds.grid.NewLabel("A")
	ds.udpA = ds.grid.NewLabel("?")
	ds.tcpA = ds.grid.NewLabel("?")
	ds.aSuccess = ds.grid.NewLabel("?")
	ds.aFail = ds.grid.NewLabel("?")

	ds.grid.NewLabel("AAAA")
	ds.udpAAAA = ds.grid.NewLabel("?")
	ds.tcpAAAA = ds.grid.NewLabel("?")
	ds.aaaaSuccess = ds.grid.NewLabel("?")
	ds.aaaaFail = ds.grid.NewLabel("?")

	ds.group.Margin()
	ds.grid.Margin()
	ds.group.Pad()
	ds.grid.Pad()

	return ds
}

// special thanks to the Element Hotel wifi in Philidelphia that allowed me to
// easily debug this code since the internet connection here blocks port 53 traffic
func (ds *dnsStatus) update() (bool, bool) {
	var results []string
	var a bool = false
	var aaaa bool = false

	log.Println("dnsStatus.update() For server", ds.server, "on", ds.hostname)
	results, _ = dnsUdpLookup(ds.server, ds.hostname, dns.TypeA)
	log.Println("dnsStatus.update() UDP type A =", results)

	if (len(results) == 0) {
		me.digStatus.set(ds.udpA, "BROKEN")
		ds.aFailc += 1
	} else {
		me.digStatus.set(ds.udpA, "WORKING")
		ds.aSuccessc += 1
		a = true
	}

	results, _ = dnsTcpLookup(ds.server, ds.hostname, dns.TypeA)
	log.Println("dnsStatus.update() TCP type A =", results)

	if (len(results) == 0) {
		me.digStatus.set(ds.tcpA, "BROKEN")
		ds.aFailc += 1
	} else {
		me.digStatus.set(ds.tcpA, "WORKING")
		ds.aSuccessc += 1
		a = true
	}

	me.digStatus.set(ds.aFail, strconv.Itoa(ds.aFailc))
	me.digStatus.set(ds.aSuccess,strconv.Itoa(ds.aSuccessc))

	results, _ = dnsUdpLookup(ds.server, ds.hostname, dns.TypeAAAA)
	log.Println("dnsStatus.update() UDP type AAAA =", results)

	if (len(results) == 0) {
		me.digStatus.set(ds.udpAAAA, "BROKEN")
		ds.aaaaFailc += 1
		me.digStatus.set(ds.aaaaFail, strconv.Itoa(ds.aaaaFailc))
	} else {
		me.digStatus.set(ds.udpAAAA, "WORKING")
		ds.aaaaSuccessc += 1
		aaaa = true
	}

	results, _ = dnsTcpLookup(ds.server, ds.hostname, dns.TypeAAAA)
	log.Println("dnsStatus.update() UDP type AAAA =", results)

	if (len(results) == 0) {
		me.digStatus.set(ds.tcpAAAA, "BROKEN")
		ds.aaaaFailc += 1
		me.digStatus.set(ds.aaaaFail, strconv.Itoa(ds.aaaaFailc))
	} else {
		me.digStatus.set(ds.tcpAAAA, "WORKING")
		ds.aaaaSuccessc += 1
		aaaa = true
	}

	me.digStatus.set(ds.aaaaFail, strconv.Itoa(ds.aaaaFailc))
	me.digStatus.set(ds.aaaaSuccess,strconv.Itoa(ds.aaaaSuccessc))

	return a, aaaa
}

func (ds *digStatus) makeHttpStatusGrid() {
	group := ds.details.NewGroup("dns.google.com via HTTPS")
	grid := group.NewGrid("LookupStatus", 2, 2)

	ds.httpGoWitCom = gadgets.NewOneLiner(grid, "go.wit.com")
	me.digStatus.set(ds.httpGoWitCom, "unknown")

	group.Pad()
	grid.Pad()
}

func (ds *digStatus) makeDnsStatusGrid() {
	var cmd, out string
	group := ds.details.NewGroup("dig results")
	grid := group.NewGrid("LookupStatus", 2, 2)

	cmd = "dig +noall +answer go.wit.com A"
	grid.NewLabel(cmd)
	ds.DnsDigUDP = grid.NewLabel("?")
	out = shell.Run(cmd)
	log.Println("makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigUDP, out)

	cmd = "dig +noall +answer go.wit.com AAAA"
	grid.NewLabel(cmd)
	ds.DnsDigTCP = grid.NewLabel("?")
	out = shell.Run(cmd)
	log.Println("makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigTCP, out)

	group.Pad()
	grid.Pad()
}

// dnsLookup performs a DNS lookup for the specified record type (e.g., "TXT", "AAAA") for a given domain.
func dnsUdpLookup(server string, domain string, recordType uint16) ([]string, error) {
	var records []string

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), recordType)
	r, _, err := c.Exchange(m, server) // If server = "1.1.1.1:53" then use Cloudflare's DNS server
	if err != nil {
		return nil, err
	}

	for _, ans := range r.Answer {
		records = append(records, ans.String())
	}

	return records, nil
}

func dnsTcpLookup(server string, domain string, recordType uint16) ([]string, error) {
	var records []string

	c := new(dns.Client)
	c.Net = "tcp" // Specify to use TCP for the query
	c.Timeout = time.Second * 5  // Set a 5-second timeout
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), recordType)
	r, _, err := c.Exchange(m, server) // If server = "1.1.1.1:53" then use Cloudflare's DNS server 
	if err != nil {
		return nil, err
	}

	for _, ans := range r.Answer {
		records = append(records, ans.String())
	}

	return records, nil
}

func (ds *digStatus) checkLookupDoH(hostname string) bool {
	var status bool = false

	domain := "go.wit.com"
	ipv6Addresses, err := dnsAAAAlookupDoH(domain)
	if err != nil {
		log.Error(err, "checkLookupDoH()")
		return status
	}

	log.Println("IPv6 Addresses for %s:\n", domain)
	for _, addr := range ipv6Addresses {
		log.Println(addr)
		me.digStatus.set(ds.httpGoWitCom, addr)
		status = true
	}
	return status
}

func (ds *digStatus) Show() {
	log.Info("digStatus.Show() window")
	if me.digStatus.hidden {
		me.digStatus.window.Show()
	}
	me.digStatus.hidden = false
}

func (ds *digStatus) Hide() {
	log.Info("digStatus.Hide() window")
	if ! me.digStatus.hidden {
		me.digStatus.window.Hide()
	}
	me.digStatus.hidden = true
}
