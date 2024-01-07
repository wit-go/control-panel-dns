// This creates a simple hello world window
package main

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/cloudflare"
	"go.wit.com/wit/control-panel-dns/smartwindow"
)

func fix() bool {
	// make and toggle the fixWindow display
	if me.fixWindow == nil {
		me.fixWindow = smartwindow.New()
		me.fixWindow.SetParent(me.myGui)
		me.fixWindow.Title("fix window")
		me.fixWindow.SetDraw(drawFixWindow)
		me.fixWindow.Vertical()
		me.fixWindow.Make()
		me.fixWindow.Draw()
		me.fixWindow.Hide()
		// me.fixWindow.Draw2()
		return false
	}
	me.fixWindow.Toggle()

	if ! me.statusDNS.Ready() {
		log.Log(CHANGE, "The IPv6 Control Panel is not Ready() yet")
		return false
	}
	if me.statusOS.ValidHostname() {
		log.Log(CHANGE, "GOOD Your hostname is VALID:", me.statusOS.GetHostname())
	} else {
		log.Log(CHANGE, "You must first fix your hostname:", me.statusOS.GetHostname())
		return false
	}
	if me.digStatus.IPv4() {
		log.Log(CHANGE, "GOOD IPv4 addresses are resolving")
	} else {
		log.Log(CHANGE, "You must first figure out why you can't look up IPv4 addresses")
		log.Log(CHANGE, "Are you on the internet at all?")
		return false
	}
	if me.digStatus.IPv6() {
		log.Log(CHANGE, "GOOD IPv6 addresses are resolving")
	} else {
		log.Log(CHANGE, "You must first figure out why you can't look up IPv6 addresses")
		return false
	}
	if ! me.statusDNS.IPv4() {
		log.Log(CHANGE, "OK   You do not have real IPv4 addresses. Nothing to fix here") 
	}
	if ! me.statusDNS.IPv6() {
		if fixIPv6dns() {
			log.Log(CHANGE, "IPv6 DNS Repair is underway")
			return false
		}
		log.Log(CHANGE, "GOOD IPv6 DNS is working!")
	}
	log.Log(CHANGE, "GOOD YOU SHOULD BE IN IPv6 BLISS")
	return true
}

func fixIPv6dns() bool {
	log.Log(INFO, "What are my IPv6 addresses?")
	var broken bool = false
	osAAAA := make(map[string]string)
	dnsAAAA := make(map[string]string)

	for _, aaaa := range me.statusOS.GetIPv6() {
		log.Log(INFO, "FOUND OS  AAAA ip", aaaa)
		osAAAA[aaaa] = "os"
	}

	log.Log(INFO, "What are the AAAA resource records in DNS?")
	for _, aaaa := range me.statusDNS.GetIPv6() {
		log.Log(INFO, "FOUND DNS AAAA ip", aaaa)
		dnsAAAA[aaaa] = "dns"
	}

	// remove old DNS entries first
	for aaaa, _ := range dnsAAAA {
		if osAAAA[aaaa] == "os" {
			log.Log(INFO, "DNS AAAA is in     OS", aaaa)
		} else {
			broken = true
			log.Log(INFO, "DNS AAAA is not in OS", aaaa)
			addToFixWindow("DELETE", aaaa)
			/*
			if deleteFromDNS(aaaa) {
				log.Log(INFO, "Delete AAAA", aaaa, "Worked")
			} else {
				log.Log(INFO, "Delete AAAA", aaaa, "Failed")
			}
			*/
		}
	}

	// now add new DNS entries
	for aaaa, _ := range osAAAA {
		if dnsAAAA[aaaa] == "dns" {
			log.Log(INFO, "OS  AAAA is in     DNS", aaaa)
		} else {
			broken = true
			log.Log(INFO, "OS  AAAA is not in DNS", aaaa)
			addToFixWindow("CREATE", aaaa)
			/*
			if addToDNS(aaaa) {
				log.Log(INFO, "Add AAAA", aaaa, "Worked")
			} else {
				log.Log(INFO, "Add AAAA", aaaa, "Failed")
			}
			*/
		}
	}

	// if anything doesn't match, return false
	return broken
}

func deleteFromDNS(aaaa string) bool {
	log.Log(CHANGE, "Delete this from DNS !!!!", aaaa)
	api := me.statusDNS.API()
	log.Log(CHANGE, "your API provider is =", api)
	if api == "cloudflare" {
		log.Log(CHANGE, "Let's try a DELETE via the Cloudflare API")
		hostname := me.statusOS.GetHostname()
		b, response := cloudflare.Delete("wit.com", hostname, aaaa)
		log.Log(CHANGE, "response was:", response)
		return b
	}
	return false
}

func addToDNS(aaaa string) bool {
	log.Log(CHANGE, "Add this to DNS !!!!", aaaa)
	api := me.statusDNS.API()
	log.Log(CHANGE, "your API provider is =", api)
	if api == "cloudflare" {
		log.Log(CHANGE, "Let's try a CREATE via the Cloudflare API")
		hostname := me.statusOS.GetHostname()
		return cloudflare.Create("wit.com", hostname, aaaa)
	}
	return false
}

func exists(m map[string]bool, s string) bool {
	if _, ok := m[s]; ok {
		return true
	}
	return false
}

var myErrorBox *errorBox

func addToFixWindow(t string, ip string) {
	log.Log(INFO, "addToFixWindow() START")
	if me.fixWindow == nil {
		log.Log(WARN, "addToFixWindow() fixWindow == nil. Can't add the error", t, ip)
		return
	}
	if myErrorBox == nil {
		box := me.fixWindow.Box()
		myErrorBox = NewErrorBox(box, t, ip)
	}
	myErrorBox.add(t, ip)
	log.Log(INFO, "addToFixWindow() END")
}

func drawFixWindow(sw *smartwindow.SmartWindow) {
	log.Log(WARN, "drawFixWindow() START")
	box := sw.Box()
	box.NewLabel("test")
}
