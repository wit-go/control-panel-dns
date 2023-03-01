// inspired from:
// https://github.com/mactsouk/opensource.com.git
// and
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

package main

// import "net"

// will try to get this hosts FQDN
import "github.com/Showmax/go-fqdn"

// this is the king of dns libraries
import "github.com/miekg/dns"

// dnssec IPv6 socket library
import "git.wit.org/jcarr/dnssecsocket"

func getHostname() {
	var err error
	me.fqdn, err = fqdn.FqdnHostname()
	if (err != nil) {
		log("FQDN hostname error =", err)
		exit()
		return
	}
	log("FQDN hostname is", me.fqdn)
}

func dnsAAAA(s string) []string {
	var aaaa []string
	// lookup the IP address from DNS
	rrset := dnssecsocket.Dnstrace(s, "AAAA")
	log(args.VerboseDNS, SPEW, rrset)
	for i, rr := range rrset {
		log(args.VerboseDNS, "r.Answer =", i, rr)
		ipaddr := dns.Field(rr, 1)
		aaaa = append(aaaa, ipaddr)
	}
	log(args.VerboseDNS, "aaaa =", aaaa)
	return aaaa
}