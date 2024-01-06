/* 
	Performs DNS queries on TCP and UDP
*/

package main

import (
	"os"
	"time"
	"strconv"
	"reflect"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"

	"github.com/miekg/dns"
)

type resolverStatus struct {
	title		string
	server		string // The DNS server. Example: "127.0.0.1:53" or "1.1.1.1:53"
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

func (rs *resolverStatus) set(a any, s string) {
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

// Makes a DNS Status Grid
func NewResolverStatus(p *gui.Node, title string, server string, hostname string) *resolverStatus {
	var rs *resolverStatus
	rs = new(resolverStatus)
	rs.parent = p
	rs.group = p.NewGroup(server + " " + title + " lookup")
	rs.grid = rs.group.NewGrid("LookupStatus", 5, 2)

	rs.server = server
	rs.hostname = hostname

	rs.grid.NewLabel("")
	rs.grid.NewLabel("UDP")
	rs.grid.NewLabel("TCP")
	rs.grid.NewLabel("Success")
	rs.grid.NewLabel("Fail")

	rs.grid.NewLabel("A")
	rs.udpA = rs.grid.NewLabel("?")
	rs.tcpA = rs.grid.NewLabel("?")
	rs.aSuccess = rs.grid.NewLabel("?")
	rs.aFail = rs.grid.NewLabel("?")

	rs.grid.NewLabel("AAAA")
	rs.udpAAAA = rs.grid.NewLabel("?")
	rs.tcpAAAA = rs.grid.NewLabel("?")
	rs.aaaaSuccess = rs.grid.NewLabel("?")
	rs.aaaaFail = rs.grid.NewLabel("?")

	rs.group.Margin()
	rs.grid.Margin()
	rs.group.Pad()
	rs.grid.Pad()

	return rs
}

// special thanks to the Element Hotel wifi in Philidelphia that allowed me to
// easily debug this code since the internet connection here blocks port 53 traffic
func (rs *resolverStatus) update() (bool, bool) {
	var results []string
	var a bool = false
	var aaaa bool = false

	log.Log(DNS, "resolverStatus.update() For server", rs.server, "on", rs.hostname)
	results, _ = dnsUdpLookup(rs.server, rs.hostname, dns.TypeA)
	log.Log(DNS, "resolverStatus.update() UDP type A =", results)

	if (len(results) == 0) {
		rs.set(rs.udpA, "BROKEN")
		rs.aFailc += 1
	} else {
		rs.set(rs.udpA, "WORKING")
		rs.aSuccessc += 1
		a = true
	}

	results, _ = dnsTcpLookup(rs.server, rs.hostname, dns.TypeA)
	log.Log(DNS, "resolverStatus.update() TCP type A =", results)

	if (len(results) == 0) {
		rs.set(rs.tcpA, "BROKEN")
		rs.aFailc += 1
	} else {
		me.digStatus.set(rs.tcpA, "WORKING")
		rs.aSuccessc += 1
		a = true
	}

	me.digStatus.set(rs.aFail, strconv.Itoa(rs.aFailc))
	me.digStatus.set(rs.aSuccess,strconv.Itoa(rs.aSuccessc))

	results, _ = dnsUdpLookup(rs.server, rs.hostname, dns.TypeAAAA)
	log.Log(DNS, "resolverStatus.update() UDP type AAAA =", results)

	if (len(results) == 0) {
		me.digStatus.set(rs.udpAAAA, "BROKEN")
		rs.aaaaFailc += 1
		me.digStatus.set(rs.aaaaFail, strconv.Itoa(rs.aaaaFailc))
	} else {
		me.digStatus.set(rs.udpAAAA, "WORKING")
		rs.aaaaSuccessc += 1
		aaaa = true
	}

	results, _ = dnsTcpLookup(rs.server, rs.hostname, dns.TypeAAAA)
	log.Log(DNS, "resolverStatus.update() UDP type AAAA =", results)

	if (len(results) == 0) {
		me.digStatus.set(rs.tcpAAAA, "BROKEN")
		rs.aaaaFailc += 1
		me.digStatus.set(rs.aaaaFail, strconv.Itoa(rs.aaaaFailc))
	} else {
		me.digStatus.set(rs.tcpAAAA, "WORKING")
		rs.aaaaSuccessc += 1
		aaaa = true
	}

	me.digStatus.set(rs.aaaaFail, strconv.Itoa(rs.aaaaFailc))
	me.digStatus.set(rs.aaaaSuccess,strconv.Itoa(rs.aaaaSuccessc))

	return a, aaaa
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
