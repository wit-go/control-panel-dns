package digitalocean

import 	(
	"github.com/digitalocean/godo"

	"go.wit.com/log"
	// "go.wit.com/gui"
)

func (d *DigitalOcean) NewDroplet(dd *godo.Droplet) *Droplet {
	if ! myDo.Ready() {return nil}

	droplet := new(Droplet)
	droplet.ready = false
	droplet.poll = dd // the information polled from the digital ocean API

	if (d.dGrid == nil) {
		d.dGrid = d.group.NewGrid("grid", 5, 1).Pad()
	}

	droplet.name = d.dGrid.NewLabel(dd.Name)

	droplet.box4 = d.dGrid.NewBox("hBox", true)
	droplet.grid4 = droplet.box4.NewGrid("grid", 2, 1).Pad()
	for _, network := range dd.Networks.V4 {
		if network.Type == "public" {
			droplet.grid4.NewLabel(network.IPAddress)
		}
	}
	droplet.box6 = d.dGrid.NewBox("hBox", true)
	droplet.grid6 = droplet.box6.NewGrid("grid", 2, 1).Pad()
	for _, network := range dd.Networks.V6 {
		if network.Type == "public" {
			droplet.grid6.NewLabel(network.IPAddress)
		}
	}

	d.dGrid.NewButton("Connect", func () {
		droplet.Connect()
	})

	d.dGrid.NewButton("Show", func () {
		droplet.Show()
	})

	droplet.ready = true
	return droplet
}

func (d *Droplet) Connect() {
	if ! d.Exists() {return}
	log.Info("droplet.Connect() ssh here")
}

/*
type Droplet struct {
	ID               int           `json:"id,float64,omitempty"`
	Name             string        `json:"name,omitempty"`
	Memory           int           `json:"memory,omitempty"`
	Vcpus            int           `json:"vcpus,omitempty"`
	Disk             int           `json:"disk,omitempty"`
	Region           *Region       `json:"region,omitempty"`
	Image            *Image        `json:"image,omitempty"`
	Size             *Size         `json:"size,omitempty"`
	SizeSlug         string        `json:"size_slug,omitempty"`
	BackupIDs        []int         `json:"backup_ids,omitempty"`
	NextBackupWindow *BackupWindow `json:"next_backup_window,omitempty"`
	SnapshotIDs      []int         `json:"snapshot_ids,omitempty"`
	Features         []string      `json:"features,omitempty"`
	Locked           bool          `json:"locked,bool,omitempty"`
	Status           string        `json:"status,omitempty"`
	Networks         *Networks     `json:"networks,omitempty"`
	Created          string        `json:"created_at,omitempty"`
	Kernel           *Kernel       `json:"kernel,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	VolumeIDs        []string      `json:"volume_ids"`
	VPCUUID          string        `json:"vpc_uuid,omitempty"`
}
*/
func (d *Droplet) Show() {
	if ! d.Exists() {return}
	log.Info("droplet:", d.name.Name)
	log.Info("droplet:", d.poll.ID, d.poll.Name, d.poll.Memory, d.poll.Disk, d.poll.Status)
	log.Spew(d.poll)
}

func (d *Droplet) Hide() {
	if ! d.Exists() {return}
	log.Info("droplet.Hide() window")
	if ! d.hidden {
		// d.window.Hide()
	}
	d.hidden = true
}

func (d *Droplet) Exists() bool {
	if ! myDo.Ready() {return false}
	if d == nil {return false}
	if d.poll == nil {return false}
	return d.ready
}
