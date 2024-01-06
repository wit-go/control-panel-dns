// inspired from:
// https://github.com/mactsouk/opensource.com.git
// and
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

package main

import (
)

//	./go-nsupdate \
//		--tsig-algorithm=hmac-sha512 \
//		--tsig-secret="OWh5/ZHIyaz7B8J9m9ZDqZ8448Pke0PTpkYbZmFcOf5a6rEzgmcwrG91u1BHi1/4us+mKKEobDPLw1x6sD+ZJw==" \
//		-i eno2 farm001.lab.wit.com

/*
func nsupdate() {
	var tsigSecret string
	log.Log(NET, "nsupdate() START")
	cmd := "go-nsupdate --tsig-algorithm=hmac-sha512"
	tsigSecret = os.Getenv("TIG_SECRET")
	cmd += " --tig-secret=\"" + tsigSecret + "\""
	cmd += " -i wlo1 " + me.statusOS.GetHostname()
	log.Log(NET, "nsupdate() RUN:", cmd)

	for s, t := range me.ipmap {
		if (t.IsReal()) {
			if (t.ipv6) {
				log.Log(NET, "nsupdate() found real AAAA =", s, "on iface", t.iface.Name)
			}
		}
	}
}
*/
