package digitalocean

import 	(
	"fmt"
	"github.com/digitalocean/godo"

	"go.wit.com/log"
	// "go.wit.com/gui"
)

func (d *DigitalOcean) NewDroplet(dd godo.Droplet) *Droplet {
	if ! myDo.Ready() {return nil}

	droplet := new(Droplet)
	droplet.ready = false
	droplet.poll = dd // the information polled from the digital ocean API

	if (d.dGrid == nil) {
		d.dGrid = d.group.NewGrid("grid", 2, 1).Pad()
	}

	droplet.name = d.dGrid.NewLabel(dd.Name)

	droplet.box4 = d.dGrid.NewBox("hBox", true)
	droplet.grid4 = droplet.box4.NewGrid("grid", 2, 1).Pad()

	fmt.Printf("Droplet: %s\n", dd.Name)
	for _, network := range dd.Networks.V4 {
		if network.Type == "public" {
			fmt.Printf("IPv4: %s\n", network.IPAddress)
			droplet.grid4.NewLabel(network.IPAddress)
			droplet.grid4.NewButton("Connect", func () {
				log.Info("ssh here", network.IPAddress)
			})
		}
	}
	for _, network := range dd.Networks.V6 {
		if network.Type == "public" {
			fmt.Printf("IPv6: %s\n", network.IPAddress)
		}
	}
	fmt.Println("-------------------------")

	droplet.ready = true
	return droplet
}

func (d *Droplet) Show() {
	if ! myDo.Ready() {return}
	log.Info("droplet.Show() window")
	if d.hidden {
		// my.window.Show()
	}
	d.hidden = false
}

func (d *Droplet) Hide() {
	if ! myDo.Ready() {return}
	log.Info("droplet.Hide() window")
	if ! d.hidden {
		// d.window.Hide()
	}
	d.hidden = true
}

func (d *Droplet) Exists() bool {
	if ! myDo.Ready() {return false}
	return true
}
