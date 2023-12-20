// inspired from:
// https://github.com/mactsouk/opensource.com.git
// and
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

package main

import (
	"git.wit.org/wit/shell"

	// dnssec IPv6 socket library
	"git.wit.org/jcarr/dnssecsocket"

	"git.wit.org/jcarr/control-panel-dns/cloudflare"
)

// will try to get this hosts FQDN
import "github.com/Showmax/go-fqdn"

// this is the king of dns libraries
import "github.com/miekg/dns"


func getHostname() {
	var err error
	var s string = "gui.Label == nil"
	s, err = fqdn.FqdnHostname()
	if (err != nil) {
		debug(LogError, "FQDN hostname error =", err)
		return
	}
	if (me.fqdn != nil) {
		if (me.hostname != s) {
			me.fqdn.SetText(s)
			me.hostname = s
			me.changed = true
		}
	}
	debug(LogNet, "FQDN =", s)

	dn := run("domainname")
	if (me.domainname.S != dn) {
		debug(LogChange, "domainname has changed from", me.domainname.S, "to", dn)
		me.domainname.SetText(dn)
		me.changed = true
	}

	hshort := run("hostname -s")
	if (me.hostshort.S != hshort) {
		debug(LogChange, "hostname -s has changed from", me.hostshort.S, "to", hshort)
		me.hostshort.SetText(hshort)
		me.changed = true
	}

	var test string
	test = hshort + "." + dn
	if (me.hostname != test) {
		debug(LogInfo, "me.hostname", me.hostname, "does not equal", test)
		if (me.hostnameStatus.S != "BROKEN") {
			debug(LogChange, "me.hostname", me.hostname, "does not equal", test)
			me.changed = true
			me.hostnameStatus.SetText("BROKEN")
		}
	} else {
		if (me.hostnameStatus.S != "VALID") {
			debug(LogChange, "me.hostname", me.hostname, "is valid")
			me.hostnameStatus.SetText("VALID")
			me.changed = true
		}
		// enable the cloudflare button if the provider is cloudflare
		if (me.cloudflareB == nil) {
			debug(LogChange, "me.cloudflare == nil; me.DnsAPI.S =", me.DnsAPI.S)
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
	debug(true, "hostname =", hostname)

	hs := run("hostname -s")
	dn := run("domainname")
	debug(true, "hostname short =", hs, "domainname =", dn)

	tmp := hs + "." + dn
	if (hostname == tmp) {
		debug(true, "hostname seems to be good", hostname)
		return true
	}

	return false
}

func digAAAA(s string) []string {
	var aaaa []string
	// lookup the IP address from DNS
	rrset := dnssecsocket.Dnstrace(s, "AAAA")
	// debug(true, args.VerboseDNS, SPEW, rrset)
	for i, rr := range rrset {
		ipaddr := dns.Field(rr, 1)
		// how the hell do you detect a RRSIG	AAAA record here?
		if (ipaddr == "28") {
			continue
		}
		debug(LogNow, "r.Answer =", i, "rr =", rr, "ipaddr =", ipaddr)
		aaaa = append(aaaa, ipaddr)
		me.ipv6s[ipaddr] = rr
	}
	debug(true, args.VerboseDNS, "aaaa =", aaaa)
	return aaaa
}
