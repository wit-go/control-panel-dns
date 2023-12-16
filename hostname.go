// inspired from:
// https://github.com/mactsouk/opensource.com.git
// and
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

package main

import (
	"log"
	"git.wit.org/wit/shell"
)

// will try to get this hosts FQDN
import "github.com/Showmax/go-fqdn"

// this is the king of dns libraries
import "github.com/miekg/dns"

// dnssec IPv6 socket library
import "git.wit.org/jcarr/dnssecsocket"

func getHostname() {
	var err error
	var s string = "gui.Label == nil"
	s, err = fqdn.FqdnHostname()
	if (err != nil) {
		log.Println("FQDN hostname error =", err)
		return
	}
	if (me.fqdn != nil) {
		if (me.hostname != s) {
			me.fqdn.SetText(s)
			me.hostname = s
			me.changed = true
		}
	}
	log.Println("FQDN =", s)
}

// returns true if the hostname is good
// check that all the OS settings are correct here
// On Linux, /etc/hosts, /etc/hostname
//      and domainname and hostname
func goodHostname(h string) bool {
	hostname := shell.Chomp(shell.Cat("/etc/hostname"))
	log.Println("hostname =", hostname)

	hs := run("hostname -s")
	dn := run("domainname")
	log.Println("hostname short =", hs, "domainname =", dn)

	tmp := hs + "." + dn
	if (hostname == tmp) {
		log.Println("hostname seems to be good", hostname)
		return true
	}

	return false
}

func dnsAAAA(s string) []string {
	var aaaa []string
	// lookup the IP address from DNS
	rrset := dnssecsocket.Dnstrace(s, "AAAA")
	log.Println(args.VerboseDNS, SPEW, rrset)
	for i, rr := range rrset {
		log.Println(args.VerboseDNS, "r.Answer =", i, rr)
		ipaddr := dns.Field(rr, 1)
		aaaa = append(aaaa, ipaddr)
	}
	log.Println(args.VerboseDNS, "aaaa =", aaaa)
	return aaaa
}
