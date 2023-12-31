package digitalocean

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"

	"go.wit.com/log"
)

// func (d *DigitalOcean) ListDroplets() bool {
func (d *DigitalOcean) ListSSHKeyID() error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	// List all keys.
	keys, _, err := client.Keys.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return err
	}

	// Find the key by name.
	for i, key := range keys {
		log.Info("found ssh i =", i)
		log.Info("found ssh key.Name =", key.Name)
		log.Info("found ssh key.Fingerprint =", key.Fingerprint)
		log.Info("found ssh key:", key)
	//	if key.Name == name {
	//		return key.Fingerprint, nil
	//	}
	}

	// return fmt.Errorf("SSH Key not found")
	return nil
}
