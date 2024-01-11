// Various Linux/Unix'y things

// https://wiki.archlinux.org/title/Dynamic_DNS

package main

import 	(
	"net"
	"strings"

	"go.wit.com/log"
	"go.wit.com/shell"
)

/*
	Check a bunch of things. If they don't work right, then things are not correctly configured
	They are things like:
		/etc/hosts
		hostname
		hostname -f
		domainname
*/
func (h *Host) verifyETC() bool {
	return true
}

func (h *Host) updateIPs(host string) {
    ips, err := net.LookupIP(host)
        if err != nil {
                log.Error(err, "updateIPs failed")
        }
        for _, ip := range ips {
                log.Println(host, ip)
        }
}

func (h *Host) setIPv4(ipv4s map[string]*IPtype) {
        for ip, t := range ipv4s {
		log.Println("IPv4", ip, t)
        }
}

/*
func lookupNSprovider(domain string) string {
	for s, d := range me.nsmap {
		log.Log(CHANGE, "lookupNS() domain =", d, "server =", s)
		if (domain == d) {
			// figure out the provider (google, cloudflare, etc)
			return s + " blah"
		}
	}
	return "blah"
}
*/

// nsLookup performs an NS lookup on the given domain name.
func lookupNS(domain string) {
	var domains string

	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		return
	}

	var servers []string
	for _, ns := range nsRecords {
		servers = append(servers, ns.Host)
	}

	// checks to see if the NS records change
	for _, server := range servers {
		server = strings.TrimRight(server, ".")
		if (me.nsmap[server] != domain) {
			log.Log(CHANGE, "lookupNS() domain", domain, "has NS", server)
			me.nsmap[server] = domain
			domains += server + "\n"
		}
	}

	var tmp string
	// checks to see if the NS records change
	for s, d := range me.nsmap {
		log.Log(CHANGE, "lookupNS() domain =", d, "server =", s)
		if (domain == d) {
			tmp += s + "\n"
			// figure out the provider (google, cloudflare, etc)
			setProvider(s)
		}
	}
	tmp = shell.Chomp(tmp)

	if (tmp != me.statusDNS.NSrr.Get()) {
		me.changed = true
		log.Log(CHANGE, "lookupNS() setting changed to me.NSrr =", tmp)
		me.statusDNS.NSrr.Set(tmp)
	}
}

// returns the second-to-last part of a domain name.
func setProvider(hostname string) {
	var provider string = ""
	parts := strings.Split(hostname, ".")
	if len(parts) >= 2 {
		provider = parts[len(parts)-2]
	}
	if me.statusDNS.GetDNSapi() != provider {
		log.Log(CHANGE, "setProvider() changed to =", provider)
	}
	me.statusDNS.SetDNSapi(provider)
}
