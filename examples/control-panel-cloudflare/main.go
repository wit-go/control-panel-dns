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

func main() {
	// parse the config file
	readConfig()

	// initialize a new GO GUI instance
	myGui = gui.New().Default()

	// draw the cloudflare control panel window
	win := cloudflare.MakeCloudflareWindow(myGui)
	win.SetText(title)

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()

	// update the config file
	saveConfig()
}
