// This is a simple example
package cloudflare

import 	(
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.wit.com/log"
)

/*
    This function should run each time
    the user chanegs anything in the GUi
    or each time something in general changes
   
    It returns a RR record which then can be
    turned into JSON and sent via http
    to cloudflare's API
*/
func DoChange() *RRT {
	var dnsRow *RRT
	dnsRow = new(RRT)

	log.Println("DoChange() START")
	if (CFdialog.proxyNode.S == "On") {
		dnsRow.Proxied = true
	} else {
		dnsRow.Proxied = false
	}
	dnsRow.Auth = CFdialog.apiNode.S
	dnsRow.Email = CFdialog.emailNode.S

	dnsRow.Domain = CFdialog.zoneNode.S
	dnsRow.ZoneID = CFdialog.zoneIdNode.S
	dnsRow.ID = CFdialog.rrNode.S

	dnsRow.Content = CFdialog.ValueNode.S
	dnsRow.Name = CFdialog.NameNode.S
	dnsRow.Type = CFdialog.TypeNode.S
	dnsRow.url = CFdialog.urlNode.S

	dnsRow.data = makeJSON(dnsRow)
	// show the JSON
	log.Println(dnsRow)

	if (CFdialog.curlNode != nil) {
		pretty, _ := FormatJSON(dnsRow.data)
		log.Println("http PUT curl =", pretty)
		CFdialog.curlNode.SetText(pretty)
	}
	return dnsRow
}

func SetRow(dnsRow *RRT) {
	log.Println("Look for changes in row", dnsRow.ID)
	if (CFdialog.proxyNode != nil) {
		log.Println("Proxy", dnsRow.Proxied, "vs", CFdialog.proxyNode.S)
		if (dnsRow.Proxied == true) {
			CFdialog.proxyNode.SetText("On")
		} else {
			CFdialog.proxyNode.SetText("Off")
		}
	}
	if (CFdialog.zoneNode != nil) {
		CFdialog.zoneNode.SetText(dnsRow.Domain)
	}
	if (CFdialog.zoneIdNode != nil) {
		CFdialog.zoneIdNode.SetText(dnsRow.ZoneID)
	}
	log.Println("zoneIdNode =", dnsRow.ZoneID)
	if (CFdialog.rrNode != nil) {
		CFdialog.rrNode.SetText(dnsRow.ID)
	}
	if (CFdialog.ValueNode != nil) {
		log.Println("Content", dnsRow.Content, "vs", CFdialog.ValueNode.S)
		CFdialog.ValueNode.SetText(dnsRow.Content)
	}
	if (CFdialog.NameNode != nil) {
		CFdialog.NameNode.SetText(dnsRow.Name)
	}
	if (CFdialog.TypeNode != nil) {
		CFdialog.TypeNode.SetText(dnsRow.Type)
	}

	if (CFdialog.urlNode != nil) {
		url := cloudflareURL + dnsRow.ZoneID + "/dns_records/" + dnsRow.ID
		CFdialog.urlNode.SetText(url)
	}

	// show the JSON
	tmp := makeJSON(dnsRow)
	log.Println(tmp)
	if (CFdialog.curlNode != nil) {
		pretty, _ := FormatJSON(tmp)
		log.Println("http PUT curl =", pretty)
		CFdialog.curlNode.SetText(pretty)
	}
}

func GetZonefile(c *ConfigT) *DNSRecords {
	var url = cloudflareURL + c.ZoneID + "/dns_records/"
	log.Println("getZonefile()", c.Domain, url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http.NewRequest error:", err)
		return nil
	}

	// Set headers
	req.Header.Set("X-Auth-Key", c.Auth)
	req.Header.Set("X-Auth-Email", c.Email)

	log.Println("getZonefile() auth, email", c.Auth, c.Email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("http.Client error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() error", err)
		return nil
	}

	var records DNSRecords
	if err := json.Unmarshal(body, &records); err != nil {
		log.Println("json.Unmarshal() error", err)
		return nil
	}

	log.Println("getZonefile() worked", records)
	return &records
}

/*
	pass in a DNS Resource Records (the stuff in a zonefile)

	This will talk to the cloudflare API and generate a resource record in the zonefile:

	For example:
	gitea.wit.com. 3600 IN CNAME git.wit.com.
	go.wit.com. 3600 IN A 1.1.1.9
	test.wit.com. 3600 IN NS ns1.wit.com.
*/
func makeJSON(dnsRow *RRT) string {
	// make a json record to send on port 80 to cloudflare
	var tmp string
	tmp = `{"content": "` + dnsRow.Content + `", `
	tmp += `"name": "` + dnsRow.Name + `", `
	tmp += `"type": "` + dnsRow.Type + `", `
	tmp+= `"ttl": "` +  "1" + `", `
	tmp += `"comment": "WIT DNS Control Panel"`
	tmp +=  `}`

	return tmp
}

// https://api.cloudflare.com/client/v4/zones
func GetZones(auth, email string) *DNSRecords {
	var url = "https://api.cloudflare.com/client/v4/zones"
	log.Println("getZones()", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http.NewRequest error:", err)
		return nil
	}

	// Set headers
	req.Header.Set("X-Auth-Key", auth)
	req.Header.Set("X-Auth-Email", email)

	log.Println("getZones() auth, email", auth, email)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("getZones() http.Client error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("getZones() ioutil.ReadAll() error", err)
		return nil
	}

	var records DNSRecords
	if err := json.Unmarshal(body, &records); err != nil {
		log.Println("getZones() json.Unmarshal() error", err)
		return nil
	}

	/* Cloudflare API returns struct[] of:
	  struct { ID string "json:\"id\""; Type string "json:\"type\""; Name string "json:\"name\"";
		Content string "json:\"content\""; Proxied bool "json:\"proxied\"";
		Proxiable bool "json:\"proxiable\""; TTL int "json:\"ttl\"" }
	*/

	// log.Println("getZones() worked", records)
	// log.Println("spew dump:")
	// spew.Dump(records)
	for _, record := range records.Result {
		log.Println("spew record:", record)
		log.Println("record:", record.Name, record.ID)

		var newc *ConfigT
		newc = new(ConfigT)

		newc.Domain = record.Name
		newc.ZoneID = record.ID
		newc.Auth = auth
		newc.Email = email

		Config[record.Name] = newc
		log.Println("zonedrop.AddText:", record.Name, record.ID)
	}
	for d, _ := range Config {
		log.Println("Config entry:", d)
	}

	return &records
}
