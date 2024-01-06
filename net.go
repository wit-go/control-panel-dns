// This creates a simple hello world window
package main

import 	(
	// "log"
	"net"
	"strings"

	"go.wit.com/log"
)

// this doesn't work
/*
func watchNetworkInterfaces() {
	// Get list of network interfaces
	interfaces, _ := net.Interfaces()

	// Set up a notification channel
	notification := make(chan net.Interface)

	log.Log(NET, "watchNet()")
	// Start goroutine to watch for changes
	go func() {
		log.Log(NET, "watchNet() func")
		for {
			log.Log(NET, "forever loop start")
			// Check for changes in each interface
			for _, i := range interfaces {
				log.Log(NET, "something on i =", i)
				if status := i.Flags & net.FlagUp; status != 0 {
					notification <- i
					log.Log(NET, "something on i =", i)
				}
			}
			log.Log(NET, "forever loop end")
		}
	}()
}
*/

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
		log.Info(i.Name, "is a new network interface. The linux kernel index =", i.Index)
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
	log.Log(NET, "me.ifmap[i] does exist. Need to compare everything.", i.Index, i.Name, val.iface.Index, val.iface.Name)
	if (val.iface.Name != i.Name) {
		log.Info(val.iface.Name, "has changed to it's name to", i.Name)
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

// delete network interfaces and ip addresses from the gui
func deleteChanges2() bool {
	var changed bool = false
	for i, t := range me.ifmap {
		if (t.gone) {
			log.Log(CHANGE, "DELETE int =", i, "name =", t.name, t.iface)
			delete(me.ifmap, i)
			changed = true
		}
		t.gone = true
	}
	for s, t := range me.ipmap {
		if (t.gone) {
			log.Log(CHANGE, "DELETE name =", s, "IPv4 =", t.ipv4)
			log.Log(CHANGE, "DELETE name =", s, "IPv6 =", t.ipv6)
			log.Log(CHANGE, "DELETE name =", s, "iface =", t.iface)
			log.Log(CHANGE, "DELETE name =", s, "ip =", t.ip)
			delete(me.ipmap, s)
			changed = true
		}
		t.gone = true
	}

	return changed
}
