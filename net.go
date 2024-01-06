// This creates a simple hello world window
package main

import 	(
	// "log"
	"net"
	"strings"

	"go.wit.com/log"
)

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func (t *IPtype) IsReal() bool {
	if (t.ip.IsPrivate() || t.ip.IsLoopback() || t.ip.IsLinkLocalUnicast()) {
		log.Log(NET, "\t\tIP is Real = false")
		return false
	} else {
		log.Log(NET, "\t\tIP is Real = true")
		return true
	}
}

func IsReal(ip *net.IP) bool {
	if (ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast()) {
		log.Log(NET, "\t\tIP is Real = false")
		return false
	} else {
		log.Log(NET, "\t\tIP is Real = true")
		return true
	}
}

/*
	These are the real IP address you have been
	given from DHCP
*/
func dhcpAAAA() []string {
	var aaaa []string

	for s, t := range me.ipmap {
		if (t.IsReal()) {
			if (t.ipv6) {
				aaaa = append(aaaa, s)
			}
		}
	}
	return aaaa
}

func realA() []string {
	var a []string

	for s, t := range me.ipmap {
		if (t.IsReal()) {
			if (t.ipv4) {
				a = append(a, s)
			}
		}
	}
	return a
}

func checkDNS() (map[string]*IPtype, map[string]*IPtype) {
	var ipv4s map[string]*IPtype
	var ipv6s map[string]*IPtype

	ipv4s = make(map[string]*IPtype)
	ipv6s = make(map[string]*IPtype)

	for s, t := range me.ipmap {
		i := t.iface
		ipt := "IPv4"
		if (t.ipv6) {
			ipt = "IPv6"
		}
		if (t.IsReal()) {
			log.Info("\tIP is Real    ", ipt, i.Index, i.Name, s)
			if (t.ipv6) {
				ipv6s[s] = t
			} else {
				ipv4s[s] = t
			}
		} else {
			log.Info("\tIP is not Real", ipt, i.Index, i.Name, s)
		}
	}
	return ipv6s, ipv4s
}
