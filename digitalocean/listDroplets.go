package digitalocean

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"
)

// ListDroplets fetches and prints out the droplets along with their IPv4 and IPv6 addresses.
func ListDroplets(token string) error {
	// OAuth token for authentication.
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	// OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// DigitalOcean client.
	client := godo.NewClient(oauthClient)

	// Context.
	ctx := context.TODO()

	// List all droplets.
	droplets, _, err := client.Droplets.List(ctx, &godo.ListOptions{})
	if err != nil {
		return err
	}

	// Iterate over droplets and print their details.
	for _, droplet := range droplets {
		fmt.Printf("Droplet: %s\n", droplet.Name)
		for _, network := range droplet.Networks.V4 {
			if network.Type == "public" {
				fmt.Printf("IPv4: %s\n", network.IPAddress)
			}
		}
		for _, network := range droplet.Networks.V6 {
			if network.Type == "public" {
				fmt.Printf("IPv6: %s\n", network.IPAddress)
			}
		}
		fmt.Println("-------------------------")
	}

	return nil
}
