package main

import (
	"fmt"
	"go.wit.com/log"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

// dnsLookupDoH performs a DNS lookup for AAAA records over HTTPS.
func lookupDoH(hostname string, rrType string) []string {
	var values []string

	// Construct the URL for a DNS query with Google's DNS-over-HTTPS API
	url := fmt.Sprintf("https://dns.google.com/resolve?name=%s&type=%s", hostname, rrType)

	log.Log(DNS, "lookupDoH()", url)
	if hostname == "" {
		log.Warn("lookupDoH() was sent a empty hostname")
		return nil
	}

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err, "error performing DNS-over-HTTPS request")
		return nil
	}
	defer resp.Body.Close()

	// Read and unmarshal the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Errorf("error reading response: %w", err))
		return nil
	}

	var data struct {
		Answer []struct {
			Data string `json:"data"`
		} `json:"Answer"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Error(fmt.Errorf("error unmarshaling response: %w", err))
		return nil
	}

	// Extract the IPv6 addresses
	for _, answer := range data.Answer {
		values = append(values, answer.Data)
	}

	return values
}
