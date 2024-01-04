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
	if (me.fqdn != nil) {
		if (me.hostname != s) {
			me.fqdn.SetText(s)
			me.hostname = s
			me.changed = true
		}
	}
	log.Log(NET, "FQDN =", s)

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
	if (me.hostname != test) {
		log.Info("me.hostname", me.hostname, "does not equal", test)
		if (me.hostnameStatusOLD.S != "BROKEN") {
			log.Log(CHANGE, "me.hostname", me.hostname, "does not equal", test)
			me.changed = true
			me.hostnameStatusOLD.SetText("BROKEN")
		}
	} else {
		if (me.hostnameStatusOLD.S != "VALID") {
			log.Log(CHANGE, "me.hostname", me.hostname, "is valid")
			me.hostnameStatusOLD.SetText("VALID")
			me.changed = true
		}
		// enable the cloudflare button if the provider is cloudflare
		if (me.cloudflareB == nil) {
			log.Log(CHANGE, "me.cloudflare == nil; me.DnsAPI.S =", me.DnsAPI.S)
			if (me.DnsAPI.S == "cloudflare") {
				me.cloudflareB = me.mainStatus.NewButton("cloudflare wit.com", func () {
					cloudflare.CreateRR(myGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
				})
			}
		}
	}
}

// returns true if the hostname is good
// check that all the OS settings are correct here
// On Linux, /etc/hosts, /etc/hostname
//      and domainname and hostname
func goodHostname(h string) bool {
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

/*
func digAAAA(s string) []string {
	var aaaa []string
	// lookup the IP address from DNS
	rrset := dnssecsocket.Dnstrace(s, "AAAA")
	// log.Spew(args.VerboseDNS, SPEW, rrset)
	for i, rr := range rrset {
		ipaddr := dns.Field(rr, 1)
		// how the hell do you detect a RRSIG	AAAA record here?
		if (ipaddr == "28") {
			continue
		}
		log.Log(NOW, "r.Answer =", i, "rr =", rr, "ipaddr =", ipaddr)
		aaaa = append(aaaa, ipaddr)
		me.ipv6s[ipaddr] = rr
	}
	log.Info(args.VerboseDNS, "aaaa =", aaaa)
	log.Println("digAAAA() returned =", aaaa)
	log.Println("digAAAA() me.ipv6s =", me.ipv6s)
	os.Exit(0)
	return aaaa
}
*/

func digAAAA(hostname string) []string {
	var blah, ipv6Addresses []string
	// domain := hostname
	recordType := dns.TypeAAAA // dns.TypeTXT

	// Cloudflare's DNS server
	blah, _ = dnsUdpLookup("1.1.1.1:53", hostname, recordType)
	log.Println("digAAAA() has BLAH =", blah)

	if (len(blah) == 0) {
		log.Println("digAAAA() RUNNING dnsAAAAlookupDoH(domain)")
		ipv6Addresses, _ = dnsAAAAlookupDoH(hostname)
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

/*
func dnsHttpsLookup(domain string, recordType uint16) ([]string, error) {
	domain := "google.com"
	dnsLookupDoH(domain string) ([]string, error) {
	ipv6Addresses, err := dnsLookupDoH(domain)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Printf("IPv6 Addresses for %s:\n", domain)
	for _, addr := range ipv6Addresses {
		log.Println(addr)
	}
}
*/
