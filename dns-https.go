package main

import (
	"fmt"
	"go.wit.com/log"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

/*
func getAAAArecords() {
	hostname := "go.wit.com"
	ipv6Addresses, err := dnsLookupDoH(hostname)
	if err != nil {
		log.Error(err, "getAAAArecords")
		return
	}

	fmt.Printf("IPv6 Addresses for %s:\n", hostname)
	for _, addr := range ipv6Addresses {
		log.Println(addr)
	}
}
*/

// dnsLookupDoH performs a DNS lookup for AAAA records over HTTPS.
func dnsAAAAlookupDoH(domain string) ([]string, error) {
	var ipv6Addresses []string

	// Construct the URL for a DNS query with Google's DNS-over-HTTPS API
	url := fmt.Sprintf("https://dns.google/resolve?name=%s&type=AAAA", domain)

	log.Println("curl", url)

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error performing DNS-over-HTTPS request: %w", err)
	}
	defer resp.Body.Close()

	// Read and unmarshal the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var data struct {
		Answer []struct {
			Data string `json:"data"`
		} `json:"Answer"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Extract the IPv6 addresses
	for _, answer := range data.Answer {
		ipv6Addresses = append(ipv6Addresses, answer.Data)
	}

	return ipv6Addresses, nil
}
