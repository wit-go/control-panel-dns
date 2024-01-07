// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.

// This is a control panel for DNS

// This is the main Linux kernel / OS code
// to check your network settings are correct
// This does (and should do) no network or external checking
// This is just the state of your OS

package linuxstatus

import 	(
	"os"
	"os/user"
	"io/ioutil"
	"strconv"
	"strings"
	"sort"

	"go.wit.com/log"
)

func linuxLoop() {
	me.changed = false

	// checks for a VALID hostname
	lookupHostname()
	if me.changed {
		log.Log(CHANGE, "lookupHostname() detected a change")
	}

	// scans the linux network intefaces for your available IPv4 & IPv6 addresses
	scanInterfaces()
	if me.changed {
		log.Log(CHANGE, "scanInterfaces() detected a change")
	}
	for i, t := range me.ifmap {
		log.Log(NET, strconv.Itoa(i) + " iface = " + t.iface.Name)
	}

	// get all the real A records from all the network interfaces linux can see
	a := realA()
	sort.Strings(a)
	tmp := strings.Join(a, "\n")
	if tmp != me.workingIPv4.Get() {
		log.Log(CHANGE, "realAAAA() your real IPv6 addresses changed")
		me.changed = true
		me.workingIPv4.Set(tmp)
	}

	// get all the real AAAA records from all the network interfaces linux can see
	aaaa := realAAAA()
	sort.Strings(aaaa)
	tmp = strings.Join(aaaa, "\n")
	if tmp != me.workingIPv6.Get() {
		log.Log(CHANGE, "realAAAA() your real IPv6 addresses changed")
		me.changed = true
		me.workingIPv6.Set(tmp)
	}

	user, _ := user.Current()
	tmp = user.Username + " (" + strconv.Itoa(os.Getuid()) + ")"
	if tmp != me.uid.Get() {
		log.Log(CHANGE, "os.Getuid =", user.Username, os.Getuid())
		me.changed = true
		me.uid.Set(tmp)
	}

	content, _ := ioutil.ReadFile("/etc/resolv.conf")
	var ns []string
	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.Split(line, " ")
		if len(parts) > 1 {
			if parts[0] == "nameserver" {
				ns = append(ns, parts[1])
			}
		}
	}
	sort.Strings(ns)
	newNS := strings.Join(ns, "\n")
	if newNS != me.resolver.Get() {
		log.Log(CHANGE, "resolver changed in /etc/resolv.conf to", ns)
		me.changed = true
		me.resolver.Set(newNS)
	}

	/*
	processName := getProcessNameByPort(53)
	fmt.Println("Process with port 53:", processName)

	commPath := filepath.Join("/proc", proc.Name(), "comm")
	comm, err := ioutil.ReadFile(commPath)
	if err != nil {
		return "", err // Error reading the process name
	}
	return strings.TrimSpace(string(comm)), nil
	*/
}
