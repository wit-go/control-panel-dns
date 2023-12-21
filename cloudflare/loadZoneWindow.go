// This is a simple example
package cloudflare

import 	(
	"log"
	"strconv"

	"go.wit.com/gui"
)

func LoadZoneWindow(n *gui.Node, c *ConfigT) {
	hostname := c.Domain
	zoneID := c.ZoneID
	log.Println("adding DNS record", hostname)

	newt := n.NewTab(hostname)
	vb := newt.NewBox("vBox", false)
	newg := vb.NewGroup("more zoneID = " + zoneID)

	// make a grid 6 things wide
	grid := newg.NewGrid("gridnuts", 6, 1)

//	grid.NewButton("Type", func () {
//		log.Println("sort by Type")
//	})
	grid.NewLabel("RR type")
	grid.NewLabel("hostname")

	grid.NewLabel("Proxy")
	grid.NewLabel("TTL")
	grid.NewLabel("Value")
	grid.NewLabel("Save")

	records := GetZonefile(c)
	for _, record := range records.Result {
		var rr RRT // dns zonefile resource record

		// copy all the JSON values into the row record.
		rr.ID = record.ID
		rr.Type = record.Type
		rr.Name = record.Name
		rr.Content = record.Content
		rr.Proxied = record.Proxied
		rr.Proxiable = record.Proxiable
		rr.ZoneID = zoneID
		// rr.Ttl = record.TTL

		rr.Domain = hostname
		rr.ZoneID = zoneID
		rr.Auth = c.Auth
		rr.Email = c.Email

		grid.NewLabel(record.Type)
		grid.NewLabel(record.Name)

		proxy := grid.NewLabel("proxy")
		if (record.Proxied) {
			proxy.SetText("On")
		} else {
			proxy.SetText("Off")
		}

		var ttl  string
		if (record.TTL == 1) {
			ttl = "Auto"
		} else {
			ttl = strconv.Itoa(record.TTL)
		}
		grid.NewLabel(ttl)

		val := grid.NewLabel("Value")
		val.SetText(record.Content)

		load := grid.NewButton("Load", nil)
		load.Custom = func () {
			name := "save stuff to cloudflare for " + rr.ID
			log.Println(name)

			/*
			rr.Domain = domainWidget.S
			rr.ZoneID = zoneWidget.S
			rr.Auth = authWidget.S
			rr.Email = emailWidget.S
			*/

			SetRow(&rr)
		}
	}

	grid.Pad()
}
