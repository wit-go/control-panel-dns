// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package linuxstatus

import 	(
	"os"
	"os/user"
	"strconv"
	"strings"

	"go.wit.com/log"
)

func linuxLoop() {
	me.changed = false
	duration := timeFunction(lookupHostname)
	log.Log(INFO, "getHostname() execution Time: ", duration, "me.changed =", me.changed)

	duration = timeFunction(scanInterfaces)
	log.Log(NET, "scanInterfaces() execution Time: ", duration)
	for i, t := range me.ifmap {
		log.Log(NET, strconv.Itoa(i) + " iface = " + t.iface.Name)
	}

	// get all the real AAAA records from all the network interfaces linux can see
	tmp := strings.Join(realAAAA(), "\n")
	tmp = sortLines(tmp)
	me.workingIPv6.Set(tmp)

	user, _ := user.Current()
	log.Log(INFO, "os.Getuid =", user.Username, os.Getuid())
	if (me.uid != nil) {
		me.uid.Set(user.Username + " (" + strconv.Itoa(os.Getuid()) + ")")
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
