// This creates a simple hello world window
package main

import 	(
	"os"
	"os/user"
	"log"
	"net"
	"git.wit.org/wit/gui"
	"github.com/davecgh/go-spew/spew"
)

// This initializes the first window
func initGUI() {
	var w *gui.Node
	gui.Config.Title = "Hello World golang wit/gui Window"
	gui.Config.Width = 640
	gui.Config.Height = 480
	gui.Config.Exit = myDefaultExit

	w = gui.NewWindow()
	w.Dump()
	addDemoTab(w, "A Simple Tab Demo")

	// TODO: add these back
	if (args.GuiDemo) {
		gui.DemoWindow()
	}

	if (args.GuiDebug) {
		gui.DebugWindow()
	}
}

func addDemoTab(window *gui.Node, title string) {
	var newNode, g, g2, tb *gui.Node
	var err error
	var name string

	newNode = window.NewTab(title)
        log.Println("addDemoTab() newNode.Dump")
	newNode.Dump()

	g = newNode.NewGroup("group 1")
	dd := g.NewDropdown("demoCombo2")
	dd.AddDropdownName("more 1")
	dd.AddDropdownName("more 2")
	dd.AddDropdownName("more 3")
	dd.OnChanged = func(*gui.Node) {
		s := dd.GetText()
		tb.SetText("hello world " + args.User + "\n" + s)
	}

	g2 = newNode.NewGroup("group 2")
	tb = g2.NewTextbox("tb")
	log.Println("tb =", tb.GetText())
	tb.OnChanged = func(*gui.Node) {
		s := tb.GetText()
		log.Println("text =", s)
	}
	g2.NewButton("hello", func () {
		log.Println("world")
		scanInterfaces()
	})
	g2.NewButton("os.Hostname()", func () {
		name, err = os.Hostname()
		log.Println("name =", name, err)
	})
	g2.NewButton("os.User()", func () {
		user, _ := user.Current()
		spew.Dump(user)
		log.Println("os.Getuid =", os.Getuid())
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
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}
