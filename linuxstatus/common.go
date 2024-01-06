// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

// reports externally if something has changed
// since the last time it was asked about it
func (ls *LinuxStatus) Changed() bool {
	if ! ls.Ready() {return false}
	if ls.changed {
		ls.changed = false
		return true
	}
	return false
}

func (ls *LinuxStatus) Make() {
	if ! ls.Ready() {return}
	log.Log(CHANGE, "Make() window ready =", ls.ready)
	ls.window.Make()
	ls.ready = true
}
func (ls *LinuxStatus) Draw() {
	if ! ls.Ready() {return}
	log.Log(CHANGE, "Draw() window ready =", ls.ready)
	ls.window.Draw()
	ls.ready = true
}
func (ls *LinuxStatus) Draw2() {
	if ! ls.Ready() {return}
	log.Log(CHANGE, "draw(ls) ready =", ls.ready)
	draw(ls)
}

func (ls *LinuxStatus) Show() {
	if ! ls.Ready() {return}
	log.Log(CHANGE, "Show() window ready =", ls.ready)
	ls.window.Show()
	ls.hidden = false
}

func (ls *LinuxStatus) Hide() {
	if ! ls.Ready() {return}
	log.Log(CHANGE, "Hide() window ready =", ls.ready)
	ls.window.Hide()
	ls.hidden = true
}

func (ls *LinuxStatus) Toggle() {
	if ! ls.Ready() {return}
	log.Log(CHANGE, "Toggle() window ready =", ls.ready)
	if ls.hidden {
		ls.Show()
	} else {
		ls.Hide()
	}
}

func (ls *LinuxStatus) Ready() bool {
	log.Log(SPEW, "Ready() maybe not ready? ls =", ls)
	if me == nil {return false}
	if ls == nil {return false}
	if ls.window == nil {return false}
	return me.ready
}

func (ls *LinuxStatus) Initialized() bool {
	log.Log(CHANGE, "checking Initialized()")
	if me == nil {return false}
	if ls == nil {return false}
	if ls.parent == nil {return false}
	return true
}

func (ls *LinuxStatus) SetParent(p *gui.Node) {
	log.Log(CHANGE, "Attempting SetParent")
	if me == nil {return}
	if ls == nil {return}
	if ls.parent == nil {
		log.Log(CHANGE, "SetParent =", p)
		ls.parent = p
		return
	} else {
		log.Log(CHANGE, "SetParent was already set to =", ls.parent)
	}
}
