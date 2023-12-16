package main

import (
	"log"
	"github.com/jsimonetti/rtnetlink"
)

// List all interfaces
func Example_listLink() {
	// Dial a connection to the rtnetlink socket
	conn, err := rtnetlink.Dial(nil)
	if err != nil {
		log.Println(logError, "Example_listLink() failed", err)
		return
	}
	defer conn.Close()

	// Request a list of interfaces
	msg, err := conn.Link.List()
	if err != nil {
		log.Println(err)
	}

	log.Println("%#v", msg)
	log.Println(SPEW, msg)
}
