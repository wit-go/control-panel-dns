package digitalocean

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"
)

func GetSSHKeyID(token, name string) (string, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	// List all keys.
	keys, _, err := client.Keys.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return "", err
	}

	// Find the key by name.
	for _, key := range keys {
		if key.Name == name {
			return key.Fingerprint, nil
		}
	}

	return "", fmt.Errorf("SSH Key not found")
}
