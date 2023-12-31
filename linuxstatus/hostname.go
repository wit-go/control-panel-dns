// figures out if your hostname is valid
// then checks if your DNS is setup correctly
package linuxstatus

import (
	"io/ioutil"
	"go.wit.com/log"

	// will try to get this hosts FQDN
	"github.com/Showmax/go-fqdn"
)

func (ls *LinuxStatus) GetDomainName() string {
	if ! me.Ready() {return ""}
	return me.domainname.Get()
}

func (ls *LinuxStatus) setDomainName() {
	if ! me.Ready() {return}

	dn := run("domainname")
	if (me.domainname.Get() != dn) {
		log.Log(CHANGE, "domainname has changed from", me.GetDomainName(), "to", dn)
		me.domainname.Set(dn)
		me.changed = true
	}
}

func (ls *LinuxStatus) GetHostname() string {
	if ! me.Ready() {return ""}
	return me.hostname.Get()
}

func (ls *LinuxStatus) ValidHostname() bool {
	if ! me.Ready() {return false}
	if me.hostnameStatus.Get() == "WORKING" {
		return true
	}
	return false
}

func (ls *LinuxStatus) setHostname(newname string) {
	if ! me.Ready() {return}
	if newname ==  me.hostname.Get() {
		return
	}
	log.Log(CHANGE, "hostname has changed from", me.GetHostname(), "to", newname)
	me.hostname.Set(newname)
	me.changed = true
}

func (ls *LinuxStatus) GetHostShort() string {
	if ! me.Ready() {return ""}
	return me.hostshort.Get()
}

func (ls *LinuxStatus) setHostShort() {
	if ! me.Ready() {return}
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
	var hostfqdn string = "broken"
	hostfqdn, err = fqdn.FqdnHostname()
	if (err != nil) {
		log.Error(err, "FQDN hostname error")
		return
	}
	log.Log(NET, "full hostname should be: ", hostfqdn)

	me.setDomainName()
	me.setHostShort()

	// these are authoritative
	// if they work wrong, your linux configuration is wrong.
	// Do not complain.
	// Fix your distro if your box is otherwise not working this way
	hshort := me.GetHostShort() // from `hostname -s`
	dn := me.GetDomainName() // from `domanname`
	hostname := me.GetHostname() // from `hostname -f`

	if hostfqdn != hostname {
		log.Log(WARN, "hostname", hostname, "does not equal fqdn.FqdnHostname()", hostfqdn)
		// TODO: figure out what is wrong
	}

	var test string
	test = hshort + "." + dn

	me.setHostname(test)

	if (hostname != test) {
		log.Log(CHANGE, "hostname", hostname, "does not equal", test)
		if (me.hostnameStatus.Get() != "BROKEN") {
			log.Log(CHANGE, "hostname", hostname, "does not equal", test)
			me.changed = true
			me.hostnameStatus.Set("BROKEN")
		}
	} else {
		if (me.hostnameStatus.Get() != "WORKING") {
			log.Log(CHANGE, "hostname", hostname, "is valid")
			me.hostnameStatus.Set("WORKING")
			me.changed = true
		}
	}
}

// returns true if the hostname is good
// check that all the OS settings are correct here
// On Linux, /etc/hosts, /etc/hostname
//      and domainname and hostname
func goodHostname() bool {
	content, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		// this needs to be a fixWindow() error
		log.Error(err)
	}

	hostname := string(content)

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
