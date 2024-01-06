// figures out if your hostname is valid
// then checks if your DNS is setup correctly
package linuxstatus

import (
	"errors"

	"go.wit.com/log"
	"go.wit.com/shell"

	// will try to get this hosts FQDN
	"github.com/Showmax/go-fqdn"
)

func (ls *LinuxStatus) GetDomainName() string {
	if ! me.Ready() {return ""}
	if me.window == nil {
		log.Log(NOW, "me.window == nil")
	} else {
		log.Log(NOW, "me.window exists, but has not been drawn")
	}
	return me.domainname.Get()
}

func (ls *LinuxStatus) setDomainName() {
	if ! me.Ready() {return}

	dn := run("domainname")
	if me.window == nil {
		log.Log(NOW, "me.window == nil")
	} else {
		log.Log(NOW, "me.window exists, but has not been drawn")
		log.Log(NOW, "me.window.Draw() =")
	}
	if (me.domainname.Get() != dn) {
		log.Log(CHANGE, "domainname has changed from", me.GetDomainName(), "to", dn)
		me.domainname.Set(dn)
		me.changed = true
	}
}

func (ls *LinuxStatus) GetHostShort() string {
	if ! me.Ready() {return ""}
	if me.window == nil {
		log.Log(NOW, "me.window == nil")
	} else {
		log.Log(NOW, "me.window exists, but has not been drawn")
	}
	return me.hostshort.Get()
}

func (ls *LinuxStatus) setHostShort() {
	if ! me.Ready() {return ""}
	hshort := run("hostname -s")
	if (me.hostshort.Get() != hshort) {
		log.Log(CHANGE, "hostname -s has changed from", me.hostshort.Get(), "to", hshort)
		me.hostshort.Set(hshort)
		me.changed = true
	}
}

func lookupHostname() {
	if ! me.Ready() {return}
	var err error
	var s string = "gui.Label == nil"
	s, err = fqdn.FqdnHostname()
	if (err != nil) {
		log.Error(err, "FQDN hostname error")
		return
	}
	log.Error(errors.New("full hostname should be: " + s))

	me.setDomainName()


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
