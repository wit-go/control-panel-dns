// This creates a simple hello world window
package main

import 	(
	"net"
	"strings"
	// "git.wit.org/wit/gui"
)

var DEBUGNET bool = false

// this doesn't work
/*
func watchNetworkInterfaces() {
	// Get list of network interfaces
	interfaces, _ := net.Interfaces()

	// Set up a notification channel
	notification := make(chan net.Interface)

	log(DEBUGNET, "watchNet()")
	// Start goroutine to watch for changes
	go func() {
		log(DEBUGNET, "watchNet() func")
		for {
			log(DEBUGNET, "forever loop start")
			// Check for changes in each interface
			for _, i := range interfaces {
				log(DEBUGNET, "something on i =", i)
				if status := i.Flags & net.FlagUp; status != 0 {
					notification <- i
					log(DEBUGNET, "something on i =", i)
				}
			}
			log(DEBUGNET, "forever loop end")
		}
	}()
}
*/

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func (t *IPtype) IsReal() bool {
	if (t.ip.IsPrivate() || t.ip.IsLoopback() || t.ip.IsLinkLocalUnicast()) {
		log(DEBUGNET, "\t\tIP is Real = false")
		return false
	} else {
		log(DEBUGNET, "\t\tIP is Real = true")
		return true
	}
}

func IsReal(ip *net.IP) bool {
	if (ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast()) {
		log(DEBUGNET, "\t\tIP is Real = false")
		return false
	} else {
		log(DEBUGNET, "\t\tIP is Real = true")
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
		log(i.Name, "is a new network interface. The linux kernel index =", i.Index)
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
	log(args.VerboseNet, "me.ifmap[i] does exist. Need to compare everything.", i.Index, i.Name, val.iface.Index, val.iface.Name)
	if (val.iface.Name != i.Name) {
		log(val.iface.Name, "has changed to it's name to", i.Name)
		me.ifmap[i.Index].iface = &i
		me.changed = true
		if (me.Interfaces != nil) {
			me.Interfaces.AddText(i.Name)
			me.Interfaces.SetText(i.Name)
		}
		return
	}
}

func realAAAA() []string {
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
			log("\tIP is Real    ", ipt, i.Index, i.Name, s)
			if (t.ipv6) {
				ipv6s[s] = t
			} else {
				ipv4s[s] = t
			}
		} else {
			log("\tIP is not Real", ipt, i.Index, i.Name, s)
		}
	}
	return ipv6s, ipv4s
}

// Will figure out if an IP address is new
func checkIP(ip *net.IPNet, i net.Interface) bool {
	log(args.VerboseNet, "\t\taddr.(type) = *net.IPNet")
	log(args.VerboseNet, "\t\taddr.(type) =", ip)
	var realip string
	realip = ip.IP.String()

	val, ok := me.ipmap[realip]
	if ok {
		log(args.VerboseNet, val.ipnet.IP.String(), "is already a defined IP address")
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
		log("\tIP is Real    ", t, i.Index, i.Name, realip)
	} else {
		log("\tIP is not Real", t, i.Index, i.Name, realip)
	}
	log(args.VerboseNet, "\t\tIP is IsPrivate() =", ip.IP.IsPrivate())
	log(args.VerboseNet, "\t\tIP is IsLoopback() =", ip.IP.IsLoopback())
	log(args.VerboseNet, "\t\tIP is IsLinkLocalUnicast() =", ip.IP.IsLinkLocalUnicast())
	// log("HERE HERE", "realip =", realip, "me.ip[realip]=", me.ipmap[realip])
	return true
}

func scanInterfaces() {
	me.changed = false
	ifaces, _ := net.Interfaces()
	// me.ifnew = ifaces
	log(DEBUGNET, SPEW, ifaces)
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// log("range ifaces = ", i)
		checkInterface(i)
		log(args.VerboseNet, "*net.Interface.Name = ", i.Name, i.Index)
		log(args.VerboseNet, SPEW, i)
		log(DEBUGNET, SPEW, addrs)
		for _, addr := range addrs {
			log(DEBUGNET, "\taddr =", addr)
			log(DEBUGNET, SPEW, addrs)
			ips, _ := net.LookupIP(addr.String())
			log(DEBUGNET, "\tLookupIP(addr) =", ips)
			switch v := addr.(type) {
			case *net.IPNet:
				checkIP(v, i)
				// log("\t\tIP is () =", ip.())
			default:
				log(DEBUGNET, "\t\taddr.(type) = NO IDEA WHAT TO DO HERE v =", v)
			}
		}
	}
	deleteChanges()
	var all4 string
	var all6 string
	for s, t := range me.ipmap {
		log("HAVE name =", s, "IPv4 =", t.ipv4)
		log("HAVE name =", s, "IPv6 =", t.ipv6)
		if (t.ipv4) {
			all4 += s + "\n"
		}
		if (t.ipv6) {
			all6 += s + "\n"
		}
	}
	all4 = strings.TrimSuffix(all4, "\r\n")
	all4 = strings.TrimSuffix(all4, "\n")
	all6 = strings.TrimSuffix(all6, "\r\n")
	all6 = strings.TrimSuffix(all6, "\n")
	me.IPv4.SetText(all4)
	me.IPv6.SetText(all6)
}

// delete network interfaces and ip addresses from the gui
func deleteChanges() {
	for i, t := range me.ifmap {
		if (t.gone) {
			log("DELETE int =", i, "name =", t.name, t.iface)
			delete(me.ifmap, i)
			me.changed = true
		}
		t.gone = true
	}
	for s, t := range me.ipmap {
		if (t.gone) {
			log("DELETE name =", s, "IPv4 =", t.ipv4)
			log("DELETE name =", s, "IPv6 =", t.ipv6)
			log("DELETE name =", s, "iface =", t.iface)
			log("DELETE name =", s, "ip =", t.ip)
			delete(me.ipmap, s)
			me.changed = true
		}
		t.gone = true
	}
}
