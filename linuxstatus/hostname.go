// figures out if your hostname is valid
// then checks if your DNS is setup correctly
package linuxstatus

import (
	"go.wit.com/log"
	"go.wit.com/shell"

	// will try to get this hosts FQDN
	"github.com/Showmax/go-fqdn"
)

func (ls *LinuxStatus) GetDomainName() string {
	if ! me.Ready() {return ""}
	return me.domainname.Get()
}

func (ls *LinuxStatus) setDomainName(dn string) {
	if ! me.Ready() {return}
	me.domainname.Set(dn)
}

func getHostname() {
	var err error
	var s string = "gui.Label == nil"
	s, err = fqdn.FqdnHostname()
	if (err != nil) {
		log.Error(err, "FQDN hostname error")
		return
	}
	log.Warn("full hostname should be:", s)

	dn := run("domainname")
	if (me.domainname.Get() != dn) {
		log.Log(CHANGE, "domainname has changed from", me.GetDomainName(), "to", dn)
		me.setDomainName(dn)
		me.changed = true
	}

	hshort := run("hostname -s")
	if (me.hostshort.Get() != hshort) {
		log.Log(CHANGE, "hostname -s has changed from", me.hostshort.Get(), "to", hshort)
		me.hostshort.Set(hshort)
		me.changed = true
	}

	/*
	var test string
	test = hshort + "." + dn
	if (me.status.GetHostname() != test) {
		log.Log(CHANGE, "me.hostname", me.status.GetHostname(), "does not equal", test)
		if (me.hostnameStatus.S != "BROKEN") {
			log.Log(CHANGE, "me.hostname", me.status.GetHostname(), "does not equal", test)
			me.changed = true
			me.hostnameStatus.SetText("BROKEN")
		}
	} else {
		if (me.hostnameStatus.S != "VALID") {
			log.Log(CHANGE, "me.hostname", me.status.GetHostname(), "is valid")
			me.hostnameStatus.SetText("VALID")
			me.changed = true
		}
	}
	*/
}

// returns true if the hostname is good
// check that all the OS settings are correct here
// On Linux, /etc/hosts, /etc/hostname
//      and domainname and hostname
func goodHostname() bool {
	hostname := shell.Chomp(shell.Cat("/etc/hostname"))
	log.Log(NOW, "hostname =", hostname)

	hs := run("hostname -s")
	dn := run("domainname")
	log.Log(NOW, "hostname short =", hs, "domainname =", dn)

	tmp := hs + "." + dn
	if (hostname == tmp) {
		log.Log(NOW, "hostname seems to be good", hostname)
		return true
	}

	return false
}