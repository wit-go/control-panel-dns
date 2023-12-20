package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getProcessNameByPort(port int) string {
	// Convert port to hex string
	portHex := strconv.FormatInt(int64(port), 16)

	// Function to search /proc/net/tcp or /proc/net/udp
	searchProcNet := func(file string) string {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return ""
		}
		// debug(LogProc, "searchProcNet() data:", string(data))

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			debug(LogProc, "searchProcNet() portHex:", portHex)
			if (len(fields) > 9) {
				debug(LogProc, "searchProcNet() fields[9]", fields[9])
			}
			debug(LogProc, "searchProcNet() lines:", line)
			if len(fields) > 1 {
				parts := strings.Split(fields[1], ":")
				if len(parts) > 1 {
					// Convert the hexadecimal string to an integer
					value, _ := strconv.ParseInt(parts[1], 16, 64)
					debug(LogProc, "searchProcNet() value, port =", value, port, "parts[1] =", parts[1])
					if (port == int(value)) {
						debug(LogProc, "searchProcNet() THIS IS THE LINE:", fields)
						return fields[9]
					}
				}
			}
		}

		return ""
	}

	// Search TCP and then UDP
	inode := searchProcNet("/proc/net/tcp")
	if inode == "" {
		inode = searchProcNet("/proc/net/udp")
	}
	debug(LogProc, "searchProcNet() inode =", inode)

	// Search for process with the inode
	procs, _ := ioutil.ReadDir("/proc")
	for _, proc := range procs {
		if !proc.IsDir() {
			continue
		}

		fdPath := filepath.Join("/proc", proc.Name(), "fd")
		fds, err := ioutil.ReadDir(fdPath)
		if err != nil {
			continue // Process might have exited; skip it
		}

		for _, fd := range fds {
			fdLink, _ := os.Readlink(filepath.Join(fdPath, fd.Name()))
			var s string
			s = "socket:["+inode+"]"
			if strings.Contains(fdLink, "socket:[") {
				debug(LogProc, "searchProcNet() fdLink has socket:", fdLink)
				debug(LogProc, "searchProcNet() proc.Name() =", proc.Name(), "s =", s)
			}
			if strings.Contains(fdLink, "socket:[35452]") {
				debug(LogProc, "searchProcNet() found proc.Name() =", proc.Name(), fdLink)
				return proc.Name()
			}
			if strings.Contains(fdLink, "socket:[35450]") {
				debug(LogProc, "searchProcNet() found proc.Name() =", proc.Name(), fdLink)
				return proc.Name()
			}
			if strings.Contains(fdLink, "socket:[35440]") {
				debug(LogProc, "searchProcNet() found proc.Name() =", proc.Name(), fdLink)
				return proc.Name()
			}
			if strings.Contains(fdLink, "socket:[21303]") {
				debug(LogProc, "searchProcNet() found proc.Name() =", proc.Name(), fdLink)
				// return proc.Name()
			}
			if strings.Contains(fdLink, "socket:["+inode+"]") {
				return proc.Name()
			}
		}
	}

	return ""
}
