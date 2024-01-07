// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package main

import 	(
	"fmt"
	"strings"
	"sort"
	"runtime"
	"time"
	"embed"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/debugger"

	"go.wit.com/control-panels/dns/linuxstatus"

	"github.com/miekg/dns"
)

//go:embed plugins/*.so
var resToolkit embed.FS

func main() {
	// parsedown()

	// initialize the maps to track IP addresses and network interfaces
	me.ipmap = make(map[string]*IPtype)
	me.dnsmap = make(map[string]*IPtype)
	me.ifmap = make(map[int]*IFtype)
	me.nsmap = make(map[string]string)

	// initialize maps for the returned DNS records
	me.ipv4s = make(map[string]dns.RR)
	me.ipv6s = make(map[string]dns.RR)

	// send all log() output to a file in /tmp
	log.SetTmp()

	me.myGui = gui.New().Default()

	log.Sleep(me.artificialSleep)
	setupControlPanelWindow()

	me.digStatus = NewDigStatusWindow(me.myGui)
	me.statusDNS = NewHostnameStatusWindow(me.myGui)

	linuxstatus.New()

	if debugger.ArgDebug() {
		log.Sleep(2)
		debugger.DebugWindow(me.myGui)
	}

	// forever monitor for network and dns changes
	log.Sleep(me.artificialSleep)
	checkNetworkChanges()
}

/*
	Poll for changes to the networking settings
*/

/* https://github.com/robfig/cron/blob/master/cron.go

// Run the cron scheduler, or no-op if already running.
func (c *Cron) Run() {
	c.runningMu.Lock()
	if c.running {
		c.runningMu.Unlock()
		return
	}
	c.running = true
	c.runningMu.Unlock()
	c.run()
}

// run the scheduler.. this is private just due to the need to synchronize
// access to the 'running' state variable.
func (c *Cron) run() {
	c.logger.Info("start")
*/

func checkNetworkChanges() {
	var lastLocal time.Time = time.Now()
	var lastDNS time.Time = time.Now()

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	for {
		time.Sleep(me.ttl.Duration)
		if (time.Since(lastLocal) > me.localSleep) {
			if (runtime.GOOS == "linux") {
				duration := timeFunction(linuxLoop)
				s := fmt.Sprint(duration)
				me.statusOS.SetSpeedActual(s)
			} else {
				// TODO: make windows and macos diagnostics
				log.Warn("Windows and MacOS don't work yet")
			}
			lastLocal = time.Now()
		}
		if (time.Since(lastDNS) > me.dnsTtl.Duration) {
			DNSloop()
			lastDNS = time.Now()
		}

		/*
		stop2 := timer2.Stop()
		if stop2 {
			fmt.Println("Timer 2 stopped")
		}
		*/
	}
}

// run this on each timeout
func DNSloop() {
	duration := timeFunction(dnsTTL)
	log.Info("dnsTTL() execution Time: ", duration)
	var s, newSpeed string
	if (duration > 5000 * time.Millisecond ) {
		newSpeed = "VERY SLOW"
	} else if (duration > 2000 * time.Millisecond ) {
		newSpeed = "SLOWER"
	} else if (duration > 500 * time.Millisecond ) {
		newSpeed = "SLOW"
	} else if (duration > 100 * time.Millisecond ) {
		newSpeed = "OK"
	} else {
		newSpeed = "FAST"
	}
	if (newSpeed != me.DnsSpeedLast) {
		log.Log(CHANGE, "dns lookup speed changed =", newSpeed)
		log.Log(CHANGE, "dnsTTL() execution Time: ", duration)
		me.DnsSpeed.SetText(newSpeed)
		me.DnsSpeedLast = newSpeed
	}
	s = fmt.Sprint(duration)
	me.DnsSpeedActual.SetText(s)
}

// This checks for changes to the network settings
// and verifies that DNS is working or not working
func dnsTTL() {
	updateDNS()
}

// run update on the LinuxStatus() window
func linuxLoop() {
	me.statusOS.Update()

	if (me.statusOS.Changed()) {
		stamp := time.Now().Format("2006/01/02 15:04:05")
		log.Log(CHANGE, "Network things changed on", stamp)
		duration := timeFunction(updateDNS)
		log.Log(CHANGE, "updateDNS() execution Time: ", duration)
	}
}

/*
	// Example usage
	duration := timeFunction(FunctionToTime)
	log.Println("Execution Time: ", duration)
*/

// timeFunction takes a function as an argument and returns the execution time.
func timeFunction(f func()) time.Duration {
	startTime := time.Now() // Record the start time
	f()                     // Execute the function
	return time.Since(startTime) // Calculate the elapsed time
}

// sortLines takes a string, splits it on newlines, sorts the lines,
// and rejoins them with newlines.
func sortLines(input string) string {
	lines := strings.Split(input, "\n")

	// Trim leading and trailing whitespace from each line
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	sort.Strings(lines)
	tmp := strings.Join(lines, "\n")
	tmp = strings.TrimLeft(tmp, "\n")
	tmp = strings.TrimRight(tmp, "\n")
	return tmp
}
