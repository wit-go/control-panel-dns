// GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// Copyright (c) 2023 WIT.COM, Inc.
// This is a control panel for DNS

package linuxstatus

import 	(
	"os"
	"os/user"
	"strconv"

	"go.wit.com/log"
)

func linuxLoop() {
	me.changed = false
	duration := timeFunction(getHostname)
	log.Log(INFO, "getHostname() execution Time: ", duration, "me.changed =", me.changed)

	duration = timeFunction(scanInterfaces)
	log.Log(NET, "scanInterfaces() execution Time: ", duration)
	for i, t := range me.ifmap {
		log.Log(NET, strconv.Itoa(i) + " iface = " + t.iface.Name)
	}

	var aaaa []string
	aaaa = dhcpAAAA()
	var all string
	for _, s := range aaaa {
		log.Log(NET, "my actual AAAA = ",s)
		all += s + "\n"
	}
	// me.IPv6.SetText(all)

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
