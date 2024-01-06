// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/log"
)

func (ls *LinuxStatus) Show() {
	log.Log(CHANGE, "linuxStatus.Show() window")
	ls.window.Show()
	ls.hidden = false
}

func (ls *LinuxStatus) Hide() {
	log.Log(CHANGE, "linuxStatus.Hide() window")
	ls.window.Hide()
	ls.hidden = true
}

func (ls *LinuxStatus) Toggle() {
	log.Log(CHANGE, "linuxStatus.Toggle() window")
	if ls.hidden {
		ls.window.Show()
	} else {
		ls.window.Hide()
	}
}

func (ls *LinuxStatus) Ready() bool {
	if me == nil {return false}
	if ls == nil {return false}
	if ls.window == nil {return false}
	return me.ready
}

func (ls *LinuxStatus) Initialized() bool {
	if me == nil {return false}
	if ls == nil {return false}
	if ls.parent == nil {return false}
	return true
}
