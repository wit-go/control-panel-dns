package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// dnsLookupDoH performs a DNS lookup for AAAA records over HTTPS.
func dnsLookupDoH(domain string) ([]string, error) {
	var ipv6Addresses []string

	// Construct the URL for a DNS query with Google's DNS-over-HTTPS API
	url := fmt.Sprintf("https://dns.google/resolve?name=%s&type=AAAA", domain)

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

/*
func main() {
	domain := "google.com"
	ipv6Addresses, err := dnsLookupDoH(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("IPv6 Addresses for %s:\n", domain)
	for _, addr := range ipv6Addresses {
		fmt.Println(addr)
	}
}
*/
