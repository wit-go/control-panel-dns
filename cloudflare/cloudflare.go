// This is a simple example
package cloudflare

import 	(
	"log"
	"os"
	"bytes"
	"io/ioutil"
	"net/http"

	"git.wit.org/wit/gui"
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

// CFdialog is everything you need forcreating 
// a new record: name, TTL, type (CNAME, A, etc)
var CFdialog RRT

// Resource Record (used in a DNS zonefile)
type RRT struct {
	cloudflareW *gui.Node	// the window node
	cloudflareB *gui.Node	// the cloudflare button

	TypeNode *gui.Node	// CNAME, A, AAAA, ...
	NameNode *gui.Node	// www, mail, ...
	ValueNode *gui.Node	// 4.2.2.2, "dkim stuff", etc

	proxyNode *gui.Node	// If cloudflare is a port 80 & 443 proxy
	ttlNode *gui.Node	// just set to 1 which means automatic to cloudflare
	curlNode *gui.Node	// shows you what you could run via curl
	resultNode *gui.Node	// what the cloudflare API returned
	saveNode *gui.Node	// button to send it to cloudflare

	zoneNode *gui.Node	// "wit.com"
	zoneIdNode *gui.Node	// cloudflare zone ID
	apiNode *gui.Node	// cloudflare API key (from environment var CF_API_KEY)
	emailNode *gui.Node	// cloudflare email (from environment var CF_API_EMAIL)

	ID     string
	Type   string
	Name   string
	Content string
	ProxyS string
	Proxied bool
	Proxiable bool
	Ttl string
}

func CreateRR(myGui *gui.Node, zone string, zoneID string) {
	if (CFdialog.cloudflareW != nil) {
		// skip this if the window has already been created
		log.Println("createRR() the cloudflare window already exists")
		CFdialog.cloudflareB.Disable()
		return
	}
	CFdialog.cloudflareW = myGui.NewWindow("cloudflare " + zone + " API")
	CFdialog.cloudflareW.Custom = func () {
		log.Println("createRR() don't really exit here")
		CFdialog.cloudflareW = nil
		CFdialog.cloudflareB.Enable()
	}

	CFdialog.ID = zoneID

	group := CFdialog.cloudflareW.NewGroup("Create a new DNS Resource Record (rr)")

	// make a grid 2 things wide
	grid := group.NewGrid("gridnuts", 2, 3)

	grid.NewLabel("zone")
	CFdialog.zoneNode = grid.NewLabel("zone")
	CFdialog.zoneNode.SetText(zone)

	grid.NewLabel("zone ID")
	CFdialog.zoneIdNode = grid.NewLabel("zoneID")
	CFdialog.zoneIdNode.SetText(zoneID)

	grid.NewLabel("shell env $CF_API_EMAIL")
	CFdialog.emailNode = grid.NewLabel("type")
	CFdialog.emailNode.SetText(os.Getenv("CF_API_EMAIL"))

	grid.NewLabel("shell env $CF_API_KEY")
	CFdialog.apiNode = grid.NewLabel("type")
	CFdialog.apiNode.SetText(os.Getenv("CF_API_KEY"))

	grid.NewLabel("Record Type")
	CFdialog.TypeNode = grid.NewCombobox("type")
	CFdialog.TypeNode.AddText("A")
	CFdialog.TypeNode.AddText("AAAA")
	CFdialog.TypeNode.AddText("CNAME")
	CFdialog.TypeNode.AddText("TXT")
	CFdialog.TypeNode.AddText("MX")
	CFdialog.TypeNode.AddText("NS")
	CFdialog.TypeNode.Custom = func () {
		CreateCurlRR()
	}
	CFdialog.TypeNode.SetText("AAAA")

	grid.NewLabel("Name (usually the hostname)")
	CFdialog.NameNode = grid.NewCombobox("name")
	CFdialog.NameNode.AddText("www")
	CFdialog.NameNode.AddText("mail")
	CFdialog.NameNode.AddText("git")
	CFdialog.NameNode.AddText("go")
	CFdialog.NameNode.AddText("blog")
	CFdialog.NameNode.AddText("ns1")
	CFdialog.NameNode.Custom = func () {
		CreateCurlRR()
	}
	CFdialog.NameNode.SetText("www")

	grid.NewLabel("Cloudflare Proxy")
	CFdialog.proxyNode = grid.NewDropdown("proxy")
	CFdialog.proxyNode.AddText("On")
	CFdialog.proxyNode.AddText("Off")
	CFdialog.proxyNode.Custom = func () {
		CreateCurlRR()
	}
	CFdialog.proxyNode.SetText("Off")

	grid.NewLabel("Value")
	CFdialog.ValueNode = grid.NewCombobox("value")
	CFdialog.ValueNode.AddText("127.0.0.1")
	CFdialog.ValueNode.AddText("2001:4860:4860::8888")
	CFdialog.ValueNode.AddText("ipv6.wit.com")
	CFdialog.ValueNode.Custom = func () {
		CreateCurlRR()
	}
	CFdialog.ValueNode.SetText("127.0.0.1")
	CFdialog.ValueNode.Expand()

	group.NewLabel("curl")
	CFdialog.curlNode = group.NewTextbox("curl")
	CFdialog.curlNode.Custom = func () {
		CreateCurlRR()
	}
	CFdialog.curlNode.SetText("put the curl text here")

	CFdialog.resultNode = group.NewTextbox("result")
	CFdialog.resultNode.SetText("API response will show here")

	CFdialog.saveNode = group.NewButton("Save", func () {
		url, data := CreateCurlRR()
		result := curl(url, data)
		CFdialog.resultNode.SetText(result)
	})
	CFdialog.saveNode.Disable()

	group.Pad()
	grid.Pad()
	grid.Expand()
}

func CreateCurlRR() (string, string) {
	// enable the Save/Create Button
	if (CFdialog.saveNode != nil) {
		CFdialog.saveNode.Enable()
	}

	if (CFdialog.TypeNode != nil) {
		CFdialog.Type = CFdialog.TypeNode.S
	}
	if (CFdialog.NameNode != nil) {
		CFdialog.Name = CFdialog.NameNode.S
	}
	if (CFdialog.proxyNode != nil) {
		if (CFdialog.proxyNode.S == "On") {
			CFdialog.ProxyS = "true"
		} else {
			CFdialog.ProxyS = "false"
		}
	}
	if (CFdialog.ValueNode != nil) {
		CFdialog.Content = CFdialog.ValueNode.S
	}
	CFdialog.Ttl = "3600"

	var url string = "https://api.cloudflare.com/client/v4/zones/" + CFdialog.ID + "/dns_records"
	// https://api.cloudflare.com/client/v4/zones/zone_identifier/dns_records \
	// var authKey string = os.Getenv("CF_API_KEY")
	// var email string = os.Getenv("CF_API_EMAIL")

	// make a json record to send on port 80 to cloudflare
	var tmp string
	tmp = `{"content": "` + CFdialog.Content + `", `
	tmp += `"name": "` + CFdialog.Name + `", `
	tmp += `"type": "` + CFdialog.Type + `", `
	tmp += `"ttl": ` +  CFdialog.Ttl + `, `
	tmp += `"proxied": ` +  CFdialog.ProxyS + `, `
	tmp += `"comment": "WIT DNS Control Panel"`
	tmp +=  `}`
	data := []byte(tmp)

	log.Println("http PUT url =", url)
	// log.Println("http PUT data =", data)
	// spew.Dump(data)
	pretty, _ := formatJSON(string(data))
	log.Println("http URL =", url)
	log.Println("http PUT data =", pretty)
	if (CFdialog.curlNode != nil) {
		CFdialog.curlNode.SetText("URL: " + url + "\n" + pretty)
	}

	return url, tmp
}

func curl(url string, tmp string) string {
	var authKey string = CFdialog.apiNode.S
	var email string = CFdialog.emailNode.S

	log.Println("curl() START")
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
	pretty, _ := formatJSON(string(body))
	return pretty
}
