/*
	This will attempt to delete a RR in a DNS zone file.

	Delete("wit.com", "test.wit.com", "1.1.1.1"
*/

package cloudflare

import 	(
	"os"

	"go.wit.com/log"
)

func Delete(zone string, hostname string, value string) bool {
	// CFdialog.emailNode.SetText(os.Getenv("CF_API_EMAIL"))
	// CFdialog.apiNode.SetText(os.Getenv("CF_API_KEY"))

	log.Info("cloudflare.Delete() START", zone, hostname, value)
	key := os.Getenv("CF_API_KEY")
	email := os.Getenv("CF_API_EMAIL")

	if (key == "") {
		log.Warn("cloudflare.Delete() MISSING environment variable CF_API_KEY")
		return false
	}
	if (email == "") {
		log.Warn("cloudflare.Delete() MISSING environment variable CF_API_EMAIL")
		return false
	}

	GetZones(key, email)
	var z *ConfigT
	for d, v := range Config {
		log.Info("cloudflare.Delete() zone =", d, "value =", v)
		if (zone == d) {
			z = Config[zone]
			log.Info("cloudflare.Delete() FOUND ZONE", zone, "ID =", z.ZoneID)
		}
	}
	if (z == nil) {
		log.Warn("cloudflare.Delete() COULD NOT FIND ZONE", zone)
		return false
	}
	log.Info("cloudflare.Delete() FOUND ZONE", z)

	records := GetZonefile(z)
	for i, record := range records.Result {
		if (record.Name == hostname) {
			log.Info("cloudflare.Delete() FOUND hostname:", i, record.ID, record.Type, record.Name, record.Content)
		}
		if (record.Content == value) {
			log.Info("cloudflare.Delete() FOUND CONTENT:", i, record.ID, record.Type, record.Name, record.Content)
			log.Info("cloudflare.Delete() DO THE ACTUAL cloudflare DELETE here")
			result := doCurlDelete(key, email, z.ZoneID, record.ID)
			pretty, _ := FormatJSON(result)
			log.Info("cloudflare.Delete() result =", pretty)
			return true
		}
	}

	log.Info("cloudflare.Delete() NEVER FOUND cloudflare value:", value)
	return false
}
