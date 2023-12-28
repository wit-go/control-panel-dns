// This creates a simple hello world window
package main

import 	(
	// "log"
	"net"
	"strings"
)

// this doesn't work
/*
func watchNetworkInterfaces() {
	// Get list of network interfaces
	interfaces, _ := net.Interfaces()

	// Set up a notification channel
	notification := make(chan net.Interface)

	debug(LogNet, "watchNet()")
	// Start goroutine to watch for changes
	go func() {
		debug(LogNet, "watchNet() func")
		for {
			debug(LogNet, "forever loop start")
			// Check for changes in each interface
			for _, i := range interfaces {
				debug(LogNet, "something on i =", i)
				if status := i.Flags & net.FlagUp; status != 0 {
					notification <- i
					debug(LogNet, "something on i =", i)
				}
			}
			debug(LogNet, "forever loop end")
		}
	}()
}
*/

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func (t *IPtype) IsReal() bool {
	if (t.ip.IsPrivate() || t.ip.IsLoopback() || t.ip.IsLinkLocalUnicast()) {
		debug(LogNet, "\t\tIP is Real = false")
		return false
	} else {
		debug(LogNet, "\t\tIP is Real = true")
		return true
	}
}

func IsReal(ip *net.IP) bool {
	if (ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast()) {
		debug(LogNet, "\t\tIP is Real = false")
		return false
	} else {
		debug(LogNet, "\t\tIP is Real = true")
		return true
	}
}

func renameInterface(i *net.Interface) {
	/*
	/sbin/ip link set eth1 down
	/sbin/ip link set eth1 name eth123
	/sbin/ip link set eth123 up
	*/
}

// Will figure out if an interface was just added
func checkInterface(i net.Interface) {
	val, ok := me.ifmap[i.Index]
	if ! ok {
		debug(i.Name, "is a new network interface. The linux kernel index =", i.Index)
		me.ifmap[i.Index] = new(IFtype)
		me.ifmap[i.Index].gone = false
		me.ifmap[i.Index].iface = &i
		me.changed = true
		if (me.Interfaces != nil) {
			me.Interfaces.AddText(i.Name)
			me.Interfaces.SetText(i.Name)
		}
		return
	}
	me.ifmap[i.Index].gone = false
	debug(LogNet, "me.ifmap[i] does exist. Need to compare everything.", i.Index, i.Name, val.iface.Index, val.iface.Name)
	if (val.iface.Name != i.Name) {
		debug(val.iface.Name, "has changed to it's name to", i.Name)
		me.ifmap[i.Index].iface = &i
		me.changed = true
		if (me.Interfaces != nil) {
			me.Interfaces.AddText(i.Name)
			me.Interfaces.SetText(i.Name)
		}
		return
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
			debug("\tIP is Real    ", ipt, i.Index, i.Name, s)
			if (t.ipv6) {
				ipv6s[s] = t
			} else {
				ipv4s[s] = t
			}
		} else {
			debug("\tIP is not Real", ipt, i.Index, i.Name, s)
		}
	}
	return ipv6s, ipv4s
}

// Will figure out if an IP address is new
func checkIP(ip *net.IPNet, i net.Interface) bool {
	debug(LogNet, "\t\taddr.(type) = *net.IPNet")
	debug(LogNet, "\t\taddr.(type) =", ip)
	var realip string
	realip = ip.IP.String()

	val, ok := me.ipmap[realip]
	if ok {
		debug(LogNet, val.ipnet.IP.String(), "is already a defined IP address")
		me.ipmap[realip].gone = false
		return false
	}

	me.ipmap[realip] = new(IPtype)
	me.ipmap[realip].gone = false
	me.ipmap[realip].ipv4 = true
	me.ipmap[realip].ipnet = ip
	me.ipmap[realip].ip = ip.IP
	me.ipmap[realip].iface = &i
	t := "IPv4"
	if (IsIPv6(ip.String())) {
		me.ipmap[realip].ipv6 = true
		me.ipmap[realip].ipv4 = false
		t = "IPv6"
		if (me.IPv6 != nil) {
			me.IPv6.SetText(realip)
		}
	} else {
		me.ipmap[realip].ipv6 = false
		me.ipmap[realip].ipv4 = true
		if (me.IPv4 != nil) {
			me.IPv4.SetText(realip)
		}
	}
	if (IsReal(&ip.IP)) {
		debug("\tIP is Real    ", t, i.Index, i.Name, realip)
	} else {
		debug("\tIP is not Real", t, i.Index, i.Name, realip)
	}
	debug(LogNet, "\t\tIP is IsPrivate() =", ip.IP.IsPrivate())
	debug(LogNet, "\t\tIP is IsLoopback() =", ip.IP.IsLoopback())
	debug(LogNet, "\t\tIP is IsLinkLocalUnicast() =", ip.IP.IsLinkLocalUnicast())
	// debug("HERE HERE", "realip =", realip, "me.ip[realip]=", me.ipmap[realip])
	return true
}

func scanInterfaces() {
	debug(LogNet, "scanInterfaces() START")
	ifaces, _ := net.Interfaces()
	// me.ifnew = ifaces
	debug(LogNet, SPEW, ifaces)
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// debug("range ifaces = ", i)
		checkInterface(i)
		debug(LogNet, "*net.Interface.Name = ", i.Name, i.Index)
		debug(LogNet, SPEW, i)
		debug(LogNet, SPEW, addrs)
		for _, addr := range addrs {
			debug(LogNet, "\taddr =", addr)
			debug(LogNet, SPEW, addrs)
			ips, _ := net.LookupIP(addr.String())
			debug(LogNet, "\tLookupIP(addr) =", ips)
			switch v := addr.(type) {
			case *net.IPNet:
				if checkIP(v, i) {
					debug(true, "scanInterfaces() IP is new () i =", v.IP.String())
				}
			default:
				debug(LogNet, "\t\taddr.(type) = NO IDEA WHAT TO DO HERE v =", v)
			}
		}
	}
	if deleteChanges() {
		me.changed = true
		debug(LogNow, "deleteChanges() detected network changes")
	}
	updateRealAAAA()
	debug(LogNet, "scanInterfaces() END")
}

// displays the IP address found on your network interfaces
func updateRealAAAA() {
	var all4 string
	var all6 string
	for s, t := range me.ipmap {
		if (t.ipv4) {
			all4 += s + "\n"
			debug(LogNet, "IPv4 =", s)
		} else if (t.ipv6) {
			all6 += s + "\n"
			debug(LogNet, "IPv6 =", s)
		} else {
			debug(LogNet, "???? =", s)
		}
	}
	all4 = sortLines(all4)
	all6 = sortLines(all6)
	if (me.IPv4.S != all4) {
		debug(LogNow, "IPv4 addresses have changed", all4)
		me.IPv4.SetText(all4)
	}
	if (me.IPv6.S != all6) {
		debug(LogNow, "IPv6 addresses have changed", all6)
		me.IPv6.SetText(all6)
	}
}

// delete network interfaces and ip addresses from the gui
func deleteChanges() bool {
	var changed bool = false
	for i, t := range me.ifmap {
		if (t.gone) {
			debug(LogChange, "DELETE int =", i, "name =", t.name, t.iface)
			delete(me.ifmap, i)
			changed = true
		}
		t.gone = true
	}
	for s, t := range me.ipmap {
		if (t.gone) {
			debug(LogChange, "DELETE name =", s, "IPv4 =", t.ipv4)
			debug(LogChange, "DELETE name =", s, "IPv6 =", t.ipv6)
			debug(LogChange, "DELETE name =", s, "iface =", t.iface)
			debug(LogChange, "DELETE name =", s, "ip =", t.ip)
			delete(me.ipmap, s)
			changed = true
		}
		t.gone = true
	}

	return changed
}
