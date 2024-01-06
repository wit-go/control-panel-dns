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
func (h *Host) checkDNS() {
	var ip4 bool = false
	var ip6 bool = false

	for s, t := range h.ipmap {
		i := t.iface
		ipt := "IPv4"
		if (t.ipv6) {
			ipt = "IPv6"
		}
		if (! t.IsReal()) {
			log.Println(args.VerboseDNS, "\tIP is not Real", ipt, i.Index, i.Name, s)
			continue
		}

		log.Println(args.VerboseDNS, "\tIP is Real    ", ipt, i.Index, i.Name, s)
		if (t.ipv6) {
			ip6 = true
		} else {
			ip4 = true
		}
	}

	if (ip4 == true) {
		log.Println(args.VerboseDNS, "IPv4 should work. Wow. You actually have a real IPv4 address")
	} else {
		log.Println(args.VerboseDNS, "IPv4 is broken. (be nice and setup ipv4-only.wit.com)")
	}
	if (ip6 == true) {
		log.Println(args.VerboseDNS, "IPv6 should be working. Need to test it here.")
	} else {
		log.Println(args.VerboseDNS, "IPv6 is broken. Need to fix it here.")
	}
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

	if (tmp != me.status.NSrr.Get()) {
		me.changed = true
		log.Log(CHANGE, "lookupNS() setting changed to me.NSrr =", tmp)
		me.status.NSrr.Set(tmp)
	}
}

// getDomain returns the second-to-last part of a domain name.
func setProvider(hostname string) {
	var provider string = ""
	parts := strings.Split(hostname, ".")
	if len(parts) >= 2 {
		provider = parts[len(parts)-2]
	}
	if (me.DnsAPI.S != provider) {
		me.changed = true
		log.Log(CHANGE, "setProvider() changed to =", provider)
		me.DnsAPI.SetText(provider)
	}
}
