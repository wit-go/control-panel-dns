// This creates a simple hello world window
package main

import 	(
	"fmt"
	"time"
	"os"
	"os/user"
	"strconv"
	"net"
	"strings"

	"go.wit.com/gui"
	"go.wit.com/log"
	"go.wit.com/shell"
	"go.wit.com/control-panel-dns/cloudflare"
)

// This setups up the dns control panel window
func setupControlPanelWindow() {
	me.window = myGui.NewWindow("DNS and IPv6 Control Panel")
	// me.window.Dump() // will dump out some info

	debug("artificial sleep of:", me.artificialSleep)
	sleep(me.artificialSleep)

	// setup the main tab
	dnsTab("DNS")
	detailsTab("Details")
	debugTab("Debug")

	// me.digStatus = NewDigStatusWindow(me.window)
}

func detailsTab(title string) {
	var g2 *gui.Node

	tab := me.window.NewTab(title)

	g2 = tab.NewGroup("Real Stuff")

	grid := g2.NewGrid("gridnuts", 2, 2)

	grid.SetNext(1,1)

	grid.NewLabel("domainname =")
	me.domainname = grid.NewLabel("domainname")

	grid.NewLabel("hostname -s =")
	me.hostshort = grid.NewLabel("hostname -s")

	grid.NewLabel("NS records =")
	me.NSrr = grid.NewLabel("NS RR's")

	grid.NewLabel("UID =")
	me.uid = grid.NewLabel("my uid")

	grid.NewLabel("Current IPv4 =")
	me.IPv4 = grid.NewLabel("?")

	grid.NewLabel("Current IPv6 =")
	me.IPv6 = grid.NewLabel("?")

	grid.NewLabel("Working Real IPv6 =")
	me.workingIPv6 = grid.NewLabel("?")

	grid.NewLabel("interfaces =")
	me.Interfaces = grid.NewCombobox("Interfaces")

	grid.NewLabel("refresh speed")
	me.LocalSpeedActual = grid.NewLabel("unknown")

	tab.Margin()
	tab.Pad()

	grid.Margin()
	grid.Pad()
}

func debugTab(title string) {
	var g2 *gui.Node

	tab := me.window.NewTab(title)

	g2 = tab.NewGroup("Real Stuff")

	g2.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
	})

	g2.NewButton("Load 'gocui'", func () {
		// this set the xterm and mate-terminal window title. maybe works generally?
		fmt.Println("\033]0;" + title + "blah \007")
		myGui.LoadToolkit("gocui")
	})

	g2.NewButton("Network Interfaces", func () {
		for i, t := range me.ifmap {
			log.Println("name =", t.iface.Name)
			log.Println("int =", i, "name =", t.name, t.iface)
			log.Println("iface = " + t.iface.Name)
		}
	})

	g2.NewButton("Hostname", func () {
		getHostname()
	})

	g2.NewButton("Actual AAAA & A", func () {
		displayDNS() // doesn't re-query anything
	})

	g2.NewButton("dig A & AAAA DNS records", func () {
		log.Println("updateDNS()")
		updateDNS()
	})

	g2.NewButton("checkDNS:", func () {
		ipv6s, ipv4s := checkDNS()
		for s, _ := range ipv6s {
			debug(LogNow, "check if", s, "is in DNS")
		}
		for s, _ := range ipv4s {
			debug(LogNow, "check if", s, "is in DNS")
		}
	})

	g2.NewButton("os.User()", func () {
		user, _ := user.Current()
		log.Println("os.Getuid =", user.Username, os.Getuid())
		if (me.uid != nil) {
			me.uid.SetText(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
		}
	})

	g2.NewButton("dig +trace", func () {
		o := shell.Run("dig +trace +noadditional DS " + me.hostname + " @8.8.8.8")
		log.Println(o)
	})

	g2.NewButton("Example_listLink()", func () {
		Example_listLink()
	})

	g2.NewButton("Escalate()", func () {
		Escalate()
	})

	g2.NewButton("LookupAddr(<raw ipv6>) == fire from /etc/hosts", func () {
		host, err := net.LookupAddr("2600:1700:afd5:6000:b26e:bfff:fe80:3c52")
		if err != nil {
			return
		}
		log.Println("host =", host)
	})

	g2.NewButton("DumpPublicDNSZone(apple.com)", func () {
		DumpPublicDNSZone("apple.com")
		dumpIPs("www.apple.com")
	})

	g2 = tab.NewGroup("debugging options")

	// DEBUG flags
	me.dbOn = g2.NewCheckbox("turn on debugging (will override all flags below)")
	me.dbOn.Custom = func() {
		DEBUGON = me.dbOn.B
	}

	me.dbNet = g2.NewCheckbox("turn on network debugging)")
	me.dbNet.Custom = func() {
		LogNet = me.dbNet.B
	}

	me.dbProc = g2.NewCheckbox("turn on /proc debugging)")
	me.dbProc.Custom = func() {
		LogProc = me.dbProc.B
	}

	// makes a slider widget
	me.ttl = cloudflare.NewDurationSlider(g2, "Loop Timeout", 10 * time.Millisecond, 5 * time.Second)
	me.ttl.Set(300 * time.Millisecond)

	// makes a slider widget
	me.dnsTtl = cloudflare.NewDurationSlider(g2, "DNS Timeout", 800 * time.Millisecond, 300 * time.Second)
	me.dnsTtl.Set(60 * time.Second)

	g2.Margin()
	g2.Pad()
}

// will return a AAAA value that needs to be deleted
func deleteAAA() string {
	var aaaa []string
	aaaa = dhcpAAAA() // your AAAA IP addresses right now
	for _, s := range aaaa {
		debug(LogNow, "DNS AAAA =", s)
		if ( me.ipmap[s] == nil) {
			return s
		}
	}
	return ""
}

// will return a AAAA value that needs to be added
func missingAAAA() string {
	var aaaa []string
	aaaa = dhcpAAAA() // your AAAA IP addresses right now
	for _, s := range aaaa {
		debug(LogNow, "missing AAAA =", s)
		return s
	}
	return ""
}

// doesn't actually do any network traffic
// it just updates the GUI
func displayDNS() string {
	var aaaa []string
	aaaa = dhcpAAAA() // your AAAA records right now
	h := me.hostname
	var all string
	var broken string = "unknown"
	for _, s := range aaaa {
		debug(LogNow, "host", h, "DNS AAAA =", s, "ipmap[s] =", me.ipmap[s])
		all += s + "\n"
		if ( me.ipmap[s] == nil) {
			debug(LogError, "THIS IS THE WRONG AAAA DNS ENTRY:  host", h, "DNS AAAA =", s)
			broken = "wrong AAAA entry"
		} else {
			if (broken == "unknown") {
				broken = "needs update"
			}
		}
	}
	all = sortLines(all)
	if (me.workingIPv6.S != all) {
		debug(LogError, "workingIPv6.SetText() to:", all)
		me.workingIPv6.SetText(all)
	}

	var a []string
	a = realA()
	all = sortLines(strings.Join(a, "\n"))
	if (all == "") {
		debug(LogInfo, "THERE IS NOT a real A DNS ENTRY")
		all = "CNAME ipv6.wit.com"
	}
	if (me.DnsA.S != all) {
		debug(LogError, "DnsA.SetText() to:", all)
		me.DnsA.SetText(all)
	}
	return broken
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

func dnsTab(title string) {
	tab := me.window.NewTab(title)

	me.mainStatus = tab.NewGroup("dns update")

	grid := me.mainStatus.NewGrid("gridnuts", 2, 2)

	grid.SetNext(1,1)

	grid.NewLabel("hostname =")
	me.fqdn = grid.NewLabel("?")
	me.hostname = ""

	grid.NewLabel("DNS AAAA =")
	me.DnsAAAA = grid.NewLabel("?")

	grid.NewLabel("DNS A =")
	me.DnsA = grid.NewLabel("?")

	me.fix = me.mainStatus.NewButton("Fix", func () {
		if (goodHostname(me.hostname)) {
			debug(LogInfo, "hostname is good:", me.hostname)
		} else {
			debug(LogError, "FIX: you need to fix your hostname here", me.hostname)
			return
		}
		// check to see if the cloudflare window exists
		/*
		if (me.cloudflareW != nil) {
			newRR.NameNode.SetText(me.hostname)
			newRR.TypeNode.SetText("AAAA")
			for s, t := range me.ipmap {
				if (t.IsReal()) {
					if (t.ipv6) {
						newRR.ValueNode.SetText(s)
						cloudflare.CreateCurlRR()
						return
					}
				}
			}
			cloudflare.CreateCurlRR()
			return
		} else {
		// nsupdate()
		// me.fixProc.Disable()
		}
		*/
	})
	me.fix.Disable()

	me.digStatusButton = me.mainStatus.NewButton("Show DNS Lookup Status", func () {
		if (me.digStatus == nil) {
			log.Info("drawing the digStatus window START")
			me.digStatus = NewDigStatusWindow(me.window)
			log.Info("drawing the digStatus window END")
			me.digStatusButton.SetText("Hide DNS Lookup Status")
		} else {
			if me.digStatus.hidden {
				me.digStatusButton.SetText("Hide DNS Lookup Status")
				me.digStatus.Show()
			} else {
				me.digStatusButton.SetText("Show DNS Lookup Status")
				me.digStatus.Hide()
			}
		}
	})

	grid.Margin()
	grid.Pad()

	statusGrid(tab)

}

func statusGrid(n *gui.Node) {
	problems := n.NewGroup("status")

	gridP := problems.NewGrid("nuts", 2, 2)

	gridP.NewLabel("DNS Status =")
	me.DnsStatus = gridP.NewLabel("unknown")

	me.statusIPv6 = cloudflare.NewOneLiner(gridP, "IPv6 working")
	me.statusIPv6.Set("known")

	gridP.NewLabel("hostname =")
	me.hostnameStatus = gridP.NewLabel("invalid")

	gridP.NewLabel("dns resolution")
	me.DnsSpeed = gridP.NewLabel("unknown")

	gridP.NewLabel("dns resolution speed")
	me.DnsSpeedActual = gridP.NewLabel("unknown")

	gridP.NewLabel("dns API provider =")
	me.DnsAPI = gridP.NewLabel("unknown")

	gridP.Margin()
	gridP.Pad()

	// TODO: these are notes for me things to figure out
	ng := n.NewGroup("TODO:")
	gridP = ng.NewGrid("nut2", 2, 2)

	gridP.NewLabel("ping.wit.com =")
	gridP.NewLabel("unknown")

	gridP.NewLabel("ping6.wit.com =")
	gridP.NewLabel("unknown")

	problems.Margin()
	problems.Pad()
	gridP.Margin()
	gridP.Pad()
}

// run everything because something has changed
func updateDNS() {
	var aaaa []string
	h := me.hostname
	if (h == "") {
		h = "test.wit.com"
	}

	me.digStatus.Update()

	// log.Println("digAAAA()")
	aaaa = digAAAA(h)
	debug(LogNow, "digAAAA() =", aaaa)

	// log.Println(SPEW, me)
	if (aaaa == nil) {
		debug(LogError, "There are no DNS AAAA records for hostname: ", h)
		me.DnsAAAA.SetText("(none)")
		if (cloudflare.CFdialog.TypeNode != nil) {
			cloudflare.CFdialog.TypeNode.SetText("AAAA new")
		}

		if (cloudflare.CFdialog.NameNode != nil) {
			cloudflare.CFdialog.NameNode.SetText(me.hostname)
		}

		d := deleteAAA()
		if (d != "") {
			if (cloudflare.CFdialog.ValueNode != nil) {
				cloudflare.CFdialog.ValueNode.SetText(d)
			}
		}
		m := missingAAAA()
		if (m != "") {
			if (cloudflare.CFdialog.ValueNode != nil) {
				cloudflare.CFdialog.ValueNode.SetText(m)
			}
			/*
			 rr := &cloudflare.RRT{
				Type:	"AAAA",
				Name:	me.hostname,
				Ttl:	"Auto",
				Proxied:	false,
				Content:	m,
			}
			cloudflare.Update(rr)
			*/
		}
	}
	status := displayDNS() // update the GUI based on dig results
	me.DnsStatus.SetText(status)

	if me.digStatus.Ready() {
		if me.digStatus.IPv6() {
			me.statusIPv6.Set("IPv6 WORKING")
		} else {
			me.statusIPv6.Set("Need VPN")
		}
	}


	// me.fix.Enable()

	user, _ := user.Current()
	log.Println("os.Getuid =", user.Username, os.Getuid())
	if (me.uid != nil) {
		me.uid.SetText(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
	}

	// lookup the NS records for your domain
	// if your host is test.wit.com, find the NS resource records for wit.com
	lookupNS(me.domainname.S)

	log.Println("updateDNS() END")
}

func suggestProcDebugging() {
	if (me.fixProc != nil) {
		// me.fixProc.Disable()
		return
	}

	me.fixProc = me.mainStatus.NewButton("Try debugging Slow DNS lookups", func () {
		debug("You're DNS lookups are very slow")
		me.dbOn.Set(true)
		me.dbProc.Set(true)

		DEBUGON = true
		LogProc = true
		processName := getProcessNameByPort(53)
		log.Println("Process with port 53:", processName)
	})
	// me.fixProc.Disable()
}
