// Various Linux/Unix'y things

// https://wiki.archlinux.org/title/Dynamic_DNS

package main

import 	(
	"log"
	"os"
	"os/exec"
	"net"
	"bytes"
	"fmt"
	"strings"

	"git.wit.org/wit/shell"
)

func CheckSuperuser() bool {
	return os.Getuid() == 0
}

func Escalate() {
	if os.Getuid() != 0 {
		cmd := exec.Command("sudo", "./control-panel-dns") // TODO: get the actual path
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			debug(LogError, "exit in Escalate()")
			exit(err)
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
		debug(LogError, "dumpIPs() failed:", err)
        }
        for _, ip := range ips {
                log.Println(host, ip)
        }
}

/*
	check if ddclient is installed, working, and/or configured
	https://github.com/ddclient/ddclient
*/
func ddclient() {
}

/*
	check if ddupdate is installed, working, and/or configured
*/
func ddupdate() {
}

func run(s string) string {
	cmdArgs := strings.Fields(s)
	// Define the command you want to run
	// cmd := exec.Command(cmdArgs)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)

	// Create a buffer to capture the output
	var out bytes.Buffer

	// Set the output of the command to the buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return ""
	}

	tmp := shell.Chomp(out.String())
	// Output the results
	debug(LogExec, "Command Output:", tmp)

	return tmp
}
