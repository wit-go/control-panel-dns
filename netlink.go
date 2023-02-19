package main

// examples of what ifconfig does
// example of AF_NETLINK change:
// https://stackoverflow.com/questions/579783/how-to-detect-ip-address-change-programmatically-in-linux/2353441#2353441
// from that page, a link to watch for any ip event:
// https://github.com/angt/ipevent/blob/master/ipevent.c

// https://github.com/mdlayher/talks : Linux, Netlink, and Go in 7 minutes or less! (GopherCon 2018, lightning talk) 

/*
	c example from ipevent.c :
	int fd = socket(PF_NETLINK, SOCK_RAW, NETLINK_ROUTE);

	struct sockaddr_nl snl = {
	    .nl_family = AF_NETLINK,
	    .nl_groups = RTMGRP_IPV4_IFADDR | RTMGRP_IPV6_IFADDR,
	};
*/

/*
import 	(
//	"os"
//	"os/exec"
	// "log"
	// "net"
	// "unix"
	"github.com/vishvananda/netlink"
	"github.com/jsimonetti/rtnetlink"
//	"git.wit.org/wit/gui"
//	"github.com/davecgh/go-spew/spew"
)

// In golang, write a function     to register with netlink to detect changes to any network interface     Use tab indentation. Do not include example usage.

func registerNetlink() error {
	// Create netlink socket
	sock, err := netlink.Socket(rtnetlink.NETLINK_ROUTE, 0)
	if err != nil {
		return err
	}
	// Register for interface change events
	err = netlink.AddMembership(sock, netlink.RTNLGRP_LINK)
	if err != nil {
		return err
	}
	// Close the socket 
	defer sock.Close()
	// Handle incoming notifications
	for {
		msgs, _, err := sock.Receive()
		if err != nil {
			return err
		}
		for _, msg := range msgs {
			switch msg.Header.Type {
			case unix.RTM_NEWLINK:
				// Do something with new link
			case unix.RTM_DELLINK:
				// Do something with deleted link
			}
		}
	}
	return nil
}
*/
