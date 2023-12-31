package digitalocean

import 	(
	"os"
	"go.wit.com/log"
	"go.wit.com/gui"
)

var myDo *DigitalOcean

// This is initializes the main DO object
// You can only have one of these
func New(p *gui.Node) *DigitalOcean {
	if myDo != nil {return myDo}
	myDo = new(DigitalOcean)
	myDo.parent = p

	myDo.ready = false

	// Your personal API token from DigitalOcean.
	myDo.token = os.Getenv("DIGITALOCEAN_TOKEN")

	myDo.window = p.NewWindow("DigitalOcean Control Panel")

	// make a group label and a grid
	myDo.group = myDo.window.NewGroup("droplets:").Pad()
	myDo.grid = myDo.group.NewGrid("grid", 2, 1).Pad()

	myDo.ready = true
	myDo.Hide()
	return myDo
}

// Returns true if the status is valid
func (d *DigitalOcean) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *DigitalOcean) Show() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Show() window")
	if d.hidden {
		d.window.Show()
	}
	d.hidden = false
}

func (d *DigitalOcean) Hide() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Hide() window")
	if ! d.hidden {
		d.window.Hide()
	}
	d.hidden = true
}

func (d *DigitalOcean) Update() bool {
	if ! d.Ready() {return false}
	if ! d.ListDroplets() {
		log.Error(d.err, "Error listing droplets")
		return false
	}
	for _, droplet := range d.droplets {
		d.NewDroplet(&droplet)
	}
	return true
}
