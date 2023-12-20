// inspired from:
// https://github.com/mactsouk/opensource.com.git
// and
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

package dnssecsocket

import "os"
import "bufio"
import "math/rand"
import "net"
import "strconv"
import "strings"
// import log "github.com/sirupsen/logrus"
// import "github.com/wercker/journalhook"

import "go.wit.com/shell"

// will try to get this hosts FQDN
// import "github.com/Showmax/go-fqdn"

import "github.com/miekg/dns"

// import "github.com/davecgh/go-spew/spew"

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func GetRemoteAddr(conn net.TCPConn) string {
	clientAddr := conn.RemoteAddr().String()
	parts := strings.Split(clientAddr, "]")
	ipv6 := parts[0]
	return ipv6[1:]
}

//
// Handle each connection
// Each client must send it's hostname as the first line
// Then each hostname is verified with DNSSEC
//
func HandleConnection(conn *net.TCPConn) {
	// Disable journalhook until it builds on Windows
	// journalhook.Enable()

	// spew.Dump(conn)
	// ipv6client := GetRemoteAddr(c)
	ipv6client := conn.RemoteAddr()
	log(args.VerboseDnssec, "Serving to %s as the IPv6 client", ipv6client)

	// setup this TCP socket as the "standard input"
	// newStdin, _ := bufio.NewReader(conn.File())
	newStdin, _ := conn.File()
	newreader := bufio.NewReader(newStdin)

	log(args.VerboseDnssec, "Waiting for the client to tell me its name")
	netData, err := newreader.ReadString('\n')
	if err != nil {
		log(args.VerboseDnssec, err)
		return
	}
	clientHostname := strings.TrimSpace(netData)
	log(args.VerboseDnssec, "Recieved client hostname as:", clientHostname)

	dnsRR := Dnstrace(clientHostname, "AAAA")
	if (dnsRR == nil) {
		log(args.VerboseDnssec, "dnsRR IS NIL")
		log(args.VerboseDnssec, "dnsRR IS NIL")
		log(args.VerboseDnssec, "dnsRR IS NIL")
		conn.Close()
		return
	}
	ipaddr := dns.Field(dnsRR[1], 1)
	log(args.VerboseDnssec, "Client claims to be:   ", ipaddr)
	log(args.VerboseDnssec, "Serving to IPv6 client:", ipv6client)

/* TODO: figure out how to fix this check
	if (ipaddr != ipv6client) {
		log(args.VerboseDnssec)
		log(args.VerboseDnssec, "DNSSEC ERROR: client IPv6 does not work")
		log(args.VerboseDnssec, "DNSSEC ERROR: client IPv6 does not work")
		log(args.VerboseDnssec, "DNSSEC ERROR: client IPv6 does not work")
		log(args.VerboseDnssec)
		conn.Close()
		return
	}
*/

	f, _ := conn.File()
//	shell.SetStdout(f)
//	shell.SpewOn()		// turn this on if you want to look at the process exit states

	// send all log() output to systemd journalctl
//	shell.UseJournalctl()

	for {
		defer shell.SetStdout(os.Stdout)
		defer conn.Close()
		netData, err := newreader.ReadString('\n')
		if err != nil {
			log(args.VerboseDnssec, err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		log(args.VerboseDnssec, "Recieved: ", temp)

		if (temp == "list") {
			log(args.VerboseDnssec, "Should run list here")
			shell.SetStdout(f)
			shell.Run("/root/bin/list.testing.com")
			shell.SetStdout(os.Stdout)
		}

		if (temp == "cpuinfo") {
			log(args.VerboseDnssec, "Should cat /proc/cpuinfo")
			shell.SetStdout(f)
			shell.Run("cat /proc/cpuinfo")
			shell.SetStdout(os.Stdout)
		}

		result := strconv.Itoa(random()) + "\n"
		conn.Write([]byte(string(result)))
	}
}
