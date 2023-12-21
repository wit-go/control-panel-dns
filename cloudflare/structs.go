// This is a simple example
package cloudflare

import 	(
	"go.wit.com/gui"
)

var cloudflareURL string = "https://api.cloudflare.com/client/v4/zones/"

// Define a struct to match the JSON structure of the response.
// This structure should be adjusted based on the actual format of the response.
type DNSRecords struct {
	Result []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Content string `json:"content"`
		Proxied bool `json:"proxied"`
		Proxiable bool `json:"proxiable"`
		TTL int `json:"ttl"`
	} `json:"result"`
}

// CFdialog is everything you need forcreating 
// a new record: name, TTL, type (CNAME, A, etc)
var CFdialog dialogT

type dialogT struct {
	rootGui *gui.Node	// the root node
	mainWindow *gui.Node	// the window node
	zonedrop *gui.Node	// the drop down menu of zones

	domainWidget *gui.Node
	zoneWidget *gui.Node
	authWidget *gui.Node
	emailWidget *gui.Node

	loadButton *gui.Node
	saveButton *gui.Node

	cloudflareW *gui.Node	// the window node
	cloudflareB *gui.Node	// the cloudflare button

	TypeNode *gui.Node	// CNAME, A, AAAA, ...
	NameNode *gui.Node	// www, mail, ...
	ValueNode *gui.Node	// 4.2.2.2, "dkim stuff", etc

	rrNode *gui.Node	// cloudflare Resource Record ID
	proxyNode *gui.Node	// If cloudflare is a port 80 & 443 proxy
	ttlNode *gui.Node	// just set to 1 which means automatic to cloudflare
	curlNode *gui.Node	// shows you what you could run via curl
	resultNode *gui.Node	// what the cloudflare API returned
	saveNode *gui.Node	// button to send it to cloudflare

	zoneNode *gui.Node	// "wit.com"
	zoneIdNode *gui.Node	// cloudflare zone ID
	apiNode *gui.Node	// cloudflare API key (from environment var CF_API_KEY)
	emailNode *gui.Node	// cloudflare email (from environment var CF_API_EMAIL)
	urlNode *gui.Node	// the URL to POST, PUT, DELETE, etc
}

// Resource Record (used in a DNS zonefile)
type RRT struct {
	ID     string
	Type   string
	Name   string
	Content string
	ProxyS string
	Proxied bool
	Proxiable bool
	Ttl string

	Domain string
	ZoneID string
	Auth string
	Email string
	url string
	data string
}

/*
	This is a structure of all the RR's (Resource Records)
	in the DNS zonefiile for a hostname. For example:

	For the host test.wit.com:

	test.wit.com A 127.0.0.1
	test.wit.com AAAA
	test.wit.com TXT email test@wit.com
	test.wit.com TXT phone 212-555-1212
	test.wit.com CNAME real.wit.com
*/
type hostT struct {
	hostname string
	RRs []ConfigT
}

type ConfigT struct {
	Domain string
	ZoneID string
	Auth string
	Email string
}

var Config map[string]*ConfigT