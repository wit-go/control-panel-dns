// figures out if your hostname is valid
// then checks if your DNS is setup correctly
package main

import (
	"strings"

	"go.wit.com/log"
	"go.wit.com/shell"
	"go.wit.com/gui/cloudflare"

	"github.com/miekg/dns"
	// will try to get this hosts FQDN
	"github.com/Showmax/go-fqdn"
)

func getHostname() {
	var err error
	var s string = "gui.Label == nil"
	s, err = fqdn.FqdnHostname()
	if (err != nil) {
		log.Error(err, "FQDN hostname error")
		return
	}
	me.status.SetHostname(s)

	dn := run("domainname")
	if (me.domainname.S != dn) {
		log.Log(CHANGE, "domainname has changed from", me.domainname.S, "to", dn)
		me.domainname.SetText(dn)
		me.changed = true
	}

	hshort := run("hostname -s")
	if (me.hostshort.S != hshort) {
		log.Log(CHANGE, "hostname -s has changed from", me.hostshort.S, "to", hshort)
		me.hostshort.SetText(hshort)
		me.changed = true
	}

	var test string
	test = hshort + "." + dn
	if (me.status.GetHostname() != test) {
		log.Log(CHANGE, "me.hostname", me.status.GetHostname(), "does not equal", test)
		if (me.hostnameStatus.S != "BROKEN") {
			log.Log(CHANGE, "me.hostname", me.status.GetHostname(), "does not equal", test)
			me.changed = true
			me.hostnameStatus.SetText("BROKEN")
		}
	} else {
		if (me.hostnameStatus.S != "VALID") {
			log.Log(CHANGE, "me.hostname", me.status.GetHostname(), "is valid")
			me.hostnameStatus.SetText("VALID")
			me.changed = true
		}
		// enable the cloudflare button if the provider is cloudflare
		if (me.cloudflareB == nil) {
			log.Log(CHANGE, "me.cloudflare == nil; me.DnsAPI.S =", me.DnsAPI.S)
			if (me.DnsAPI.S == "cloudflare") {
				me.cloudflareB = me.mainStatus.NewButton("cloudflare wit.com", func () {
					cloudflare.CreateRR(me.myGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
				})
			}
		}
	}
}

// returns true if the hostname is good
// check that all the OS settings are correct here
// On Linux, /etc/hosts, /etc/hostname
//      and domainname and hostname
func goodHostname() bool {
	hostname := shell.Chomp(shell.Cat("/etc/hostname"))
	log.Log(NOW, "hostname =", hostname)

	hs := run("hostname -s")
	dn := run("domainname")
	log.Log(NOW, "hostname short =", hs, "domainname =", dn)

	tmp := hs + "." + dn
	if (hostname == tmp) {
		log.Log(NOW, "hostname seems to be good", hostname)
		return true
	}

	return false
}

func digAAAA(hostname string) []string {
	var blah, ipv6Addresses []string
	// domain := hostname
	recordType := dns.TypeAAAA // dns.TypeTXT

	// Cloudflare's DNS server
	blah, _ = dnsUdpLookup("1.1.1.1:53", hostname, recordType)
	log.Println("digAAAA() has BLAH =", blah)

	if (len(blah) == 0) {
		log.Println("digAAAA() RUNNING dnsAAAAlookupDoH(domain)")
		ipv6Addresses = lookupDoH(hostname, "AAAA")
		log.Println("digAAAA() has ipv6Addresses =", strings.Join(ipv6Addresses, " "))
		for _, addr := range ipv6Addresses {
			log.Println(addr)
		}
		return ipv6Addresses
	}

	// TODO: check digDoH vs blah, if so, then port 53 TCP and/or UDP is broken or blocked
	log.Println("digAAAA() has BLAH =", blah)

	return blah
}
