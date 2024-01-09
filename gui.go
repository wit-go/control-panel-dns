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
	"go.wit.com/apps/control-panel-dns/smartwindow"
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
	me.window.Vertical()

	hbox := me.window.Box().NewBox("bw hbox", true)

	statusGrid(hbox)

	// some artificial padding to make the last row of buttons look less wierd
	gr := hbox.NewGroup("Development and Debugging Windows")
	gr = gr.NewBox("vbox", false)

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

	gr.NewButton("test smartwindow()", func () {
		if me.fixWindow == nil {
			me.fixWindow = smartwindow.New()
			me.fixWindow.SetParent(me.myGui)
			me.fixWindow.Title("smart window test")
			me.fixWindow.SetDraw(testSmartWindow)
			me.fixWindow.Vertical()
			me.fixWindow.Make()
			me.fixWindow.Draw()
			me.fixWindow.Hide()
			return
		}
		me.fixWindow.Toggle()
	})

	gr.NewButton("Show Errors", func () {
		me.problems.Toggle()
	})
	me.autofix = gr.NewCheckbox("Auto-correct Errors")
	me.autofix.Set(false)

	// These are your problems
	me.problems = NewErrorBox(me.window.Box(), "Errors", "has problems?")
	me.problems.addIPerror(RR, USER, "1:1:1:1")
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
	me.apiButton = gridP.NewButton("unknown wit.com", func () {
		log.Log(CHANGE, "WHAT API ARE YOU USING?")
		provider := me.statusDNS.GetDNSapi()
		if provider == "cloudflare" {
			if me.witcom != nil {
				me.witcom.Toggle()
			} else {
				me.witcom = cloudflare.CreateRR(me.myGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
			}
		}
	})
}
