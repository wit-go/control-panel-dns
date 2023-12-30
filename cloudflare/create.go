/*
	This will attempt to create a RR in a DNS zone file.

	Create("wit.com", "test.wit.com", "1.1.1.1"
*/

package cloudflare

import 	(
	"os"

	"go.wit.com/log"
)

func Create(zone string, hostname string, value string) bool {
	log.Info("cloudflare.Create() START", zone, hostname, value)
	key := os.Getenv("CF_API_KEY")
	email := os.Getenv("CF_API_EMAIL")

	if (key == "") {
		log.Warn("cloudflare.Create() MISSING environment variable CF_API_KEY")
		return false
	}
	if (email == "") {
		log.Warn("cloudflare.Create() MISSING environment variable CF_API_EMAIL")
		return false
	}

	GetZones(key, email)
	var z *ConfigT
	for d, v := range Config {
		log.Info("cloudflare.Create() zone =", d, "value =", v)
		if (zone == d) {
			z = Config[zone]
			log.Info("cloudflare.Create() FOUND ZONE", zone, "ID =", z.ZoneID)
		}
	}
	if (z == nil) {
		log.Warn("cloudflare.Create() COULD NOT FIND ZONE", zone)
		return false
	}
	log.Info("cloudflare.Create() FOUND ZONE", z)

	// make a json record to send on port 80 to cloudflare
	var data string
	data = `{"content": "` + value + `", `
	data += `"name": "` + hostname + `", `
	data += `"type": "AAAA", `
	data += `"ttl": "1", `
	data += `"comment": "WIT DNS Control Panel"`
	data +=  `}`

	result := doCurlCreate(key, email, z.ZoneID, data)
	pretty, _ := FormatJSON(result)
	log.Info("cloudflare.Create() result =", pretty)
	return true
}
