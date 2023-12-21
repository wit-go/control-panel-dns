// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package main

import 	(
	"fmt"
	"strings"
	"sort"
	"strconv"
	"runtime"
	"time"
	"embed"

	"go.wit.com/gui"
	"github.com/miekg/dns"
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
	me.nsmap = make(map[string]string)

	// initialize maps for the returned DNS records
	me.ipv4s = make(map[string]dns.RR)
	me.ipv6s = make(map[string]dns.RR)

	// will set all debugging flags
	// gui.SetDebug(true)

	// myGui = gui.New().InitEmbed(resToolkit).LoadToolkit("gocui")
	myGui = gui.New().Default()

	sleep(me.artificialSleep)
	setupControlPanelWindow()

	/*
	if (args.GuiDebug) {
		gui.DebugWindow()
	}
	gui.ShowDebugValues()
	*/

	// forever monitor for network and dns changes
	sleep(me.artificialSleep)
	checkNetworkChanges()
}

/*
	Poll for changes to the networking settings
*/
func checkNetworkChanges() {
	var lastLocal time.Time = time.Now()
	var lastDNS time.Time = time.Now()
	/*
func timeFunction(f func()) time.Duration {
	startTime := time.Now() // Record the start time
	f()                     // Execute the function
	return time.Since(startTime) // Calculate the elapsed time
}
*/
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
				me.LocalSpeedActual.SetText(s)
			} else {
				// TODO: make windows and macos diagnostics
				debug(LogError, "Windows and MacOS don't work yet")
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
	debug(LogInfo, "dnsTTL() execution Time: ", duration)
	var s, newSpeed string
	if (duration > 5000 * time.Millisecond ) {
		newSpeed = "VERY BAD"
		suggestProcDebugging()
	} else if (duration > 2000 * time.Millisecond ) {
		newSpeed = "BAD"
		suggestProcDebugging()
	} else if (duration > 500 * time.Millisecond ) {
		suggestProcDebugging()
		newSpeed = "SLOW"
	} else if (duration > 100 * time.Millisecond ) {
		newSpeed = "OK"
		if (me.fixProc != nil) {
			me.fixProc.Disable()
		}
	} else {
		newSpeed = "FAST"
		if (me.fixProc != nil) {
			me.fixProc.Disable()
		}
	}
	if (newSpeed != me.DnsSpeedLast) {
		debug(LogChange, "dns lookup speed changed =", newSpeed)
		debug(LogChange, "dnsTTL() execution Time: ", duration)
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

func linuxLoop() {
	me.changed = false
	debug(LogNet, "FQDN =", me.fqdn.GetText())
	duration := timeFunction(getHostname)
	debug(LogInfo, "getHostname() execution Time: ", duration, "me.changed =", me.changed)

	duration = timeFunction(scanInterfaces)
	debug(LogNet, "scanInterfaces() execution Time: ", duration)
	for i, t := range me.ifmap {
		debug(LogNet, strconv.Itoa(i) + " iface = " + t.iface.Name)
	}

	var aaaa []string
	aaaa = realAAAA()
	var all string
	for _, s := range aaaa {
		debug(LogNet, "my actual AAAA = ",s)
		all += s + "\n"
	}
	// me.IPv6.SetText(all)

	if (me.changed) {
		stamp := time.Now().Format("2006/01/02 15:04:05")
		debug(LogChange, "Network things changed on", stamp)
		duration := timeFunction(updateDNS)
		debug(LogChange, "updateDNS() execution Time: ", duration)
	}

	/*
	processName := getProcessNameByPort(53)
	fmt.Println("Process with port 53:", processName)

	commPath := filepath.Join("/proc", proc.Name(), "comm")
	comm, err := ioutil.ReadFile(commPath)
	if err != nil {
		return "", err // Error reading the process name
	}
	return strings.TrimSpace(string(comm)), nil
	*/
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
