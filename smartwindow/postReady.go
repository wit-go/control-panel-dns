// This creates a 'smart window'
// it should work even when it is hidden
// from the gui toolkit plugins
package smartwindow

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

/*
	all these functions run after the window is Ready()
	so they should all start with that check
*/

// reports externally if something has changed
// since the last time it was asked about it
func (sw *SmartWindow) Changed() bool {
	if ! sw.Ready() {return false}

	if sw.changed {
		sw.changed = false
		return true
	}
	return false
}

func (sw *SmartWindow) Show() {
	if ! sw.Ready() {return}

	log.Log(WARN, "Show() window ready =", sw.ready)
	sw.window.Show()
	sw.hidden = false
}

func (sw *SmartWindow) Hide() {
	if ! sw.Ready() {return}

	log.Log(WARN, "Hide() window ready =", sw.ready)
	sw.window.Hide()
	sw.hidden = true
}

func (sw *SmartWindow) Toggle() {
	if ! sw.Ready() {return}

	log.Log(WARN, "Toggle() window ready =", sw.ready)
	if sw.hidden {
		sw.Show()
	} else {
		sw.Hide()
	}
}

func (sw *SmartWindow) Box() *gui.Node {
	if ! sw.Ready() {return nil}

	return sw.box
}

func (sw *SmartWindow) Draw() {
	if ! sw.Ready() {return}

	log.Log(WARN, "Draw() window ready")
	sw.window.Draw()

	if sw.vertical {
		sw.box = sw.window.NewBox("bw vbox", false)
		log.Log(WARN, "BasicWindow.Custom() made vbox")
	} else {
		sw.box = sw.window.NewBox("bw hbox", true)
		log.Log(WARN, "BasicWindow.Custom() made hbox")
	}
	if (sw.populate != nil) {
		log.Log(WARN, "Make() trying to run Custom sw.populate() here")
		sw.populate(sw)
	}
}
