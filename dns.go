// Various Linux/Unix'y things

// https://wiki.archlinux.org/title/Dynamic_DNS

package main

import 	(
//	"os"
//	"os/exec"
	"log"
	"net"
//	"git.wit.org/wit/gui"
//	"github.com/davecgh/go-spew/spew"
)

type IPtype struct {
	// IP		string
	IPv4		bool
	IPv6		bool
	LinkLocal	bool
}

type Host struct {
	Name		string
	domainname	string
	hostname	string
	fqdn		string
	ips		map[string]*IPtype
}

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
                log.Fatal(err)
        }
        for _, ip := range ips {
                log.Println(host, ip)
        }
}
