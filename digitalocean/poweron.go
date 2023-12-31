package digitalocean

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"

	"go.wit.com/log"
)

func (d *DigitalOcean) PowerOn(dropletID int) error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	// Create a request to power on the droplet.
	_, _, err := client.DropletActions.PowerOn(ctx, dropletID)
	if err != nil {
		return err
	}

	log.Printf("Power-on signal sent to droplet with ID: %d\n", dropletID)
	return nil
}

func (d *DigitalOcean) PowerOff(dropletID int) error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	// Create a request to power on the droplet.
	_, _, err := client.DropletActions.PowerOff(ctx, dropletID)
	if err != nil {
		return err
	}

	log.Printf("Power-on signal sent to droplet with ID: %d\n", dropletID)
	return nil
}
