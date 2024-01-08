// testing the smart window
package main

import 	(
	"go.wit.com/log"
	// "go.wit.com/gui/cloudflare"
	"go.wit.com/apps/control-panel-dns/smartwindow"
)

func testSmartWindow(sw *smartwindow.SmartWindow) {
	log.Log(WARN, "testSmartWindow() START")
	grid := sw.Box().NewGrid("test", 5, 1)
	grid.NewLabel("test 1")
	grid.NewLabel("test 1")
	grid.NewLabel("test 2")
	grid.NewLabel("test 2")
	grid.NewLabel("test 3")
	grid.NewLabel("test 3")
	grid.NewLabel("test 3")
	grid.NewLabel("test 3")
}
