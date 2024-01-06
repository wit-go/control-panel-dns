// This creates a simple hello world window
package main

import 	(
	"go.wit.com/log"
)

func fix() bool {
	log.Warn("")
	if ! me.status.Ready() {
		log.Warn("The IPv6 Control Panel is not Ready() yet")
		return false
	}
	if me.status.ValidHostname() {
		log.Warn("Your hostname is VALID:", me.status.GetHostname())
	} else {
		log.Warn("You must first fix your hostname:", me.status.GetHostname())
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
	if ! me.status.IPv4() {
		log.Warn("You do not have real IPv4 addresses. Nothing to fix here") 
	}
	if ! me.status.IPv6() {
		log.Warn("IPv6 DNS is broken. Check what is broken here")
		log.Warn("What are my IPv6 addresses?")
		log.Warn("What are the AAAA resource records in DNS?")
		return false
	}
	log.Warn("YOU SHOULD BE IN IPv6 BLISS")
	return true
}
