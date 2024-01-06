package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"

	"go.wit.com/log"
	"github.com/miekg/dns"
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

func digAAAA(hostname string) []string {
	var blah, ipv6Addresses []string
	// domain := hostname
	recordType := dns.TypeAAAA // dns.TypeTXT

	// Cloudflare's DNS server
	blah, _ = dnsUdpLookup("1.1.1.1:53", hostname, recordType)
	log.Println("digAAAA() has BLAH =", blah)

	if (len(blah) == 0) {
		log.Println("digAAAA() RUNNING dnsAAAAlookupDoH(domain)")
		ipv6Addresses = lookupDoH(hostname, "AAAA")
		log.Println("digAAAA() has ipv6Addresses =", strings.Join(ipv6Addresses, " "))
		for _, addr := range ipv6Addresses {
			log.Println(addr)
		}
		return ipv6Addresses
	}

	// TODO: check digDoH vs blah, if so, then port 53 TCP and/or UDP is broken or blocked
	log.Println("digAAAA() has BLAH =", blah)

	return blah
}
