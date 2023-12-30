// This is a simple example
package cloudflare

import 	(
	"io/ioutil"
	"net/http"
	"bytes"

	"go.wit.com/log"
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

func doCurlDelete(auth string, email string, zoneId string, rrId string) string {
	var err error
	var req *http.Request

	if zoneId == "" {
		log.Warn("doCurlDelete() zoneId == nil")
		return ""
	}

	if rrId == "" {
		log.Warn("doCurlDelete() rrId == nil")
		return ""
	}

	data := []byte("")

	url := "https://api.cloudflare.com/client/v4/zones/" + zoneId + "/dns_records/" + rrId

	req, err = http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(data))

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Key", auth)
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

	return string(body)
}

func doCurlCreate(auth string, email string, zoneId string, data string) string {
	var err error
	var req *http.Request

	if zoneId == "" {
		log.Warn("doCurlDelete() zoneId == nil")
		return ""
	}

	url := "https://api.cloudflare.com/client/v4/zones/" + zoneId + "/dns_records/"

	log.Info("doCurlCreate() POST url =", url)
	log.Info("doCurlCreate() POST Auth =", auth)
	log.Info("doCurlCreate() POST Email =", email)
	log.Info("doCurlCreate() POST data =", data)

	req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer( []byte(data) ))

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Key", auth)
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

	return string(body)
}

func doCurl(method string, rr *RRT) string {
	var err error
	var req *http.Request

	data := []byte(rr.data)

	if (method == "PUT") {
		req, err = http.NewRequest(http.MethodPut, rr.url, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(http.MethodPost, rr.url, bytes.NewBuffer(data))
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Key", rr.Auth)
	req.Header.Set("X-Auth-Email", rr.Email)

	log.Println("http PUT url =", rr.url)
	log.Println("http PUT Auth =", rr.Auth)
	log.Println("http PUT Email =", rr.Email)
	log.Println("http PUT data =", rr.data)

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

	return string(body)
}

func curlPost(dnsRow *RRT) string {
	var authKey string = dnsRow.Auth
	var email string = dnsRow.Email

	url := dnsRow.url
	tmp := dnsRow.data

	log.Println("curlPost() START")
	log.Println("curlPost() authkey = ", authKey)
	log.Println("curlPost() email   = ", email)
	log.Println("curlPost() url     = ", url)
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
