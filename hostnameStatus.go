/* 
	figures out if your hostname is valid
	then checks if your DNS is setup correctly
*/

package main

import (
	"os"
	"fmt"
	"time"
	"reflect"
	"strings"

	"go.wit.com/log"
	"go.wit.com/gui"
	"go.wit.com/control-panel-dns/cloudflare"
)

type hostnameStatus struct {
	ready		bool
	hidden		bool

	hostname	string // my hostname. Example: "test.wit.com"

	window	*gui.Node

	status		*cloudflare.OneLiner
	summary		*cloudflare.OneLiner
	speed		*cloudflare.OneLiner
	speedActual	*cloudflare.OneLiner

	dnsA		*cloudflare.OneLiner
	dnsAAAA		*cloudflare.OneLiner
	dnsAPI		*cloudflare.OneLiner

	statusIPv4	*cloudflare.OneLiner
	statusIPv6	*cloudflare.OneLiner

	// Details Group
	currentIPv4	*cloudflare.OneLiner
	currentIPv6	*cloudflare.OneLiner
}

func NewHostnameStatusWindow(p *gui.Node) *hostnameStatus {
	var hs *hostnameStatus
	hs = new(hostnameStatus)

	hs.ready = false
	hs.hidden = true
	hs.hostname = me.hostname

	hs.window = p.NewWindow( hs.hostname + " Status")
	hs.window.Custom = func () {
		hs.hidden = true
		hs.window.Hide()
	}
	box := hs.window.NewBox("hBox", true)
	group := box.NewGroup("Summary")
	grid := group.NewGrid("LookupStatus", 2, 2)

	hs.status	= cloudflare.NewOneLiner(grid, "status").Set("unknown")
	hs.statusIPv4	= cloudflare.NewOneLiner(grid, "IPv4").Set("unknown")
	hs.statusIPv6	= cloudflare.NewOneLiner(grid, "IPv6").Set("unknown")

	group.Pad()
	grid.Pad()

	group = box.NewGroup("Details")
	grid = group.NewGrid("LookupDetails", 2, 2)

	hs.currentIPv4	= cloudflare.NewOneLiner(grid, "Current IPv4")
	hs.currentIPv6	= cloudflare.NewOneLiner(grid, "Current IPv6")

	hs.dnsAPI	= cloudflare.NewOneLiner(grid, "dns API provider").Set("unknown")
	hs.dnsA		= cloudflare.NewOneLiner(grid, "dns IPv4 resource records").Set("unknown")
	hs.dnsAAAA	= cloudflare.NewOneLiner(grid, "dns IPv6 resource records").Set("unknown")
	hs.speed	= cloudflare.NewOneLiner(grid, "speed").Set("unknown")
	hs.speedActual	= cloudflare.NewOneLiner(grid, "actual").Set("unknown")

	group.Pad()
	grid.Pad()

	hs.hidden = false
	hs.ready = true
	return hs
}

func (hs *hostnameStatus) Update() {
	log.Info("hostnameStatus() Update() START")
	if hs == nil {
		log.Error("hostnameStatus() Update() hs == nil")
		return
	}
	duration := timeFunction(func () {
		hs.updateStatus()
	})
	s := fmt.Sprint(duration)
	hs.set(hs.speedActual, s)

	if (duration > 500 * time.Millisecond ) {
		hs.set(hs.speed, "SLOW")
	} else if (duration > 100 * time.Millisecond ) {
		hs.set(hs.speed, "OK")
	} else {
		hs.set(hs.speed, "FAST")
	}
	log.Info("hostnameStatus() Update() END")
}

// Returns true if the status is valid
func (hs *hostnameStatus) Ready() bool {
	if hs == nil {return false}
	return hs.ready
}

// Returns true if IPv4 is working
func (hs *hostnameStatus) IPv4() bool {
	if ! hs.Ready() {return false}
	if (hs.statusIPv4.Get() == "OK") {
		return true
	}
	if (hs.statusIPv4.Get() == "GOOD") {
		return true
	}
	return false
}

// Returns true if IPv6 is working
func (hs *hostnameStatus) IPv6() bool {
	if ! hs.Ready() {return false}
	if (hs.statusIPv6.Get() == "GOOD") {
		return true
	}
	return false
}

func (hs *hostnameStatus) setIPv4(s string) {
	hs.statusIPv4.Set(s)
	if ! hs.Ready() {return}
}

func (hs *hostnameStatus) setIPv6(s string) {
	hs.statusIPv6.Set(s)
	if ! hs.Ready() {return}
}

func (hs *hostnameStatus) set(a any, s string) {
	if ! hs.Ready() {return}
	if hs.hidden {
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
	var ol *cloudflare.OneLiner
	if reflect.TypeOf(a) == reflect.TypeOf(ol) {
		ol = a.(*cloudflare.OneLiner)
		if ol == nil {
			log.Println("ol = nil", reflect.TypeOf(a), "a =", a)
			return
		}
		log.Println("SETTING ol:", ol)
		ol.Set(s)
		return
	}
	log.Error("unknown type TypeOf(a) =", reflect.TypeOf(a), "a =", a)
	os.Exit(0)
}

func (hs *hostnameStatus) updateStatus() {
	var s string
	var vals []string
	log.Info("updateStatus() START")
	if ! hs.Ready() { return }

	vals = lookupDoH(hs.hostname, "AAAA")

	log.Println("IPv6 Addresses for ", hs.hostname, "=", vals)
	if len(vals) == 0 {
		s = "(none)"
		hs.setIPv6("NEED VPN")
	} else {
		for _, addr := range vals {
			log.Println(addr)
			s += addr + " (DELETE)"
			hs.setIPv6("NEEDS DELETE")
		}
	}
	hs.set(hs.dnsAAAA, s)

	vals = lookupDoH(hs.hostname, "A")
	log.Println("IPv4 Addresses for ", hs.hostname, "=", vals)
	s = strings.Join(vals, "\n")
	if (s == "") {
		s = "(none)"
		hs.setIPv4("NEEDS CNAME")
	}
	hs.set(hs.dnsA, s)

	vals = lookupDoH(hs.hostname, "CNAME")
	s = strings.Join(vals, "\n")
	if (s != "") {
		hs.set(hs.dnsA, "CNAME " + s)
		hs.setIPv4("GOOD")
	}

	hs.currentIPv4.Set(me.IPv4.S)
	hs.currentIPv6.Set(me.IPv6.S)

	if hs.IPv4() && hs.IPv4() {
		hs.status.Set("GOOD")
	} else {
		hs.status.Set("BROKEN")
	}

	hs.dnsAPI.Set(me.DnsAPI.S)
}

func (hs *hostnameStatus) Show() {
	log.Info("hostnameStatus.Show() window")
	if hs.hidden {
		hs.window.Show()
	}
	hs.hidden = false
}

func (hs *hostnameStatus) Hide() {
	log.Info("hostnameStatus.Hide() window")
	if ! hs.hidden {
		hs.window.Hide()
	}
	hs.hidden = true
}
