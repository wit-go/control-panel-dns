// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package main

import 	(
	"strconv"
	"runtime"
	// "net"
	arg "github.com/alexflint/go-arg"
	"git.wit.org/wit/gui"
)

var p *arg.Parser

func main() {
	p = arg.MustParse(&args)
	parsedown()

	// initialize the maps to track IP addresses and network interfaces
	me.ipmap = make(map[string]*IPtype)
	me.dnsmap = make(map[string]*IPtype)
	me.ifmap = make(map[int]*IFtype)
	me.dnsTTL = 5		// recheck DNS is working every 2 minutes // TODO: watch rx packets?

	go checkNetworkChanges()

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
	gui.SetDebug(args.GuiDebug)
	gui.Main(initGUI)
}

/*
	Poll for changes to the networking settings
*/
func checkNetworkChanges() {
	var ttl int = 3
	for {
		sleep(1)
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
	output("FQDN = " + me.fqdn.GetText() + "\n", true)
	getHostname()
	scanInterfaces()
	for i, t := range me.ifmap {
		output(strconv.Itoa(i) + " iface = " + t.iface.Name + "\n", true)
	}
	var aaaa []string
	aaaa = realAAAA()
	for _, s := range aaaa {
		output("my actual AAAA = " + s + "\n", true)
	}
	// loggo()
}
