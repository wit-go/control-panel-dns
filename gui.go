// This creates a simple hello world window
package main

import 	(
	"os"
	"os/user"
	"strconv"
	"net"
	"git.wit.org/wit/gui"
	"git.wit.org/wit/shell"
	"github.com/davecgh/go-spew/spew"
)

// This initializes the first window
func initGUI() {
	gui.Config.Title = "DNS and IPv6 Control Panel"
	gui.Config.Width = 1024
	gui.Config.Height = 480
	gui.Config.Exit = myDefaultExit

	me.window = gui.NewWindow()
	addDNSTab("DNS")

	if (args.GuiDebug) {
		gui.DebugWindow()
	}
	// gui.ShowDebugValues()
}

func addDNSTab(title string) {
	var g2 *gui.Node

	me.tab = me.window.NewTab(title)
        // log("addDemoTab() newNode.Dump")
	// newNode.Dump()

	g2 = me.tab.NewGroup("Real Stuff")

	g2.NewButton("Network Interfaces", func () {
		for i, t := range me.ifmap {
			log("name =", t.iface.Name)
			log("int =", i, "name =", t.name, t.iface)
			log("iface = " + t.iface.Name)
		}
	})
	g2.NewButton("Hostname", func () {
		getHostname()
	})
	g2.NewButton("Actual AAAA", func () {
		var aaaa []string
		aaaa = realAAAA()
		for _, s := range aaaa {
			log("my actual AAAA = ", s)
		}
	})

	g2.NewButton("checkDNS()", func () {
		checkDNS()
	})
	g2.NewButton("os.User()", func () {
		user, _ := user.Current()
		spew.Dump(user)
		log("os.Getuid =", user.Username, os.Getuid())
		if (me.uid != nil) {
			me.uid.AddText(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
			me.uid.SetText(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
		}
	})
	g2.NewButton("Example_listLink()", func () {
		Example_listLink()
	})
	g2.NewButton("Escalate()", func () {
		Escalate()
	})
	g2.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
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

	nsupdateGroup(me.tab)

	tmp := me.tab.NewGroup("output")
	me.output = tmp.NewTextbox("some output")
	me.output.Custom = func() {
		s := me.output.GetText()
		log("output text =", s)
	}
}

func myDefaultExit(n *gui.Node) {
        log("You can Do exit() things here")
	os.Exit(0)
}

func nsupdateGroup(w *gui.Node) {
	g := w.NewGroup("dns update")

	grid := g.NewGrid("fucknuts", 2, 2)

	grid.SetNext(1,1)
	grid.NewLabel("hostname =")
	// grid.SetNext(1,2)
	me.fqdn = grid.NewCombobox("foo(0,1)")
	me.fqdn.AddText("fire.lab.wit.com")
	me.fqdn.AddText("mirrors.wit.com")
	me.fqdn.SetText("sad.lab.wit.org")

	// grid.SetNext(2,1)
	grid.NewLabel("UID =")
	// grid.SetNext(2,2)
	me.uid = grid.NewCombobox("foo(1,1)")
	me.uid.AddText("root (0)")
	me.uid.AddText("mail (8)")
	me.uid.AddText("jcarr (1000)")

	grid.NewLabel("IPv4 =")
	me.IPv4 = grid.NewCombobox("foo(2,1)")

	grid.NewLabel("IPv6 =")
	me.IPv6 = grid.NewCombobox("foo(1,3)")

	grid.NewLabel("interfaces =")
	me.Interfaces = grid.NewCombobox("foo(1,3)")

	grid.NewLabel("DNS AAAA =")
	me.DnsAAAA = grid.NewLabel("need to lookup")

	grid.NewLabel("DNS A =")
	me.DnsA = grid.NewLabel("need to lookup")

	grid.NewLabel("DNS Status =")
	me.DnsStatus = grid.NewLabel("unknown")

	g.NewButton("DNS AAAA", func () {
		var aaaa []string
		h := me.hostname
		if (h == "") {
			h = "unknown.lab.wit.org"
			// h = "hpdevone.lab.wit.org"
		}
		aaaa = dnsAAAA(h)
		log(SPEW, me)
		if (aaaa == nil) {
			log("There are no DNS AAAA records for hostname: ", h)
		}
		var broken int = 0
		var all string
		for _, s := range aaaa {
			log("host", h, "DNS AAAA =", s)
			all += s + "\n"
			if ( me.ipmap[s] == nil) {
				log("THIS IS THE WRONG AAAA DNS ENTRY:  host", h, "DNS AAAA =", s)
				broken = 2
			} else {
				if (broken == 0) {
					broken = 1
				}
			}
		}
		me.DnsAAAA.SetText(all)
		if (broken == 1) {
			me.DnsStatus.SetText("WORKING")
		} else {
			me.DnsStatus.SetText("Broken")
		}
	})
	g.NewButton("dig +trace", func () {
		o := shell.Run("dig +trace +noadditional DS " + me.fqdn.GetText() + " @8.8.8.8")
		output(o, false)
		// log(o)
	})
}

var outJunk string
func output(s string, a bool) {
	if (a) {
		outJunk += s
	} else {
		outJunk = s
	}
	me.output.SetText(outJunk)
	log(outJunk)
}
