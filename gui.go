// This creates a simple hello world window
package main

import 	(
	"os"
	"log"
	"git.wit.org/wit/gui"
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
}

func myDefaultExit(n *gui.Node) {
        log.Println("You can Do exit() things here")
	os.Exit(0)
}
