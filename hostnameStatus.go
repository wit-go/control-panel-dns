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
	"sort"
	"errors"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

type hostnameStatus struct {
	ready		bool
	hidden		bool

	lastname	string // used to watch for changes in the hostname

	window		*gadgets.BasicWindow

	// Primary Directives
	status		*gadgets.OneLiner
	summary		*gadgets.OneLiner
	statusIPv4	*gadgets.OneLiner
	statusIPv6	*gadgets.OneLiner

	hostname	*gadgets.OneLiner
	domainname	*gadgets.OneLiner

	// what the current IP addresses your network has given you
	currentIPv4	*gadgets.OneLiner
	currentIPv6	*gadgets.OneLiner

	// what the DNS servers have
	NSrr		*gadgets.OneLiner
	dnsA		*gadgets.OneLiner
	dnsAAAA		*gadgets.OneLiner
	dnsAPI		*gadgets.OneLiner

	speed		*gadgets.OneLiner
	speedActual	*gadgets.OneLiner

	// Actions
//	dnsValue	*gui.Node
//	dnsAction	*gui.Node
}

func NewHostnameStatusWindow(p *gui.Node) *hostnameStatus {
	var hs *hostnameStatus
	hs = new(hostnameStatus)

	hs.ready = false
	hs.hidden = true
	// hs.hostname = me.hostname

	hs.window = gadgets.NewBasicWindow(p, "fix hostname here" + " Status")
	hs.window.Draw()
	hs.window.Hide()

	group := hs.window.Box().NewGroup("Summary")
	grid := group.NewGrid("LookupStatus", 2, 2)

	hs.status	= gadgets.NewOneLiner(grid, "status").Set("unknown")
	hs.statusIPv4	= gadgets.NewOneLiner(grid, "IPv4").Set("unknown")
	hs.statusIPv6	= gadgets.NewOneLiner(grid, "IPv6").Set("unknown")

	group.Pad()
	grid.Pad()

	group = hs.window.Box().NewGroup("Details")
	grid = group.NewGrid("LookupDetails", 2, 2)

	hs.hostname	= gadgets.NewOneLiner(grid, "hostname")
	hs.domainname	= gadgets.NewOneLiner(grid, "domain name")
	hs.currentIPv4	= gadgets.NewOneLiner(grid, "Current IPv4")
	hs.currentIPv6	= gadgets.NewOneLiner(grid, "Current IPv6")

	hs.NSrr		= gadgets.NewOneLiner(grid, "dns NS records").Set("unknown")
	hs.dnsAPI	= gadgets.NewOneLiner(grid, "dns API provider").Set("unknown")
	hs.dnsA		= gadgets.NewOneLiner(grid, "dns IPv4 resource records").Set("unknown")
	hs.dnsAAAA	= gadgets.NewOneLiner(grid, "dns IPv6 resource records").Set("unknown")

	hs.speed	= gadgets.NewOneLiner(grid, "speed").Set("unknown")
	hs.speedActual	= gadgets.NewOneLiner(grid, "actual").Set("unknown")

	group.Pad()
	grid.Pad()

	/*
	group = hs.window.Box().NewGroup("Actions")
	grid = group.NewGrid("LookupDetails", 2, 2)

	hs.dnsValue	= grid.NewLabel("3.4.5.6")
	hs.dnsAction	= grid.NewButton("CHECK",  func () {
		log.Warn("should", hs.dnsAction.S, "here for", hs.dnsValue.S)
		if (hs.dnsAction.S == "DELETE") {
			hs.deleteDNSrecord(hs.dnsValue.S)
		}
		if (hs.dnsAction.S == "CREATE") {
			hs.createDNSrecord(hs.dnsValue.S)
		}
	})
	*/

	group.Pad()
	grid.Pad()

	hs.hidden = false
	hs.ready = true
	return hs
}

func (hs *hostnameStatus) Domain() string {
	if ! hs.Ready() {return ""}
	return hs.domainname.Get()
}

func (hs *hostnameStatus) API() string {
	if ! hs.Ready() {return ""}
	return hs.dnsAPI.Get()
}

func (hs *hostnameStatus) Update() {
	log.Info("hostnameStatus() Update() START")
	if hs == nil {
		log.Error(errors.New("hostnameStatus() Update() hs == nil"))
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
	if ! hs.Ready() {return}
	hs.statusIPv4.Set(s)
}

func (hs *hostnameStatus) setIPv6(s string) {
	if ! hs.Ready() {return}
	hs.statusIPv6.Set(s)
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
	var ol *gadgets.OneLiner
	if reflect.TypeOf(a) == reflect.TypeOf(ol) {
		ol = a.(*gadgets.OneLiner)
		if ol == nil {
			// log.Println("ol = nil", reflect.TypeOf(a), "a =", a)
			return
		}
		// log.Println("SETTING ol:", ol)
		ol.Set(s)
		return
	}
	log.Warn("unknown type TypeOf(a) =", reflect.TypeOf(a), "a =", a)
	os.Exit(0)
}

// returns true if AAAA record already exists in DNS
func (hs *hostnameStatus) existsAAAA(s string) bool {
	log.Log(NOW, "existsAAAA() try to see if AAAA is already set", s)
	return false
}

/*
// figure out if I'm missing any IPv6 address in DNS
func (hs *hostnameStatus) missingAAAA() bool {
	var aaaa []string
	aaaa = dhcpAAAA()
	for _, s := range aaaa {
		log.Log(NET, "my actual AAAA = ",s)
		if hs.existsAAAA(s) {
			log.Log(NOW, "my actual AAAA already exists in DNS =",s)
		} else {
			log.Log(NOW, "my actual AAAA is missing from DNS",s)
			hs.dnsValue.SetText(s)
			hs.dnsAction.SetText("CREATE")
			return true
		}
	}

	return false
}
*/

func (hs *hostnameStatus) GetIPv6() []string {
	if ! hs.Ready() { return nil}
	return strings.Split(hs.dnsAAAA.Get(), "\n")
}

func (hs *hostnameStatus) updateStatus() {
	if ! hs.Ready() { return }
	var s string
	var vals []string
	log.Log(STATUS, "updateStatus() START")

	// copy the OS status over
	lasthostname := hs.hostname.Get()
	hostname := me.statusOS.GetHostname()

	// hostname changed or was setup for the first time. Set the window title, etc
	if lasthostname != hostname {
		me.changed = true
		hs.hostname.Set(hostname)
		hs.window.Title(hostname + " Status")
		me.statusDNSbutton.Set(hostname + " status")
	}
	hs.domainname.Set(me.statusOS.GetDomainName())

	tmp := me.statusOS.GetIPv4()
	sort.Strings(tmp)
	hs.currentIPv4.Set(strings.Join(tmp, "\n"))

	tmp = me.statusOS.GetIPv6()
	sort.Strings(tmp)
	hs.currentIPv6.Set(strings.Join(tmp, "\n"))

	if me.statusOS.ValidHostname() {
		vals = lookupDoH(me.statusOS.GetHostname(), "AAAA")

		log.Log(STATUS, "DNS IPv6 Addresses for ", me.statusOS.GetHostname(), "=", vals)
		if len(vals) == 0 {
			s = "(none)"
		} else {
			hs.setIPv6("Check for real IPv6 addresses here")
			/*
			if hs.missingAAAA() {
				hs.setIPv6("Add the missing IPv6 address")
			}
			*/
			for _, addr := range vals {
				log.Log(STATUS, addr)
				s += addr + " (DELETE)" + "\n"
				hs.setIPv6("NEEDS DELETE")
				// hs.dnsValue.SetText(addr)
				// hs.dnsAction.SetText("DELETE")
			}
		}
		sort.Strings(vals)
		hs.dnsAAAA.Set(strings.Join(vals, "\n"))

		vals = lookupDoH(me.statusOS.GetHostname(), "A")
		log.Log(STATUS, "IPv4 Addresses for ", me.statusOS.GetHostname(), "=", vals)
		s = strings.Join(vals, "\n")
		if (s == "") {
			s = "(none)"
			hs.setIPv4("NEEDS CNAME")
		}
		hs.set(hs.dnsA, s)

		vals = lookupDoH(me.statusOS.GetHostname(), "CNAME")
		s = strings.Join(vals, "\n")
		if (s != "") {
			hs.set(hs.dnsA, "CNAME " + s)
			hs.setIPv4("GOOD")
		}
	}

	if hs.IPv4() && hs.IPv6() {
		hs.status.Set("GOOD")
	} else {
		hs.status.Set("BROKEN")
	}

	// lookup the DNS provider
	// hs.dnsAPI.Set(me.DnsAPI.S)
	lookupNSprovider("wit.com")
}

func (hs *hostnameStatus) Show() {
	log.Log(STATUS, "hostnameStatus.Show() window")
	if hs.hidden {
		hs.window.Show()
	}
	hs.hidden = false
}

func (hs *hostnameStatus) Hide() {
	log.Log(STATUS, "hostnameStatus.Hide() window")
	if ! hs.hidden {
		hs.window.Hide()
	}
	hs.hidden = true
}

func (hs *hostnameStatus) GetDNSapi() string {
	return me.APIprovider
}
