package smartwindow

import 	(
	"go.wit.com/log"

	"go.wit.com/gui/gadgets"
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

	log.Log(WARN, "Creating the Window")
	sw.window = gadgets.NewBasicWindow(sw.parent, sw.title)
	sw.ready = true
}
