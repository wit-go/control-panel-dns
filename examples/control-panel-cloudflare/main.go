package main

import 	(
	"go.wit.com/log"
	"go.wit.com/gui"
	"go.wit.com/control-panel-dns/cloudflare"
)

var title string = "Cloudflare DNS Control Panel"

var myGui *gui.Node

// var cloudflareURL string = "https://api.cloudflare.com/client/v4/zones/"

func main() {
	// send all log() output to a file in /tmp
	log.SetTmp()

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
