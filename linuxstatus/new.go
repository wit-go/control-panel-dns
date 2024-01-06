// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/log"

	"go.wit.com/gui/gadgets"
)

func New() *LinuxStatus {
	if me != nil {
		log.Warn("You have done New() twice. You can only do this once")
		return me
	}
	me = &LinuxStatus {
		hidden: true,
		ready: false,
	}
	
	return me
}

func (ls *LinuxStatus) InitWindow() {
	if ! ls.Initialized() {
		log.Warn("LinuxStatus() is not initalized yet (no parent for the window?)")
		return
	}
	if ls.window != nil {
		log.Warn("You already have a window")
		ls.ready = true
		return
	}

	ls.ready = true
	log.Warn("Creating the Window")
	ls.window = gadgets.NewBasicWindow(ls.parent, "Linux OS Details")
}
