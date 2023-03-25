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
	var s string = "gui.Label == nil"
	s, err = fqdn.FqdnHostname()
	if (err != nil) {
		log("FQDN hostname error =", err)
		exit()
		return
	}
	if (me.fqdn != nil) {
		// s =  me.fqdn.GetText()
		log("trying to update gui.Label")
		// me.fqdn.AddText(s)
		me.fqdn.SetText(s)
		me.hostname = s
	}
	log("FQDN =", s)
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
