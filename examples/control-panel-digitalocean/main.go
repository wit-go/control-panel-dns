package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"go.wit.com/log"
	"go.wit.com/gui"
	"github.com/digitalocean/godo"
	"go.wit.com/control-panel-dns/digitalocean"
)

var title string = "Digital Ocean Control Panel"

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

func main() {
	// Your personal API token from DigitalOcean.
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		log.Fatal("Please set your DigitalOcean API token in the DIGITALOCEAN_TOKEN environment variable")
	}

	// List droplets and their IP addresses.
	err := digitalocean.ListDroplets(token)
	if err != nil {
		log.Fatalf("Error listing droplets: %s\n", err)
	}

	// initialize a new GO GUI instance
	myGui := gui.New().Default()

	// draw the cloudflare control panel window
	win := digitalocean.MakeWindow(myGui)
	win.SetText(title)

	// This is just a optional goroutine to watch that things are alive
	gui.Watchdog()
	gui.StandardExit()

	os.Exit(0)

	// Parameters for the droplet you wish to create.
	name := "ipv6.wit.com"
	region := "nyc1" // New York City region.
	size := "s-1vcpu-1gb" // Size of the droplet.
	image := "ubuntu-20-04-x64" // Image slug for Ubuntu 20.04 (LTS) x64.

	// Create a new droplet.
	droplet, err := createDropletNew(token, name, region, size, image)
	if err != nil {
		log.Fatalf("Something went wrong: %s\n", err)
	}

	fmt.Printf("Created droplet ID %d with name %s\n", droplet.ID, droplet.Name)
}

// createDroplet creates a new droplet in the specified region with the given name.
func createDropletNew(token, name, region, size, image string) (*godo.Droplet, error) {
	// Create an OAuth2 token.
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

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
		IPv6: true, // Enable IPv6
	}

	// Create the droplet.
	ctx := context.TODO()
	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}

	return newDroplet, nil
}