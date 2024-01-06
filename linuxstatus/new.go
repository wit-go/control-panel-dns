// This creates a simple hello world window
package linuxstatus

import 	(
	"go.wit.com/log"

	"go.wit.com/gui/gadgets"
)

func New() *LinuxStatus {
	if me != nil {
		log.Log(WARN, "You have done New() twice. You can only do this once")
		return me
	}
	me = &LinuxStatus {
		hidden: true,
		ready: false,
	}
	me.ifmap = make(map[int]*IFtype)
	me.ipmap = make(map[string]*IPtype)
	
	return me
}

func (ls *LinuxStatus) InitWindow() {
	if ! ls.Initialized() {
		log.Log(WARN, "not initalized yet (no parent for the window?)")
		return
	}
	if ls.window != nil {
		log.Log(WARN, "You already have a window")
		ls.ready = true
		return
	}

	log.Log(WARN, "Creating the Window")
	ls.window = gadgets.NewBasicWindow(ls.parent, "Linux OS Details")
	ls.ready = true
}
