// This is a simple example
package cloudflare

import 	(
	"os"
	"log"

	"go.wit.com/gui"
)

// This creates a window
func MakeCloudflareWindow(n *gui.Node) *gui.Node {
	CFdialog.rootGui = n
	var t *gui.Node

	log.Println("buttonWindow() START")

	CFdialog.mainWindow = n.NewWindow("Cloudflare Config")

	// this tab has the master cloudflare API credentials
	makeConfigWindow(CFdialog.mainWindow)

	t = CFdialog.mainWindow.NewTab("Zones")
	vb := t.NewBox("vBox", false)
	g1 := vb.NewGroup("zones")

	// make dropdown list of zones
	CFdialog.zonedrop = g1.NewDropdown("zone")
	CFdialog.zonedrop.AddText("example.org")
	for d, _ := range Config {
		CFdialog.zonedrop.AddText(d)
	}
	CFdialog.zonedrop.AddText("stablesid.org")

	CFdialog.zonedrop.Custom = func () {
		domain := CFdialog.zonedrop.S
		log.Println("custom dropdown() zone (domain name) =", CFdialog.zonedrop.Name, domain)
		if (Config[domain] == nil) {
			log.Println("custom dropdown() Config[domain] = nil for domain =", domain)
			CFdialog.domainWidget.SetText(domain)
			CFdialog.zoneWidget.SetText("")
			CFdialog.authWidget.SetText("")
			CFdialog.emailWidget.SetText("")
		} else {
			log.Println("custom dropdown() a =", domain, Config[domain].ZoneID, Config[domain].Auth, Config[domain].Email)
			CFdialog.domainWidget.SetText(Config[domain].Domain)
			CFdialog.zoneWidget.SetText(Config[domain].ZoneID)
			CFdialog.authWidget.SetText(Config[domain].Auth)
			CFdialog.emailWidget.SetText(Config[domain].Email)
		}
	}

	more := g1.NewGroup("data")
	showCloudflareCredentials(more)

	makeDebugWindow(CFdialog.mainWindow)
	return CFdialog.mainWindow
}

func makeConfigWindow(n *gui.Node) {
	t := n.NewTab("Get Zones")
	vb := t.NewBox("vBox", false)
	g1 := vb.NewGroup("Cloudflare API Config")

	g1.NewLabel("If you have an API key with access to list all of /n your zone files, enter it here. \n \n Alternatively, you can set the enviroment variables: \n env $CF_API_KEY \n env $CF_API_EMAIL\n")

	// make grid to display credentials
	grid := g1.NewGrid("credsGrid", 2, 4) // width = 2

	grid.NewLabel("Auth Key")
	aw := grid.NewEntryLine("CF_API_KEY")
	aw.SetText(os.Getenv("CF_API_KEY"))

	grid.NewLabel("Email")
	ew := grid.NewEntryLine("CF_API_EMAIL")
	ew.SetText(os.Getenv("CF_API_EMAIL"))

	var url string = "https://api.cloudflare.com/client/v4/zones/"
	grid.NewLabel("Cloudflare API")
	grid.NewLabel(url)

	grid.Pad()

	vb.NewButton("getZones()", func () {
		log.Println("getZones()")
		GetZones(aw.S, ew.S)
		for d, _ := range Config {
			CFdialog.zonedrop.AddText(d)
		}
	})

	vb.NewButton("cloudflare wit.com", func () {
		CreateRR(CFdialog.rootGui, "wit.com", "3777302ac4a78cd7fa4f6d3f72086d06")
	})

	t.Pad()
	t.Margin()
	vb.Pad()
	vb.Margin()
	g1.Pad()
	g1.Margin()
}

func makeDebugWindow(window *gui.Node) {
	t2 := window.NewTab("debug")
	g := t2.NewGroup("debug")
	g.NewButton("Load 'gocui'", func () {
		CFdialog.rootGui.LoadToolkit("gocui")
	})

	g.NewButton("Load 'andlabs'", func () {
		CFdialog.rootGui.LoadToolkit("andlabs")
	})

	g.NewButton("gui.DebugWindow()", func () {
		gui.DebugWindow()
	})

	g.NewButton("List all Widgets", func () {
		CFdialog.rootGui.ListChildren(true)
	})
	g.NewButton("Dump all Widgets", func () {
		CFdialog.rootGui.Dump()
	})
}

func showCloudflareCredentials(box *gui.Node) {
	// make grid to display credentials
	grid := box.NewGrid("credsGrid", 2, 4) // width = 2

	grid.NewLabel("Domain")
	CFdialog.domainWidget = grid.NewEntryLine("CF_API_DOMAIN")

	grid.NewLabel("Zone ID")
	CFdialog.zoneWidget = grid.NewEntryLine("CF_API_ZONEID")

	grid.NewLabel("Auth Key")
	CFdialog.authWidget = grid.NewEntryLine("CF_API_KEY")

	grid.NewLabel("Email")
	CFdialog.emailWidget = grid.NewEntryLine("CF_API_EMAIL")

	var url string = "https://api.cloudflare.com/client/v4/zones/"
	grid.NewLabel("Cloudflare API")
	grid.NewLabel(url)

	grid.Pad()

	CFdialog.loadButton = box.NewButton("Load Cloudflare DNS zonefile", func () {
		var domain ConfigT
		domain.Domain = CFdialog.domainWidget.S
		domain.ZoneID = CFdialog.zoneWidget.S
		domain.Auth = CFdialog.authWidget.S
		domain.Email = CFdialog.emailWidget.S
		LoadZoneWindow(CFdialog.mainWindow, &domain)
	})
}
