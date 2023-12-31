package digitalocean

import 	(
	"errors"
	"sort"
	"strings"
	"github.com/digitalocean/godo"

	"go.wit.com/log"
	// "go.wit.com/gui"
)

func (d *DigitalOcean) NewDroplet(dd *godo.Droplet) *Droplet {
	if ! myDo.Ready() {return nil}

	// check if the droplet ID already exists
	if (d.dropMap[dd.ID] != nil) {
		log.Error(errors.New("droplet.NewDroplet() already exists"))
		return d.dropMap[dd.ID]
	}

	droplet := new(Droplet)
	droplet.ready = false
	droplet.poll = dd // the information polled from the digital ocean API
	droplet.ID = dd.ID

	if (d.dGrid == nil) {
		d.dGrid = d.group.NewGrid("grid", 10, 1).Pad()
	}

	droplet.name = d.dGrid.NewLabel(dd.Name)

	var ipv4 []string
	var ipv6 []string
	for _, network := range dd.Networks.V4 {
		if network.Type == "public" {
			ipv4 = append(ipv4, network.IPAddress)
		}
	}

	for _, network := range dd.Networks.V6 {
		if network.Type == "public" {
			ipv6 = append(ipv6, network.IPAddress)
		}
	}
	sort.Strings(ipv4)
	sort.Strings(ipv6)
	droplet.ipv4 = d.dGrid.NewLabel(strings.Join(ipv4, "\n"))
	droplet.ipv6 = d.dGrid.NewLabel(strings.Join(ipv6, "\n"))

	droplet.sizeSlug = d.dGrid.NewLabel(dd.SizeSlug)

	droplet.status = d.dGrid.NewLabel(dd.Status)

	droplet.connect = d.dGrid.NewButton("Connect", func () {
		droplet.Connect()
	})

	droplet.edit = d.dGrid.NewButton("Edit", func () {
		droplet.Show()
	})

	droplet.poweroff = d.dGrid.NewButton("Power Off", func () {
		droplet.PowerOff()
	})

	droplet.poweron = d.dGrid.NewButton("Power On", func () {
		droplet.PowerOn()
	})

	droplet.destroy = d.dGrid.NewButton("Destroy", func () {
		droplet.Destroy()
	})

	droplet.ready = true
	return droplet
}

func (d *Droplet) Active() bool {
	if ! d.Ready() {return false}
	log.Info("droplet.Active() status: ", d.poll.Status)
	if (d.status.S == "active") {
		return true
	}
	return false
}

// Returns true if the droplet is finished installing
func (d *Droplet) Ready() bool {
	if d == nil {return false}
	return d.ready
}

// Returns true if the droplet is running
func (d *Droplet) On() bool {
	if ! d.Ready() {return false}
	return true
}

func (d *Droplet) HasIPv4() bool {
	if ! d.Ready() {return false}
	return true
}

func (d *Droplet) GetIPv4() string {
	if ! d.Ready() {return ""}
	return d.ipv4.GetText()
}

func (d *Droplet) GetIPv6() string {
	if ! d.Ready() {return ""}
	log.Info("droplet GetIPv6 has: n.GetText()", d.ipv6.GetText())
	return d.ipv6.GetText()
}

func (d *Droplet) Connect() {
	if ! d.Ready() {return}
	if d.HasIPv4() {
		ipv4 := d.GetIPv4()
		log.Info("droplet has IPv4 =", ipv4)
	}
	log.Info("droplet.Connect() here", d.GetIPv4(), d.GetIPv6())
}

func (d *Droplet) Update(dpoll *godo.Droplet) {
	if ! d.Exists() {return}
	d.poll = dpoll
	log.Info("droplet.Update()", dpoll.Name, "dpoll.Status =", dpoll.Status)
	log.Spew(dpoll)
	d.status.SetText(dpoll.Status)
	if d.Active() {
		d.poweron.Disable()
		d.destroy.Disable()
		d.connect.Enable()
		d.poweroff.Enable()
	} else {
		d.poweron.Enable()
		d.destroy.Enable()
		d.poweroff.Disable()
		d.connect.Disable()
	}
}

func (d *Droplet) PowerOn() {
	if ! d.Exists() {return}
	log.Info("droplet.PowerOn() should do it here")
	myDo.PowerOn(d.ID)
}

func (d *Droplet) PowerOff() {
	if ! d.Exists() {return}
	log.Info("droplet.PowerOff() here")
	myDo.PowerOff(d.ID)
}

func (d *Droplet) Destroy() {
	if ! d.Exists() {return}
	log.Info("droplet.Destroy() ID =", d.ID, "Name =", d.name)
	myDo.deleteDroplet(d)
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
