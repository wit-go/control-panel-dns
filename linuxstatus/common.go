// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

func (ls *LinuxStatus) Draw() {
	log.Log(CHANGE, "linuxStatus.Draw() window")
	if ! ls.Ready() {return}
	log.Log(CHANGE, "linuxStatus.Draw() window ready =", ls.ready)
	ls.window.Draw()
	ls.ready = true
}
func (ls *LinuxStatus) Draw2() {
	log.Log(CHANGE, "draw(ls)")
	if ! ls.Ready() {return}
	log.Log(CHANGE, "draw(ls) ready =", ls.ready)
	draw(ls)
}

func (ls *LinuxStatus) Show() {
	log.Log(CHANGE, "linuxStatus.Show() window")
	if ! ls.Ready() {return}
	log.Log(CHANGE, "linuxStatus.Show() window ready =", ls.ready)
	ls.window.Show()
	ls.hidden = false
}

func (ls *LinuxStatus) Hide() {
	log.Log(CHANGE, "linuxStatus.Hide() window")
	if ! ls.Ready() {return}
	log.Log(CHANGE, "linuxStatus.Hide() window ready =", ls.ready)
	ls.window.Hide()
	ls.hidden = true
}

func (ls *LinuxStatus) Toggle() {
	log.Log(CHANGE, "linuxStatus.Toggle() window")
	if ! ls.Ready() {return}
	log.Log(CHANGE, "linuxStatus.Toggle() window ready =", ls.ready)
	if ls.hidden {
		ls.Show()
	} else {
		ls.Hide()
	}
}

func (ls *LinuxStatus) Ready() bool {
	log.Log(CHANGE, "Ready() ls =", ls)
	if me == nil {return false}
	if ls == nil {return false}
	if ls.window == nil {return false}
	return me.ready
}

func (ls *LinuxStatus) Initialized() bool {
	log.Log(CHANGE, "checking Initialized() ls =", ls)
	if me == nil {return false}
	if ls == nil {return false}
	if ls.parent == nil {return false}
	return true
}

func (ls *LinuxStatus) SetParent(p *gui.Node) {
	log.Log(CHANGE, "Attempting SetParent =", p)
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
