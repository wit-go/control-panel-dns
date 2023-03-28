// Various Linux/Unix'y things

// https://wiki.archlinux.org/title/Dynamic_DNS

package main

import 	(
	"net"
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
                log(logError, "updateIPs failed", err)
        }
        for _, ip := range ips {
                log(host, ip)
        }
}

func (h *Host) setIPv4(ipv4s map[string]*IPtype) {
        for ip, t := range ipv4s {
		log("IPv4", ip, t)
        }
}

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
			log(args.VerboseDNS, "\tIP is not Real", ipt, i.Index, i.Name, s)
			continue
		}

		log(args.VerboseDNS, "\tIP is Real    ", ipt, i.Index, i.Name, s)
		if (t.ipv6) {
			ip6 = true
		} else {
			ip4 = true
		}
	}

	if (ip4 == true) {
		log(args.VerboseDNS, "IPv4 should work. Wow. You actually have a real IPv4 address")
	} else {
		log(args.VerboseDNS, "IPv4 is broken. (be nice and setup ipv4-only.wit.com)")
	}
	if (ip6 == true) {
		log(args.VerboseDNS, "IPv6 should be working. Need to test it here.")
	} else {
		log(args.VerboseDNS, "IPv6 is broken. Need to fix it here.")
	}
}
