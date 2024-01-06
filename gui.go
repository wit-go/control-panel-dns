// This creates a simple hello world window
package main

import 	(
	"time"
	"os"
	"strings"

	"go.wit.com/log"

	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/gui/cloudflare"
	"go.wit.com/gui/debugger"
	"go.wit.com/control-panels/dns/linuxstatus"
)

// This setups up the dns control panel window
func setupControlPanelWindow() {
	log.Info("artificial sleep of:", me.artificialSleep)
	log.Sleep(me.artificialSleep)

	// setup the main tab
	mainWindow("DNS and IPv6 Control Panel")
	debugTab("Debug")
}

func debugTab(title string) {
	var g2 *gui.Node

	me.debug = gadgets.NewBasicWindow(me.myGui, title)
	me.debug.Draw()
	me.debug.Hide()

	g2 = me.debug.Box().NewGroup("Real Stuff")

	g2.NewButton("GO GUI Debug Window", func () {
		debugger.DebugWindow(me.myGui)
	})

	g2.NewButton("dig A & AAAA DNS records", func () {
		log.Println("updateDNS()")
		updateDNS()
	})

	g2.NewButton("dig +trace", func () {
		log.Log(NOW, "TODO: redo this")
		// o := shell.Run("dig +trace +noadditional DS " + me.hostname + " @8.8.8.8")
		// log.Println(o)
	})

	g2.NewButton("getProcessNameByPort()", func () {
		processName := linuxstatus.GetProcessNameByPort(53)
		log.Info("Process with port 53:", processName)
	})

	g2 = me.debug.Box().NewGroup("debugging options")

	// makes a slider widget
	me.ttl = gadgets.NewDurationSlider(g2, "Loop Timeout", 10 * time.Millisecond, 5 * time.Second)
	me.ttl.Set(300 * time.Millisecond)

	// makes a slider widget
	me.dnsTtl = gadgets.NewDurationSlider(g2, "DNS Timeout", 800 * time.Millisecond, 300 * time.Second)
	me.dnsTtl.Set(60 * time.Second)

	g2.Margin()
	g2.Pad()

	me.debug.Hide()
}

/*
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
*/

// doesn't actually do any network traffic
// it just updates the GUI
func displayDNS() string {
	var aaaa []string
	aaaa = append(aaaa, "blah", "more")
	// h := me.hostname
	var all string
	var broken string = "unknown"
	for _, s := range aaaa {
		log.Log(STATUS, "host", "fixme", "DNS AAAA =", s, "ipmap[s] =", me.ipmap[s])
		all += s + "\n"
		if ( me.ipmap[s] == nil) {
			log.Warn("THIS IS THE WRONG AAAA DNS ENTRY:  host", "fixme", "DNS AAAA =", s)
			broken = "wrong AAAA entry"
		} else {
			if (broken == "unknown") {
				broken = "needs update"
			}
		}
	}

	var a []string
	a = append(a, "fixme")
	all = sortLines(strings.Join(a, "\n"))
	if (all == "") {
		log.Log(NOW, "THERE IS NOT a real A DNS ENTRY")
		all = "CNAME ipv6.wit.com"
	}
	if (me.DnsA.S != all) {
		log.Log(NOW, "DnsA.SetText() to:", all)
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

	gr := me.window.Box().NewGroup("dns update")
	grid := gr.NewGrid("gridnuts", 2, 2)

	grid.SetNext(1,1)

	me.hostname = gadgets.NewOneLiner(grid, "hostname =").Set("unknown")
	me.DnsAAAA = gadgets.NewOneLiner(grid, "DNS AAAA =").Set("unknown")

	grid.NewLabel("DNS A =")
	me.DnsA = grid.NewLabel("?")

	// This is where you figure out what to do next to fix the problems
	gr.NewButton("fix", func () {
		fix()
	})

	grid.Margin()
	grid.Pad()

	statusGrid(me.window.Box())

	gr = me.window.Box().NewGroup("debugging")
	gr.NewButton("hostname status", func () {
		if ! me.status.Ready() {return}
		me.status.window.Toggle()
	})

	gr.NewButton("linuxstatus.New()", func () {
		if (me.statusOS == nil) {
			me.statusOS = linuxstatus.New()
		}
		me.statusOS.SetParent(me.myGui)
		me.statusOS.InitWindow()
		me.statusOS.Make()
		me.statusOS.Draw2()
	})
	gr.NewButton("statusOS.Ready()", func () {
		me.statusOS.Ready()
	})
	gr.NewButton("statusOS.Draw()", func () {
		me.statusOS.Draw()
		me.statusOS.Draw2()
	})
	gr.NewButton("statusOS.Update()", func () {
		me.statusOS.Update()
	})
	gr.NewButton("Linux Status", func () {
		me.statusOS.Toggle()
	})
	gr.NewButton("resolver status", func () {
		if ! me.digStatus.Ready() {return}
		me.digStatus.window.Toggle()
	})
	gr.NewButton("cloudflare wit.com", func () {
		cloudflare.CreateRR(me.myGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
	})
	gr.NewButton("Debug", func () {
		me.debug.Toggle()
	})
}

func statusGrid(n *gui.Node) {
	problems := n.NewGroup("status")

	gridP := problems.NewGrid("nuts", 2, 2)

	gridP.NewLabel("hostname =")
	me.hostnameStatus = gridP.NewLabel("invalid")

	gridP.NewLabel("DNS Status =")
	me.DnsStatus = gridP.NewLabel("unknown")

	me.statusIPv6 = gadgets.NewOneLiner(gridP, "IPv6 working")
	me.statusIPv6.Set("known")

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
	me.digStatus.Update()
	me.status.Update()

	// log.Println("digAAAA()")

	if me.statusOS.ValidHostname() {
		var aaaa []string
		h := me.statusOS.GetHostname()
		aaaa = digAAAA(h)
		log.Log(NOW, "digAAAA() for", h, "=", aaaa)

		// log.Println(SPEW, me)
		if (aaaa == nil) {
			log.Warn("There are no DNS AAAA records for hostname: ", h)
			me.DnsAAAA.Set("(none)")
			if (cloudflare.CFdialog.TypeNode != nil) {
				cloudflare.CFdialog.TypeNode.SetText("AAAA new")
			}
	
			if (cloudflare.CFdialog.NameNode != nil) {
				cloudflare.CFdialog.NameNode.SetText(h)
			}
	
			/*
			d := deleteAAA()
			if (d != "") {
				if (cloudflare.CFdialog.ValueNode != nil) {
					cloudflare.CFdialog.ValueNode.SetText(d)
				}
			}
			*/
//			m := missingAAAA()
//			if (m != "") {
//				if (cloudflare.CFdialog.ValueNode != nil) {
//					cloudflare.CFdialog.ValueNode.SetText(m)
//				}
//				/*
//				 rr := &cloudflare.RRT{
//					Type:	"AAAA",
//					Name:	me.hostname,
//					Ttl:	"Auto",
//					Proxied:	false,
//					Content:	m,
//				}
//				cloudflare.Update(rr)
//				*/
//			}
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


	// lookup the NS records for your domain
	// if your host is test.wit.com, find the NS resource records for wit.com
	lookupNS(me.statusOS.GetDomainName())

	log.Println("updateDNS() END")
}
