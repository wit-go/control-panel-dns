package main

import (
	"github.com/jsimonetti/rtnetlink"
)

// List all interfaces
func Example_listLink() {
	// Dial a connection to the rtnetlink socket
	conn, err := rtnetlink.Dial(nil)
	if err != nil {
		exit(err)
	}
	defer conn.Close()

	// Request a list of interfaces
	msg, err := conn.Link.List()
	if err != nil {
		log(err)
	}

	log("%#v", msg)
	log(SPEW, msg)
}
