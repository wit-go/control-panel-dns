// This creates a simple hello world window
package main

import 	(
	"time"
	"os"

	"go.wit.com/log"

	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/gui/cloudflare"
	"go.wit.com/gui/debugger"
	"go.wit.com/gui/gadgets/logsettings"
	// "go.wit.com/control-panels/dns/linuxstatus"
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

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

func mainWindow(title string) {
	me.window = gadgets.NewBasicWindow(me.myGui, title)

	gr := me.window.Box().NewGroup("dns update")

	// This is where you figure out what to do next to fix the problems
	me.fixButton = gr.NewButton("Check Errors", func () {
		if ! fix() {
			log.Log(CHANGE, "boo. IPv6 isn't working yet")
			return
		}
		log.Log(CHANGE, "IPv6 WORKED")
		// update everything here visually for the user
		// hostname := me.statusOS.GetHostname()
		// me.hostname.Set(hostname)
		me.hostnameStatus.Set("WORKING")
		me.DnsStatus.Set("WORKING")
	})

	statusGrid(me.window.Box())

	gr = me.window.Box().NewGroup("debugging")
	me.statusDNSbutton = gr.NewButton("hostname status", func () {
		if ! me.statusDNS.Ready() {return}
		me.statusDNS.window.Toggle()
	})
	gr.NewButton("Linux Status", func () {
		me.statusOS.Toggle()
	})
	gr.NewButton("resolver status", func () {
		if ! me.digStatus.Ready() {return}
		me.digStatus.window.Toggle()
	})
	gr.NewButton("cloudflare wit.com", func () {
		if me.witcom != nil {
			me.witcom.Toggle()
		}
		me.witcom = cloudflare.CreateRR(me.myGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
	})
	gr.NewButton("Debug", func () {
		me.debug.Toggle()
	})

	var myLS *logsettings.LogSettings
	gr.NewButton("Logging Settings", func () {
		if myLS == nil {
			// initialize the log settings window (does not display it)
			myLS = logsettings.New(me.myGui)
			return
		}
		myLS.Toggle()
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
	me.statusDNS.Update()

	if me.digStatus.Ready() {
		if me.digStatus.IPv6() {
			me.statusIPv6.Set("IPv6 WORKING")
		} else {
			me.statusIPv6.Set("Need VPN")
		}
	}

	// lookup the NS records for your domain
	// if your host is test.wit.com, find the NS resource records for wit.com
	lookupNS(me.statusOS.GetDomainName())

	log.Println("updateDNS() END")
}
