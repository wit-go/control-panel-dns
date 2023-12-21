// This is a simple example
package main

import 	(
	"go.wit.com/gui"
	"go.wit.com/control-panel-dns/cloudflare"
)

var title string = "Cloudflare DNS Control Panel"
var outfile string = "/tmp/guilogfile"
var configfile string = ".config/wit/cloudflare"

var myGui *gui.Node

// var buttonCounter int = 5
// var gridW int = 5
// var gridH int = 3

// var mainWindow, more, more2 *gui.Node

// var cloudflareURL string = "https://api.cloudflare.com/client/v4/zones/"

/*
var zonedrop *gui.Node
var domainWidget *gui.Node
var masterSave *gui.Node

var zoneWidget *gui.Node
var authWidget *gui.Node
var emailWidget *gui.Node

var loadButton *gui.Node
var saveButton *gui.Node
*/

func main() {
	// parse the config file
	readConfig()

	// initialize a new GO GUI instance
	myGui = gui.New().Default()

	// draw the cloudflare control panel window
	cloudflare.MakeCloudflareWindow(myGui)

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()

	// update the config file
	saveConfig()
}
