// Various Linux/Unix'y things

// https://wiki.archlinux.org/title/Dynamic_DNS

package main

import 	(
	"os"
	"os/exec"
	"log"
	"net"
//	"git.wit.org/wit/gui"
//	"github.com/davecgh/go-spew/spew"
)

func CheckSuperuser() bool {
	return os.Getuid() == 0
}

func Escalate() {
	if os.Getuid() != 0 {
		cmd := exec.Command("sudo", "./control-panel-dns")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// You need permission to do a zone transfer. Otherwise:
// dig +noall +answer +multiline lab.wit.org any 
// dig +all +multiline fire.lab.wit.org # gives the zonefile header (ttl vals)
func DumpPublicDNSZone(zone string) {
	entries, err := net.LookupHost(zone)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		log.Println(entry)
	}
}

func dumpIPs(host string) {
    ips, err := net.LookupIP(host)
        if err != nil {
                log.Fatal(err)
        }
        for _, ip := range ips {
                log.Println(host, ip)
        }
}
