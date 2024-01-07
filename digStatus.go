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
	"strings"
	"reflect"
	"errors"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/shell"
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
	dsLocalhost	*resolverStatus
	dsLocalNetwork	*resolverStatus
	dsCloudflare	*resolverStatus
	dsGoogle	*resolverStatus
	DnsDigUDP	*gui.Node
	DnsDigTCP	*gui.Node

	httpGoWitCom	*gadgets.OneLiner
	statusHTTP	*gadgets.OneLiner
}

func NewDigStatusWindow(p *gui.Node) *digStatus {
	var ds *digStatus
	ds = new(digStatus)

	ds.ready = false
	ds.hidden = true

	ds.window = gadgets.NewBasicWindow(p, "DNS Resolver Status")
	ds.window.Draw()
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
	ds.dsLocalhost		= NewResolverStatus(ds.details, "(localhost)", "127.0.0.1:53", "go.wit.com")
	ds.dsLocalNetwork	= NewResolverStatus(ds.details, "(Local Network)", "192.168.86.1:53", "go.wit.com")
	ds.dsCloudflare		= NewResolverStatus(ds.details, "(cloudflare)", "1.1.1.1:53", "go.wit.com")
	ds.dsGoogle		= NewResolverStatus(ds.details, "(google)", "8.8.8.8:53", "go.wit.com")
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

func (ds *digStatus) setIPv4status(s string) {
	ds.statusIPv4 = s
	if ! ds.Ready() {return}
	me.digStatus.set(ds.status, s)
}

func (ds *digStatus) setIPv6status(s string) {
	ds.statusIPv6 = s
	if ! ds.Ready() {return}
	me.digStatus.set(ds.statusAAAA, s)
}

func (ds *digStatus) SetIPv6(s string) {
	if ! ds.Ready() {return}
	log.Warn("Should SetIPv6() here to", s)
	log.Warn("Should SetIPv6() here to", s)
	log.Warn("Should SetIPv6() here to", s)
	log.Warn("Should SetIPv6() here to", s)
	me.DnsAAAA.Set(s)
	// me.digStatus.set(ds.httpGoWitCom, addr)
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

	if me.statusOS.ValidHostname() {
		if ds.checkLookupDoH(me.statusOS.GetHostname()) {
			log.Log(DNS, "updateDnsStatus() HTTP DNS lookups working")
			me.digStatus.set(ds.statusHTTP, "WORKING")
		} else {
			log.Log(DNS, "updateDnsStatus() HTTP DNS lookups not working")
			log.Log(DNS, "updateDnsStatus() It's really unlikely you are on the internet")
			me.digStatus.set(ds.statusHTTP, "BROKEN")
		}
	} else {
		me.digStatus.set(ds.statusHTTP, "INVALID HOSTNAME")
	}

	if (ipv4) {
		log.Log(DNS, "updateDnsStatus() IPv4 A lookups working")
		ds.setIPv4status("OK")
	} else {
		log.Log(DNS, "updateDnsStatus() IPv4 A lookups not working. No internet?")
		ds.setIPv4status("No Internet?")
	}
	if (ipv6) {
		log.Log(DNS, "updateDnsStatus() IPv6 AAAA lookups working")
		ds.setIPv4status("GOOD")
		ds.setIPv6status("GOOD")
	} else {
		log.Log(DNS, "updateDnsStatus() IPv6 AAAA lookups are not working")
		ds.setIPv6status("Need VPN")
	}

	cmd = "dig +noall +answer www.wit.com A"
	out = shell.Run(cmd)
	log.Log(DNS, "makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigUDP, out)

	cmd = "dig +noall +answer www.wit.com AAAA"
	out = shell.Run(cmd)
	log.Log(DNS, "makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigTCP, out)

	/*
	g2.NewButton("dig +trace", func () {
		log.Log(NOW, "TODO: redo this")
		// o := shell.Run("dig +trace +noadditional DS " + me.hostname + " @8.8.8.8")
		// log.Println(o)
	})
	*/
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
	log.Log(DNS, "makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigUDP, out)

	cmd = "dig +noall +answer go.wit.com AAAA"
	grid.NewLabel(cmd)
	ds.DnsDigTCP = grid.NewLabel("?")
	out = shell.Run(cmd)
	log.Log(DNS, "makeDnsStatusGrid() dig", out)
	me.digStatus.set(ds.DnsDigTCP, out)

	group.Pad()
	grid.Pad()
}

func (ds *digStatus) checkLookupDoH(hostname string) bool {
	var status bool = false

	ipv6Addresses := lookupDoH(hostname, "AAAA")

	log.Log(DNS, "IPv6 Addresses for ", hostname)
	var s []string
	for _, addr := range ipv6Addresses {
		log.Log(DNS, addr)
		s = append(s, addr)
		status = true
	}
	me.digStatus.SetIPv6(strings.Join(s, "\n"))
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
