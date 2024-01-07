package smartwindow

import 	(
	"go.wit.com/log"
)

func New() *SmartWindow {
	sw := SmartWindow {
		hidden: true,
		ready: false,
	}
	
	return &sw
}

func (sw *SmartWindow) InitWindow() {
	if sw == nil {
		log.Log(WARN, "not initalized yet (no parent for the window?)")
		return
	}
	if sw.window != nil {
		log.Log(WARN, "You already have a SmartWindow")
		sw.ready = true
		return
	}
}
