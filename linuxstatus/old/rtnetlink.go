package main

import (
	"github.com/jsimonetti/rtnetlink"
	"go.wit.com/log"
)

// List all interfaces
func Example_listLink() {
	// Dial a connection to the rtnetlink socket
	conn, err := rtnetlink.Dial(nil)
	if err != nil {
		log.Error(err, "Example_listLink() failed")
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
