// This creates a simple hello world window
package main

import 	(
	"time"
	"os"
	"os/user"
	"strconv"
//	"net"
	"strings"

	"go.wit.com/log"
	"go.wit.com/shell"

	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/gui/cloudflare"
	"go.wit.com/gui/debugger"
)

// This setups up the dns control panel window
func setupControlPanelWindow() {
	log.Info("artificial sleep of:", me.artificialSleep)
	log.Sleep(me.artificialSleep)

	// setup the main tab
	mainWindow("DNS and IPv6 Control Panel")
	detailsTab("Details")
	debugTab("Debug")

	// me.digStatus = NewDigStatusWindow(me.window)
}

func detailsTab(title string) {
	var g2 *gui.Node

	me.details = gadgets.NewBasicWindow(me.myGui, title)
	g2 = me.details.Box().NewGroup("Real Stuff")

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

	grid.Margin()
	grid.Pad()
}

func debugTab(title string) {
	var g2 *gui.Node

	win := gadgets.NewBasicWindow(me.myGui, title)

	g2 = win.Box().NewGroup("Real Stuff")

	g2.NewButton("GO GUI Debug Window", func () {
		debugger.DebugWindow(me.myGui)
	})

	g2.NewButton("getHostname() looks at the OS settings", func () {
		getHostname()
	})

	g2.NewButton("dig A & AAAA DNS records", func () {
		log.Println("updateDNS()")
		updateDNS()
	})

	g2.NewButton("dig +trace", func () {
		o := shell.Run("dig +trace +noadditional DS " + me.hostname + " @8.8.8.8")
		log.Println(o)
	})

	g2 = win.Box().NewGroup("debugging options")

	// makes a slider widget
	me.ttl = gadgets.NewDurationSlider(g2, "Loop Timeout", 10 * time.Millisecond, 5 * time.Second)
	me.ttl.Set(300 * time.Millisecond)

	// makes a slider widget
	me.dnsTtl = gadgets.NewDurationSlider(g2, "DNS Timeout", 800 * time.Millisecond, 300 * time.Second)
	me.dnsTtl.Set(60 * time.Second)

	g2.Margin()
	g2.Pad()
}

// will return a AAAA value that needs to be deleted
func deleteAAA() string {
	var aaaa []string
	aaaa = dhcpAAAA() // your AAAA IP addresses right now
	for _, s := range aaaa {
		log.Log(NOW, "DNS AAAA =", s)
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
		log.Log(NOW, "missing AAAA =", s)
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
		log.Log(NOW, "host", h, "DNS AAAA =", s, "ipmap[s] =", me.ipmap[s])
		all += s + "\n"
		if ( me.ipmap[s] == nil) {
			log.Warn("THIS IS THE WRONG AAAA DNS ENTRY:  host", h, "DNS AAAA =", s)
			broken = "wrong AAAA entry"
		} else {
			if (broken == "unknown") {
				broken = "needs update"
			}
		}
	}
	all = sortLines(all)
	if (me.workingIPv6.S != all) {
		log.Warn("workingIPv6.SetText() to:", all)
		me.workingIPv6.SetText(all)
	}

	var a []string
	a = realA()
	all = sortLines(strings.Join(a, "\n"))
	if (all == "") {
		log.Info("THERE IS NOT a real A DNS ENTRY")
		all = "CNAME ipv6.wit.com"
	}
	if (me.DnsA.S != all) {
		log.Warn("DnsA.SetText() to:", all)
		me.DnsA.SetText(all)
	}
	return broken
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

func mainWindow(title string) {
	me.window = gadgets.NewBasicWindow(me.myGui, title)

	me.mainStatus = me.window.Box().NewGroup("dns update")
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
			log.Info("hostname is good:", me.hostname)
		} else {
			log.Warn("FIX: you need to fix your hostname here", me.hostname)
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

	me.digStatusButton = me.mainStatus.NewButton("Resolver Status", func () {
		if (me.digStatus == nil) {
			log.Info("drawing the digStatus window START")
			me.digStatus = NewDigStatusWindow(me.myGui)
			log.Info("drawing the digStatus window END")
			me.digStatusButton.SetText("Hide DNS Lookup Status")
			me.digStatus.Update()
			return
		}
		if me.digStatus.hidden {
			me.digStatusButton.SetText("Hide Resolver Status")
			me.digStatus.Show()
			me.digStatus.Update()
		} else {
			me.digStatusButton.SetText("Resolver Status")
			me.digStatus.Hide()
		}
	})
	me.hostnameStatusButton = me.mainStatus.NewButton("Show hostname DNS Status", func () {
		if (me.hostnameStatus == nil) {
			me.hostnameStatus = NewHostnameStatusWindow(me.myGui)
			me.hostnameStatusButton.SetText("Hide " + me.hostname + " DNS Status")
			me.hostnameStatus.Update()
			return
		}
		if me.hostnameStatus.hidden {
			me.hostnameStatusButton.SetText("Hide " + me.hostname + " DNS Status")
			me.hostnameStatus.Show()
			me.hostnameStatus.Update()
		} else {
			me.hostnameStatusButton.SetText("Show " + me.hostname + " DNS Status")
			me.hostnameStatus.Hide()
		}
	})

	grid.Margin()
	grid.Pad()

	statusGrid(me.window.Box())

	gr := me.window.Box().NewGroup("debugging")
	gr.NewButton("GO GUI Debugger", func () {
		debugger.DebugWindow(me.myGui)
	})
	gr.NewButton("Details", func () {
		me.details.Toggle()
	})
}

func statusGrid(n *gui.Node) {
	problems := n.NewGroup("status")

	gridP := problems.NewGrid("nuts", 2, 2)

	gridP.NewLabel("DNS Status =")
	me.DnsStatus = gridP.NewLabel("unknown")

	me.statusIPv6 = gadgets.NewOneLiner(gridP, "IPv6 working")
	me.statusIPv6.Set("known")

	gridP.NewLabel("hostname =")
	me.hostnameStatusOLD = gridP.NewLabel("invalid")

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
	me.hostnameStatus.Update()

	// log.Println("digAAAA()")
	aaaa = digAAAA(h)
	log.Log(NOW, "digAAAA() =", aaaa)

	// log.Println(SPEW, me)
	if (aaaa == nil) {
		log.Warn("There are no DNS AAAA records for hostname: ", h)
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
		log.Warn("You're DNS lookups are very slow")
		me.dbOn.Set(true)
		me.dbProc.Set(true)

		processName := getProcessNameByPort(53)
		log.Info("Process with port 53:", processName)
	})
	// me.fixProc.Disable()
}
