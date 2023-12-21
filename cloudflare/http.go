// This is a simple example
package cloudflare

import 	(
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

/*
curl --request POST \
  --url https://api.cloudflare.com/client/v4/zones/zone_identifier/dns_records \
  --header 'Content-Type: application/json' \
  --header 'X-Auth-Email: ' \
  --data '{
  "content": "198.51.100.4",
  "name": "example.com",
  "proxied": false,
  "type": "A",
  "comment": "Domain verification record",
  "tags": [
    "owner:dns-team"
  ],
  "ttl": 3600
}'
*/

func httpPut(dnsRow *RRT) (string, string) {
	var url string = cloudflareURL + dnsRow.ZoneID + "/dns_records/" + dnsRow.ID
	var authKey string = dnsRow.Auth
	var email string = dnsRow.Email

	var tmp string
	tmp = makeJSON(dnsRow)
	data := []byte(tmp)

	log.Println("http PUT url =", url)
	// log.Println("http PUT data =", data)
	// spew.Dump(data)
	pretty, _ := FormatJSON(string(data))
	log.Println("http PUT data =", pretty)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Key", authKey)
	req.Header.Set("X-Auth-Email", email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return tmp, fmt.Sprintf("blah err =", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return tmp, fmt.Sprintf("blah err =", err)
	}
	// log.Println("http PUT body =", body)
	// spew.Dump(body)

	return tmp, string(body)
}

func curlPost(dnsRow *RRT) string {
	var authKey string = dnsRow.Auth
	var email string = dnsRow.Email

	url := dnsRow.url
	tmp := dnsRow.data

	log.Println("curl() START")
	log.Println("curl() authkey = ", authKey)
	log.Println("curl() email   = ", email)
	log.Println("curl() url     = ", url)
	data := []byte(tmp)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Key", authKey)
	req.Header.Set("X-Auth-Email", email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	// log.Println("http PUT body =", body)
	// spew.Dump(body)

	log.Println("result =", string(body))
	log.Println("curl() END")
	pretty, _ := FormatJSON(string(body))
	return pretty
}
