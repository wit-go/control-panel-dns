// This creates a simple hello world window
package main

import 	(
	"log"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"net"
	"git.wit.org/wit/gui"
	"git.wit.org/wit/shell"
	"github.com/davecgh/go-spew/spew"
)

// This setups up the dns control panel window
func setupControlPanelWindow() {
	// me.window = myGui.New2().Window("DNS and IPv6 Control Panel").Standard()
	me.window = myGui.NewWindow("DNS and IPv6 Control Panel").Standard()
	me.window.Dump()

	sleep(me.artificialSleep)
	dnsTab("DNS")
	debugTab("Debug")

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
	g2.NewButton("Actual AAAA", func () {
		var aaaa []string
		aaaa = realAAAA()
		for _, s := range aaaa {
			log.Println("my actual AAAA = ", s)
		}
	})

	g2.NewButton("Update DNS", func () {
		log.Println("updateDNS()")
		updateDNS()
	})

	g2.NewButton("checkDNS()", func () {
		checkDNS()
	})
	g2.NewButton("os.User()", func () {
		user, _ := user.Current()
		spew.Dump(user)
		log.Println("os.Getuid =", user.Username, os.Getuid())
		if (me.uid != nil) {
			me.uid.SetText(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
		}
	})
	g2.NewButton("dig +trace", func () {
		o := shell.Run("dig +trace +noadditional DS " + me.hostname + " @8.8.8.8")
		log.Println(o)
		// log.Println(o)
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

	g2.NewLabel("control panel TTL (in tenths of seconds)")
	ttl := g2.NewSlider("dnsTTL", 1, 100)
	ttl.Set(me.dnsTTL * 10)
	ttl.Custom = func () {
		me.dnsTTL = ttl.I / 10
		log.Println("dnsTTL =", me.dnsTTL)
	}

	g2.NewLabel("control panel loop delay (in tenths of seconds)")
	ttl2 := g2.NewSlider("dnsTTL", 1, 100)
	ttl2.Set(me.dnsTTLsleep)
	ttl2.Custom = func () {
		me.dnsTTLsleep = float64(ttl2.I) / 10
		log.Println("dnsTTLsleep =", me.dnsTTLsleep)
	}
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}

func dnsTab(title string) {
	tab := me.window.NewTab(title)

	g := tab.NewGroup("dns update")

	grid := g.NewGrid("gridnuts", 2, 2)

	grid.SetNext(1,1)
	grid.NewLabel("hostname =")
	me.fqdn = grid.NewLabel("?")
	me.hostname = ""

	grid.NewLabel("UID =")
	me.uid = grid.NewLabel("?")

	grid.NewLabel("DNS AAAA =")
	me.DnsAAAA = grid.NewLabel("?")

	grid.NewLabel("DNS A =")
	me.DnsA = grid.NewLabel("?")

	grid.NewLabel("IPv4 =")
	me.IPv4 = grid.NewLabel("?")

	grid.NewLabel("IPv6 =")
	me.IPv6 = grid.NewLabel("?")

	grid.NewLabel("interfaces =")
	me.Interfaces = grid.NewCombobox("Interfaces")

	grid.NewLabel("DNS Status =")
	me.DnsStatus = grid.NewLabel("unknown")

	me.fix = g.NewButton("Fix", func () {
		if (goodHostname(me.hostname)) {
			log.Println("hostname is good:", me.hostname)
		} else {
			log.Println("you need to fix your hostname here", me.hostname)
			return
		}
		nsupdate()
	})
	me.fix.Disable()

	statusGrid(tab)

}

func statusGrid(n *gui.Node) {
	problems := n.NewGroup("status")

	gridP := problems.NewGrid("nuts", 2, 2)

	gridP.NewLabel("DNS Status =")
	gridP.NewLabel("unknown")

	gridP.NewLabel("hostname =")
	gridP.NewLabel("invalid")

	gridP.NewLabel("dns provider =")
	gridP.NewLabel("unknown")

	gridP.NewLabel("IPv6 working =")
	gridP.NewLabel("unknown")

	gridP.NewLabel("dns resolution =")
	gridP.NewLabel("unknown")
}

/*
var outJunk string
func output(s string, a bool) {
	if (a) {
		outJunk += s
	} else {
		outJunk = s
	}
	me.output.SetText(outJunk)
	log.Println(outJunk)
}
*/

func updateDNS() {
	var aaaa []string
	h := me.hostname
	if (h == "") {
		h = "unknown.lab.wit.org"
		// h = "hpdevone.lab.wit.org"
	}
	log.Println("dnsAAAA()()")
	aaaa = dnsAAAA(h)
	log.Println("dnsAAAA()()")
	log.Println(SPEW, me)
	if (aaaa == nil) {
		log.Println("There are no DNS AAAA records for hostname: ", h)
	}
	var broken int = 0
	var all string
	for _, s := range aaaa {
		log.Println("host", h, "DNS AAAA =", s)
		all += s + "\n"
		if ( me.ipmap[s] == nil) {
			log.Println("THIS IS THE WRONG AAAA DNS ENTRY:  host", h, "DNS AAAA =", s)
			broken = 2
		} else {
			if (broken == 0) {
				broken = 1
			}
		}
	}
	all = strings.TrimSpace(all)
	me.DnsAAAA.SetText(all)
	if (broken == 1) {
		me.DnsStatus.SetText("WORKING")
	} else {
		me.DnsStatus.SetText("BROKEN")
		me.fix.Enable()
		log.Println("Need to run go-nsupdate here")
		nsupdate()
	}

	user, _ := user.Current()
	spew.Dump(user)
	log.Println("os.Getuid =", user.Username, os.Getuid())
	if (me.uid != nil) {
		me.uid.SetText(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
	}
	log.Println("updateDNS() END")
}
