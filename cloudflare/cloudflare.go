// This is a simple example
package cloudflare

import 	(
	"log"
	"os"

	"go.wit.com/gui"
)

func init() {
	Config = make(map[string]*ConfigT)
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

	grid.NewLabel("Resource Record ID")
	CFdialog.rrNode = grid.NewLabel("type")
	CFdialog.rrNode.SetText(os.Getenv("cloudflare RR id"))

	grid.NewLabel("Record Type")
	CFdialog.TypeNode = grid.NewCombobox("type")
	CFdialog.TypeNode.AddText("A")
	CFdialog.TypeNode.AddText("AAAA")
	CFdialog.TypeNode.AddText("CNAME")
	CFdialog.TypeNode.AddText("TXT")
	CFdialog.TypeNode.AddText("MX")
	CFdialog.TypeNode.AddText("NS")
	CFdialog.TypeNode.Custom = func () {
		DoChange()
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
		DoChange()
	}
	CFdialog.NameNode.SetText("www")

	grid.NewLabel("Cloudflare Proxy")
	CFdialog.proxyNode = grid.NewDropdown("proxy")
	CFdialog.proxyNode.AddText("On")
	CFdialog.proxyNode.AddText("Off")
	CFdialog.proxyNode.Custom = func () {
		DoChange()
	}
	CFdialog.proxyNode.SetText("Off")

	grid.NewLabel("Value")
	CFdialog.ValueNode = grid.NewCombobox("value")
	CFdialog.ValueNode.AddText("127.0.0.1")
	CFdialog.ValueNode.AddText("2001:4860:4860::8888")
	CFdialog.ValueNode.AddText("ipv6.wit.com")
	CFdialog.ValueNode.Custom = func () {
		DoChange()
	}
	CFdialog.ValueNode.SetText("127.0.0.1")
	CFdialog.ValueNode.Expand()

	grid.NewLabel("URL")
	CFdialog.urlNode = grid.NewLabel("URL")

	group.NewLabel("curl")
	CFdialog.curlNode = group.NewTextbox("curl")
	CFdialog.curlNode.Custom = func () {
		DoChange()
	}
	CFdialog.curlNode.SetText("put the curl text here")

	CFdialog.resultNode = group.NewTextbox("result")
	CFdialog.resultNode.SetText("API response will show here")

	CFdialog.saveNode = group.NewButton("Save", func () {
		dnsRow := DoChange()
		result := curlPost(dnsRow)
		CFdialog.resultNode.SetText(result)
		// CreateCurlRR()
		// url, data := CreateCurlRR()
		// result := curl(url, data)
		// CFdialog.resultNode.SetText(result)
	})
	// CFdialog.saveNode.Disable()

	group.Pad()
	grid.Pad()
	grid.Expand()
}

/*
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
	pretty, _ := FormatJSON(string(data))
	log.Println("http URL =", url)
	log.Println("http PUT data =", pretty)
	if (CFdialog.curlNode != nil) {
		CFdialog.curlNode.SetText("URL: " + url + "\n" + pretty)
	}

	return url, tmp
}
*/
