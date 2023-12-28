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
	"log"
	"fmt"
	"time"
	"strconv"

	"github.com/miekg/dns"
	"go.wit.com/gui"
	"go.wit.com/control-panel-dns/cloudflare"
	"go.wit.com/shell"
)

type digStatus struct {
	ready		bool
	statusIPv4	string
	statusIPv6	string

	parent	*gui.Node
	window	*gui.Node
	group	*gui.Node
	grid	*gui.Node
	box	*gui.Node

	summary		*gui.Node
	status		*cloudflare.OneLiner
	statusAAAA	*cloudflare.OneLiner
	speed		*cloudflare.OneLiner
	speedActual	*cloudflare.OneLiner

	details		*gui.Node
	dsLocalhost	*dnsStatus
	dsLocalNetwork	*dnsStatus
	dsCloudflare	*dnsStatus
	dsGoogle	*dnsStatus
	DnsDigUDP	*gui.Node
	DnsDigTCP	*gui.Node
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

	ds.window = p.NewWindow("DNS Lookup Status")
	ds.box = ds.window.NewBox("hBox", true)

	// summary of the current state of things
	ds.summary = ds.box.NewGroup("Summary")

	b := ds.summary.NewBox("hBox", true)
	ds.status = cloudflare.NewOneLiner(b, "status")
	ds.status.Set("unknown")

	b = ds.summary.NewBox("hBox", true)
	ds.statusAAAA = cloudflare.NewOneLiner(b, "IPv6 status")
	ds.statusAAAA.Set("unknown")

	b = ds.summary.NewBox("hBox", true)
	ds.speed = cloudflare.NewOneLiner(b, "speed")
	ds.speed.Set("unknown")

	b = ds.summary.NewBox("hBox", true)
	ds.speedActual = cloudflare.NewOneLiner(b, "actual")
	ds.speedActual.Set("unknown")

	// make the area to store the raw details
	ds.details = ds.box.NewGroup("Details")
	ds.dsLocalhost = NewDnsStatus(ds.details, "(localhost)", "127.0.0.1:53", "go.wit.com")
	ds.dsLocalNetwork = NewDnsStatus(ds.details, "(Local Network)", "172.22.0.1:53", "go.wit.com")
	ds.dsCloudflare = NewDnsStatus(ds.details, "(cloudflare)", "1.1.1.1:53", "go.wit.com")
	ds.dsGoogle = NewDnsStatus(ds.details, "(google)", "8.8.8.8:53", "go.wit.com")
	ds.makeDnsStatusGrid()

	return ds
}

func (ds *digStatus) Update() {
	duration := timeFunction(func () { ds.updateDnsStatus() })
	s := fmt.Sprint(duration)
	ds.speedActual.Set(s)

	if (duration > 500 * time.Millisecond ) {
		ds.speed.Set("SLOW")
	} else if (duration > 100 * time.Millisecond ) {
		ds.speed.Set("OK")
	} else {
		ds.speed.Set("FAST")
	}
}

// Returns true if the status is valid
func (ds *digStatus) Ready() bool {
	return ds.ready
}

// Returns true if IPv4 is working
func (ds *digStatus) IPv4() bool {
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
	if (ds.statusIPv6 == "GOOD") {
		return true
	}
	return false
}

func (ds *digStatus) setIPv4(s string) {
	ds.status.Set(s)
	ds.statusIPv4 = s
}

func (ds *digStatus) setIPv6(s string) {
	ds.statusAAAA.Set(s)
	ds.statusIPv6 = s
}

func (ds *digStatus) updateDnsStatus() {
	var cmd, out string
	var ipv4, ipv6 bool

	ipv4, ipv6 = ds.dsLocalhost.Update()
	ipv4, ipv6 = ds.dsLocalNetwork.Update()
	ipv4, ipv6 = ds.dsCloudflare.Update()
	ipv4, ipv6 = ds.dsGoogle.Update()

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
	ds.DnsDigUDP.SetText(out)

	cmd = "dig +noall +answer www.wit.com AAAA"
	out = shell.Run(cmd)
	log.Println("makeDnsStatusGrid() dig", out)
	ds.DnsDigTCP.SetText(out)

	ds.ready = true
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
func (ds *dnsStatus) Update() (bool, bool) {
	var results []string
	var a bool = false
	var aaaa bool = false

	log.Println("dnsStatus.Update() For server", ds.server, "on", ds.hostname)
	results, _ = dnsUdpLookup(ds.server, ds.hostname, dns.TypeA)
	log.Println("dnsStatus.Update() UDP type A =", results)

	if (len(results) == 0) {
		ds.udpA.SetText("BROKEN")
		ds.aFailc += 1
	} else {
		ds.udpA.SetText("WORKING")
		ds.aSuccessc += 1
		a = true
	}

	results, _ = dnsTcpLookup(ds.server, ds.hostname, dns.TypeA)
	log.Println("dnsStatus.Update() TCP type A =", results)

	if (len(results) == 0) {
		ds.tcpA.SetText("BROKEN")
		ds.aFailc += 1
	} else {
		ds.tcpA.SetText("WORKING")
		ds.aSuccessc += 1
		a = true
	}

	ds.aFail.SetText(strconv.Itoa(ds.aFailc))
	ds.aSuccess.SetText(strconv.Itoa(ds.aSuccessc))

	results, _ = dnsUdpLookup(ds.server, ds.hostname, dns.TypeAAAA)
	log.Println("dnsStatus.Update() UDP type AAAA =", results)

	if (len(results) == 0) {
		ds.udpAAAA.SetText("BROKEN")
		ds.aaaaFailc += 1
		ds.aaaaFail.SetText(strconv.Itoa(ds.aaaaFailc))
	} else {
		ds.udpAAAA.SetText("WORKING")
		ds.aaaaSuccessc += 1
		aaaa = true
	}

	results, _ = dnsTcpLookup(ds.server, ds.hostname, dns.TypeAAAA)
	log.Println("dnsStatus.Update() UDP type AAAA =", results)

	if (len(results) == 0) {
		ds.tcpAAAA.SetText("BROKEN")
		ds.aaaaFailc += 1
		ds.aaaaFail.SetText(strconv.Itoa(ds.aaaaFailc))
	} else {
		ds.tcpAAAA.SetText("WORKING")
		ds.aaaaSuccessc += 1
		aaaa = true
	}

	ds.aaaaFail.SetText(strconv.Itoa(ds.aaaaFailc))
	ds.aaaaSuccess.SetText(strconv.Itoa(ds.aaaaSuccessc))

	return a, aaaa
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
	ds.DnsDigUDP.SetText(out)

	cmd = "dig +noall +answer go.wit.com AAAA"
	grid.NewLabel(cmd)
	ds.DnsDigTCP = grid.NewLabel("?")
	out = shell.Run(cmd)
	log.Println("makeDnsStatusGrid() dig", out)
	ds.DnsDigTCP.SetText(out)

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
