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

	var aaaa []string
	aaaa = getAAAA(me.fqdn)
	log("AAAA =", aaaa)
}

func getAAAA(s string) []string {
	// lookup the IP address from DNS
	dnsRR := dnssecsocket.Dnstrace(s, "AAAA")
	log(args.VerboseDNS, SPEW, dnsRR)
	if (dnsRR == nil) {
		return nil
	}
	ipaddr1 := dns.Field(dnsRR, 1)
	ipaddr2 := dns.Field(dnsRR, 2)
	log("ipaddr", ipaddr1, ipaddr2)
	return []string{ipaddr1, ipaddr2}
}
