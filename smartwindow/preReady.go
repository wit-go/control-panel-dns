// these are things you can config
package smartwindow

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

/*	for now, run these before the window is ready
	That is, they all start with:

	if ! sw.Initialized() {return}
	if sw.Ready() {return}
*/

func (sw *SmartWindow) Title(title string) {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	log.Log(WARN, "Title() =", title)
	sw.title = title
}

func (sw *SmartWindow) SetParent(p *gui.Node) {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	log.Log(WARN, "SetParent")
	if sw.parent == nil {
		log.Log(WARN, "SetParent =", p)
		sw.parent = p
		return
	} else {
		log.Log(WARN, "SetParent was already set. TODO: Move to new parent")
	}
}

func (sw *SmartWindow) SetDraw(f func(*SmartWindow)) {
	log.Log(WARN, "SetDraw() START")
	if ! sw.Initialized() {
		log.Log(WARN, "SetDraw() Failed. sw.Initialized == false")
		return
	}
	if sw.Ready() {
		log.Log(WARN, "SetDraw() Failed. sw.Ready() == true")
		return
	}

	sw.populate = f
	log.Log(WARN, "SetDraw() END sw.populate is set")
}

func (sw *SmartWindow) Make() {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}
	log.Log(WARN, "Make() START")

	sw.window = sw.parent.RawWindow(sw.title)
	sw.window.Custom = func() {
		log.Warn("BasicWindow.Custom() closed. TODO: handle this", sw.title)
	}
	log.Log(WARN, "Make() END sw.window = RawWindow() (not sent to toolkits)")
	sw.ready = true
}

func (sw *SmartWindow) Vertical() {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	log.Log(WARN, "Vertical() setting vertical = true")
	sw.vertical = true
}
