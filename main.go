// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package main

import 	(
	"log"
	"strconv"
	"runtime"
	"time"
	"embed"
	"git.wit.org/wit/gui"
)

var myGui *gui.Node

//go:embed plugins/*.so
var resToolkit embed.FS

func main() {
	// parsedown()

	// initialize the maps to track IP addresses and network interfaces
	me.ipmap = make(map[string]*IPtype)
	me.dnsmap = make(map[string]*IPtype)
	me.ifmap = make(map[int]*IFtype)
	me.dnsTTL = 2		// recheck DNS is working every 2 minutes // TODO: watch rx packets?

	// will set all debugging flags
	// gui.SetDebug(true)

	// myGui = gui.New().InitEmbed(resToolkit).LoadToolkit("gocui")
	myGui = gui.New().Default()
	sleep(me.artificialSleep)
	setupControlPanelWindow()
	sleep(me.artificialSleep)
	if (args.GuiDebug) {
		gui.DebugWindow()
	}
	gui.ShowDebugValues()

	// forever monitor for network and dns changes
	checkNetworkChanges()
}

/*
	Poll for changes to the networking settings
*/
func checkNetworkChanges() {
	var ttl int = 0
	for {
		sleep(me.dnsTTLsleep)
		ttl -= 1
		if (ttl < 0) {
			if (runtime.GOOS == "linux") {
				dnsTTL()
			} else {
				log.Println("Windows and MacOS don't work yet")
			}
			ttl = me.dnsTTL
		}
	}
}

// This checks for changes to the network settings
// and verifies that DNS is working or not working
func dnsTTL() {
	me.changed = false
	log.Println("FQDN =", me.fqdn.GetText())
	getHostname()
	scanInterfaces()
	for i, t := range me.ifmap {
		log.Println(strconv.Itoa(i) + " iface = " + t.iface.Name)
	}

	var aaaa []string
	aaaa = realAAAA()
	var all string
	for _, s := range aaaa {
		log.Println("my actual AAAA = ",s)
		all += s + "\n"
	}
	// me.IPv6.SetText(all)

	if (me.changed) {
		stamp := time.Now().Format("2006/01/02 15:04:05")
		log.Println(logError, "Network things changed on", stamp)
		updateDNS()
	}
}
