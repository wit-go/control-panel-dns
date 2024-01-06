// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/log"
)

func (hs *LinuxStatus) Show() {
	log.Log(CHANGE, "linuxStatus.Show() window")
	hs.window.Show()
	hs.hidden = false
}

func (hs *LinuxStatus) Hide() {
	log.Log(CHANGE, "linuxStatus.Hide() window")
	hs.window.Hide()
	hs.hidden = true
}

func (hs *LinuxStatus) Toggle() {
	log.Log(CHANGE, "linuxStatus.Toggle() window")
	if hs.hidden {
		hs.window.Show()
	} else {
		hs.window.Hide()
	}
}

func (hs *LinuxStatus) Ready() bool {
	if me == nil {return false}
	if hs == nil {return false}
	if hs.window == nil {return false}
	return me.ready
}
