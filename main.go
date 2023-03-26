// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package main

import 	(
	"strconv"
	"runtime"
	"time"
	arg "github.com/alexflint/go-arg"
	"git.wit.org/wit/gui"
)

var p *arg.Parser
var myGui *gui.Node

func main() {
	p = arg.MustParse(&args)
	parsedown()

	// initialize the maps to track IP addresses and network interfaces
	me.ipmap = make(map[string]*IPtype)
	me.dnsmap = make(map[string]*IPtype)
	me.ifmap = make(map[int]*IFtype)
	me.dnsTTL = 5		// recheck DNS is working every 2 minutes // TODO: watch rx packets?


	log()
	log(true, "this is true")
	log(false, "this is false")
	sleep(.4)
	sleep(.3)
	sleep(.2)
	sleep("done scanning net")
	// exit("done scanning net")

	// Example_listLink()
	// exit()

	log("Toolkit = ", args.Toolkit)
	// gui.InitPlugins([]string{"andlabs"})
	// gui.SetDebug(true)
	// myGui = gui.Main(initGUI)
	myGui = gui.Start()
	sleep(1)
	initGUI()
	sleep(1)
	checkNetworkChanges()
}

/*
	Poll for changes to the networking settings
*/
func checkNetworkChanges() {
	var ttl int = 0
	var ttlsleep int = 5
	for {
		sleep(ttlsleep)
		ttl -= 1
		if (ttl < 0) {
			if (runtime.GOOS == "linux") {
				dnsTTL()
			} else {
				log("Windows and MacOS don't work yet")
			}
			ttl = me.dnsTTL
		}
	}
}

// Run this every once and a while
func dnsTTL() {
	me.changed = false
	log("FQDN =", me.fqdn.GetText())
	getHostname()
	scanInterfaces()
	for i, t := range me.ifmap {
		log(strconv.Itoa(i) + " iface = " + t.iface.Name)
	}

	var aaaa []string
	aaaa = realAAAA()
	var all string
	for _, s := range aaaa {
		log("my actual AAAA = ",s)
		all += s + "\n"
	}
	// me.IPv6.SetText(all)

	if (me.changed) {
		stamp := time.Now().Format("2006/01/02 15:04:05")
		s := stamp + " Network things changed"
		log(logError, "Network things changed on", stamp)
		updateDNS()
		me.output.SetText(s)

	}
}
