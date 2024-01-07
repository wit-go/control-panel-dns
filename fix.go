// This creates a simple hello world window
package main

import 	(
	"go.wit.com/log"
)

func fix() bool {
	log.Warn("")
	if ! me.statusDNS.Ready() {
		log.Warn("The IPv6 Control Panel is not Ready() yet")
		return false
	}
	if me.statusOS.ValidHostname() {
		log.Warn("Your hostname is VALID:", me.statusOS.GetHostname())
	} else {
		log.Warn("You must first fix your hostname:", me.statusOS.GetHostname())
		return false
	}
	if me.digStatus.IPv4() {
		log.Warn("IPv4 addresses are resolving")
	} else {
		log.Warn("You must first figure out why you can't look up IPv4 addresses")
		log.Warn("Are you on the internet at all?")
		return false
	}
	if me.digStatus.IPv6() {
		log.Warn("IPv6 addresses are resolving")
	} else {
		log.Warn("You must first figure out why you can't look up IPv6 addresses")
		return false
	}
	if ! me.statusDNS.IPv4() {
		log.Warn("You do not have real IPv4 addresses. Nothing to fix here") 
	}
	if ! me.statusDNS.IPv6() {
		log.Warn("IPv6 DNS is broken. Check what is broken here")
		fixIPv6dns()
		return false
	}
	log.Warn("YOU SHOULD BE IN IPv6 BLISS")
	return true
}

func fixIPv6dns() {
	log.Warn("What are my IPv6 addresses?")
	osAAAA := make(map[string]string)
	dnsAAAA := make(map[string]string)

	for _, aaaa := range me.statusOS.GetIPv6() {
		log.Warn("FOUND OS  AAAA ip", aaaa)
		osAAAA[aaaa] = "os"
	}

	log.Warn("What are the AAAA resource records in DNS?")
	for _, aaaa := range me.statusDNS.GetIPv6() {
		log.Warn("FOUND DNS AAAA ip", aaaa)
		dnsAAAA[aaaa] = "dns"
	}

	// remove old DNS entries first
	for aaaa, _ := range dnsAAAA {
		if osAAAA[aaaa] == "dns" {
			log.Warn("DNS AAAA is not in OS", aaaa)
			if deleteFromDNS(aaaa) {
				log.Warn("Delete AAAA", aaaa, "Worked")
			} else {
				log.Warn("Delete AAAA", aaaa, "Failed")
			}
		} else {
			log.Warn("DNS AAAA is in     OS", aaaa)
		}
	}

	// now add new DNS entries
	for aaaa, _ := range osAAAA {
		if dnsAAAA[aaaa] == "dns" {
			log.Warn("OS  AAAA is in     DNS", aaaa)
		} else {
			log.Warn("OS  AAAA is not in DNS", aaaa)
			if addToDNS(aaaa) {
				log.Warn("Add AAAA", aaaa, "Worked")
			} else {
				log.Warn("Add AAAA", aaaa, "Failed")
			}
		}
	}
}

func deleteFromDNS(aaaa string) bool {
	log.Warn("deleteFromDNS", aaaa)
	return false
}

func addToDNS(aaaa string) bool {
	log.Warn("TODO: Add this to DNS !!!!", aaaa)
	log.Warn("what is your API provider?")
	return false
}

func exists(m map[string]bool, s string) bool {
	if _, ok := m[s]; ok {
		return true
	}
	return false
}
