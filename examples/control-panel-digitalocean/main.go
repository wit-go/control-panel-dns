package main

import (
	"go.wit.com/log"
	"go.wit.com/gui"
	"go.wit.com/control-panel-dns/digitalocean"
)

var title string = "Cloud App"
var myGui *gui.Node
var myDo *digitalocean.DigitalOcean

func main() {
	// initialize a new GO GUI instance
	myGui = gui.New().Default()

	// draw the main window
	cloudApp(myGui)

	log.Sleep(1)
	myDo = digitalocean.New(myGui)
	myDo.Update()
	myDo.Show()

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()
}

func cloudApp(n *gui.Node) *gui.Node {
	win := n.NewWindow(title)

	// make a group label and a grid
	group := win.NewGroup("data").Pad()
	grid := group.NewGrid("grid", 2, 1).Pad()

	grid.NewButton("New()", func () {
		myDo = digitalocean.New(myGui)
	})
	grid.NewLabel("initializes the DO golang gui package")

	grid.NewButton("Show", func () {
		myDo.Show()
	})
	grid.NewLabel("will show the DO window")

	grid.NewButton("Hide", func () {
		myDo.Hide()
	})
	grid.NewLabel("will hide the DO window")

	grid.NewButton("Update", func () {
		myDo.Update()
	})
	grid.NewLabel("polls DO via the API to find the state of all your droplets")

	grid.NewButton("Create", func () {
		// myDo.Create("jcarr.wit.com")
		digitalocean.InitCreateWindow()
	})
	grid.NewLabel("makes a new droplet")

	return win
}
