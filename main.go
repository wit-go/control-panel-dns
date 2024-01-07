// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package main

import 	(
	"fmt"
	// "runtime"
	"time"
	"embed"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/debugger"

	"go.wit.com/wit/control-panel-dns/linuxstatus"

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

	me.statusOS = linuxstatus.New()
	me.statusOS.SetParent(me.myGui)
	me.statusOS.InitWindow()
	me.statusOS.Make()
	me.statusOS.Draw()
	me.statusOS.Draw2()

	if debugger.ArgDebug() {
		go func() {
			log.Sleep(2)
			debugger.DebugWindow(me.myGui)
		}()
	}

	log.Sleep(me.artificialSleep)

	// TCP & UDP port 53 lookups + DNS over HTTP lookups + os.Exec(dig)
	go myTicker(60 * time.Second, "DNSloop", func() {
		me.digStatus.Update()

		if me.digStatus.Ready() {
			current := me.statusIPv6.Get()
			if me.digStatus.IPv6() {
				if current != "WORKING" {
					log.Log(CHANGE, "IPv6 resolution is WORKING")
					me.statusIPv6.Set("WORKING")
				}
			} else {
				if current != "Need VPN" {
					log.Log(CHANGE, "IPv6 resolution seems to have broken")
					me.statusIPv6.Set("Need VPN")
				}
			}
		}
	})

	// checks if your DNS records are still broken
	// if everything is working, then it just ignores
	// things until the timeout happens

	lastProvider := "unknown"
	go myTicker(10 * time.Second, "DNSloop", func() {
		log.Log(INFO, "me.statusDNS.Update() START")
		me.statusDNS.Update()

		provider := me.statusDNS.GetDNSapi()
		if provider != lastProvider {
			log.Log(CHANGE, "Your DNS API provider appears to have changed to", provider)
			lastProvider = provider
			me.apiButton.SetText(provider + " wit.com")
		}
		if provider == "cloudflare" {
			me.DnsAPIstatus.Set("WORKING")
		}
	})

	// probes the OS network settings
	go myTicker(500 * time.Millisecond, "me.statusOS,Update()", func() {
		duration := timeFunction( func() {
			me.statusOS.Update()

			if me.statusOS.ValidHostname() {
				if me.hostnameStatus.GetText() != "WORKING" {
					me.hostnameStatus.Set("WORKING")
					me.changed = true
				}
			}

			// re-check DNS API provider
			if (me.statusOS.Changed()) {
				// lookup the NS records for your domain
				// if your host is test.wit.com, find the NS resource records for wit.com
				lookupNS(me.statusOS.GetDomainName())
				log.Log(CHANGE, "updateDNS() END")
			}
		})
		s := fmt.Sprint(duration)
		me.statusOS.SetSpeedActual(s)
	})

	// check the four known things to see if they are all WORKING
	myTicker(10 * time.Second, "MAIN LOOP", func() {
		if me.hostnameStatus.GetText() != "WORKING" {
			log.Log(CHANGE, "The hostname is not WORKING yet", me.hostnameStatus.GetText())
			return
		}
		if me.statusIPv6.Get() != "WORKING" {
			log.Log(CHANGE, "IPv6 DNS lookup has not been confirmed yet", me.statusIPv6.Get())
			return
		}
		if me.DnsStatus.GetText() != "WORKING" {
			log.Log(CHANGE, "Your IPv6 DNS settings have not been confirmed yet", me.DnsStatus.GetText()) 
			return
		}
		if me.DnsAPIstatus.GetText() != "WORKING" {
			log.Log(CHANGE, "The DNS API provider is not yet working", me.DnsAPIstatus.GetText())
			return
		}
		log.Log(CHANGE, "EVERYTHING IS WORKING. YOU HAVE IPv6 BLISS. TODO: don't check so often now")
	})
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

func timeStamp() string {
	stamp := time.Now().Format("2006/01/02 15:04:05")
	log.Log(CHANGE, "Network things changed on", stamp)
	return stamp
}


func myTicker(t time.Duration, name string, f func()) {
	ticker := time.NewTicker(t)
	defer ticker.Stop()
	done := make(chan bool)
	/*
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	*/
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			log.Log(INFO, name, "Current time: ", t)
			f()
		}
	}
}
