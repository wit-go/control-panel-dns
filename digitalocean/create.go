package digitalocean

import (
	"context"
	"golang.org/x/oauth2"
	"github.com/digitalocean/godo"

	"go.wit.com/log"
	"go.wit.com/gui/gadgets"
	// "go.wit.com/gui"
)

/*
// createDroplet creates a new droplet in the specified region with the given name.
func createDroplet(token, name, region, size, image string) (*godo.Droplet, error) {
	// Create an OAuth2 token.
	tokenSource := &oauth2.Token{
		AccessToken: token,
	}

	// Create an OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// Create a DigitalOcean client with the OAuth2 client.
	client := godo.NewClient(oauthClient)

	// Define the create request.
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: image,
		},
	}

	// Create the droplet.
	ctx := context.TODO()
	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}

	return newDroplet, nil
}
*/

func (d *DigitalOcean) Create(name string, region string) {
	// region := "nyc1" // New York City region.
	size := "s-1vcpu-1gb" // Size of the droplet.
	image := "ubuntu-20-04-x64" // Image slug for Ubuntu 20.04 (LTS) x64.

	return
	// Create a new droplet.
	droplet, err := d.createDropletNew(name, region, size, image)
	if err != nil {
		log.Fatalf("Something went wrong: %s\n", err)
	}

	log.Printf("Created droplet ID %d with name %s\n", droplet.ID, droplet.Name)
}

// createDroplet creates a new droplet in the specified region with the given name.
func (d *DigitalOcean) createDropletNew(name, region, size, image string) (*godo.Droplet, error) {
	// Create an OAuth2 token.
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})

	// Create an OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// Create a DigitalOcean client with the OAuth2 client.
	client := godo.NewClient(oauthClient)

	var sshKeys []godo.DropletCreateSSHKey

	// Find the key by name.
	for i, key := range d.sshKeys {
		log.Info("found ssh i =", i, key.Name)
		log.Verbose("found ssh key.Name =", key.Name)
		log.Verbose("found ssh key.Fingerprint =", key.Fingerprint)
		log.Verbose("found ssh key:", key)
		/*
		sshKeys = []godo.DropletCreateSSHKey{
			{ID: key.ID},
		}
		*/
		sshKeys = append(sshKeys, godo.DropletCreateSSHKey{ID: key.ID})
	} 

	// Define the create request.
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: image,
		},
		IPv6: true, // Enable IPv6
		SSHKeys: sshKeys, // Add SSH key IDs here
	}

	// Create the droplet.
	ctx := context.TODO()
	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}

	return newDroplet, nil
}

var myCreate *windowCreate

// This is initializes the main DO object
// You can only have one of these
func InitCreateWindow() *windowCreate {
	if ! myDo.Ready() {return nil}
	if myCreate != nil {return myCreate}
	myCreate = new(windowCreate)
	myCreate.ready = false

	myCreate.window = myDo.parent.NewWindow("Create Droplet")

	// make a group label and a grid
	myCreate.group = myCreate.window.NewGroup("droplets:").Pad()
	myCreate.grid = myCreate.group.NewGrid("grid", 2, 1).Pad()
	
	myCreate.name = gadgets.NewBasicEntry(myCreate.grid, "Name").Set("test.wit.com")


	myCreate.zone = gadgets.NewBasicDropdown(myCreate.grid, "Region")

	regions := myDo.listRegions()

	// Print details of each region.
	log.Info("Available Regions:")
	for i, region := range regions {
		log.Infof("i: %d, Slug: %s, Name: %s, Available: %v\n", i, region.Slug, region.Name, region.Available)
		log.Spew(i, region)
		myCreate.zone.Add(region.Name)
	}

	var zone godo.Region
	myCreate.zone.Custom = func() {
		s := myCreate.zone.Get()
		log.Info("create droplet region changed to:", s)
		for _, region := range regions {
			if s == region.Name {
				log.Info("Found region! slug =", region.Slug, region)
				zone = region
			}
		}
	}

	myCreate.grid.NewLabel("makes a new droplet")
	myCreate.grid.NewButton("Create", func () {
		name := myCreate.name.Get()
		log.Info("create droplet name =", name, "zone =", zone.Slug)
		myDo.Create(name, zone.Slug)
	})

	myCreate.ready = true
	myDo.create = myCreate
	return myCreate
}

// Returns true if the status is valid
func (d *windowCreate) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *windowCreate) Show() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Show() window")
	if d.hidden {
		d.window.Show()
	}
	d.hidden = false
}

func (d *windowCreate) Hide() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Hide() window")
	if ! d.hidden {
		d.window.Hide()
	}
	d.hidden = true
}
