// This creates a simple hello world window
package main

import 	(
	"os"
	"os/user"
	"net"
	"git.wit.org/wit/gui"
	"github.com/davecgh/go-spew/spew"
)

// This initializes the first window
func initGUI() {
	var w *gui.Node
	gui.Config.Title = "DNS and IPv6 Control Panel"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myDefaultExit

	w = gui.NewWindow()
	w.Dump()
	addDNSTab(w, "DNS")

	// TODO: add these back
	if (args.GuiDemo) {
		gui.DemoWindow()
	}

	if (args.GuiDebug) {
		gui.DebugWindow()
	}
}

func addDNSTab(window *gui.Node, title string) {
	var newNode, g, g2, tb *gui.Node

	newNode = window.NewTab(title)
        log("addDemoTab() newNode.Dump")
	newNode.Dump()

	g = newNode.NewGroup("junk")
	dd := g.NewDropdown("demoCombo2")
	dd.AddDropdownName("more 1")
	dd.AddDropdownName("more 2")
	dd.AddDropdownName("more 3")
	dd.OnChanged = func(*gui.Node) {
		s := dd.GetText()
		tb.SetText("hello world " + args.User + "\n" + s)
		log("text =", s)
	}
	g.NewLabel("UID =")
	g.NewButton("hello", func () {
		log("world")
	})


	g2 = newNode.NewGroup("Real Stuff")
	tb = g2.NewTextbox("tb")
	log("tb =", tb.GetText())
	tb.OnChanged = func(*gui.Node) {
		s := tb.GetText()
		log("text =", s)
	}
	g2.NewButton("Network Interfaces", func () {
		for i, t := range me.ifmap {
			log("name =", t.iface.Name)
			log("int =", i, "name =", t.name, t.iface)
			dd.AddDropdownName(t.iface.Name)
		}
	})
	g2.NewButton("Hostname", func () {
		getHostname()
		g.NewLabel("FQDN = " + me.fqdn)
	})
	g2.NewButton("Actual AAAA", func () {
		var aaaa []string
		aaaa = realAAAA()
		for _, s := range aaaa {
			g.NewLabel("my actual AAAA = " + s)
		}
	})
	g2.NewButton("checkDNS()", func () {
		checkDNS()
	})
	g2.NewButton("os.User()", func () {
		user, _ := user.Current()
		spew.Dump(user)
		log("os.Getuid =", os.Getuid())
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
		log("host =", host)
	})
	g2.NewButton("DumpPublicDNSZone(apple.com)", func () {
		DumpPublicDNSZone("apple.com")
		dumpIPs("www.apple.com")
	})
}

func myDefaultExit(n *gui.Node) {
        log("You can Do exit() things here")
	os.Exit(0)
}
