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
	"go.wit.com/control-panels/dns/smartwindow"
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

	g2.NewButton("dig A & AAAA DNS records (updateDNS())", func () {
		log.Log(CHANGE, "updateDNS() going to run:")
	})

	g2 = me.debug.Box().NewGroup("debugging options")
	gridP := g2.NewGrid("nuts", 2, 1)

	// makes a slider widget
	me.ttl = gadgets.NewDurationSlider(gridP, "Loop Timeout", 10 * time.Millisecond, 5 * time.Second)
	me.ttl.Set(300 * time.Millisecond)

	// makes a slider widget
	me.dnsTtl = gadgets.NewDurationSlider(gridP, "DNS Timeout", 800 * time.Millisecond, 300 * time.Second)
	me.dnsTtl.Set(60 * time.Second)

	gridP.NewLabel("dns resolution")
	me.DnsSpeed = gridP.NewLabel("unknown")

	gridP.NewLabel("dns resolution speed")
	me.DnsSpeedActual = gridP.NewLabel("unknown")

	gridP.NewLabel("Test speed")
	newGrid := gridP.NewGrid("nuts", 2, 1).Pad()

	g2.Margin()
	g2.Pad()

	newGrid.NewLabel("ping.wit.com =")
	newGrid.NewLabel("unknown")

	newGrid.NewLabel("ping6.wit.com =")
	newGrid.NewLabel("unknown")

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
		me.fixButton.SetText("No Errors!")
		me.fixButton.Disable()
	})

	statusGrid(me.window.Box())

	gr = me.window.Box().NewGroup("")
/*
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
*/
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
	gr.NewButton("Show Errors", func () {
		if me.fixWindow == nil {
			me.fixWindow = smartwindow.New()
			me.fixWindow.SetParent(me.myGui)
			me.fixWindow.Title("fix window")
			me.fixWindow.SetDraw(drawFixWindow)
			me.fixWindow.Vertical()
			me.fixWindow.Make()
			me.fixWindow.Draw()
			me.fixWindow.Hide()
			return
		}
		me.fixWindow.Toggle()
	})
}


func statusGrid(n *gui.Node) {
	problems := n.NewGroup("status")
	problems.Margin()
	problems.Pad()

	gridP := problems.NewGrid("nuts", 3, 1)
	gridP.Margin()
	gridP.Pad()

	gridP.NewLabel("hostname =")
	me.hostnameStatus = gridP.NewLabel("invalid")
	gridP.NewButton("Linux Status", func () {
		me.statusOS.Toggle()
	})

	me.statusIPv6 = gadgets.NewOneLiner(gridP, "DNS Lookup")
	me.statusIPv6.Set("known")
	gridP.NewButton("resolver status", func () {
		if ! me.digStatus.Ready() {return}
		me.digStatus.window.Toggle()
	})

	gridP.NewLabel("DNS Status")
	me.DnsStatus = gridP.NewLabel("unknown")
	me.statusDNSbutton = gridP.NewButton("hostname status", func () {
		if ! me.statusDNS.Ready() {return}
		me.statusDNS.window.Toggle()
	})

	gridP.NewLabel("DNS API")
	me.DnsAPIstatus = gridP.NewLabel("unknown")
	var apiButton *gui.Node
	apiButton = gridP.NewButton("unknown wit.com", func () {
		log.Log(CHANGE, "WHAT API ARE YOU USING?")
		provider := me.statusDNS.GetDNSapi()
		apiButton.SetText(provider + " wit.com")
		if provider == "cloudflare" {
			me.DnsAPIstatus.Set("WORKING")
			return

			if me.witcom != nil {
				me.witcom.Toggle()
			}
			me.witcom = cloudflare.CreateRR(me.myGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
		}
	})

	n.NewGroup("NOTES")

}
