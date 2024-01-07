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
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	sw.populate = f
}

func (sw *SmartWindow) Make() {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	log.Log(WARN, "Make() window ready =", sw.ready)
	sw.window.Make()
	if (sw.populate != nil) {
		log.Log(WARN, "Make() trying to run Custom sw.populate() here")
		sw.populate(sw)
	}
	sw.ready = true
}

func (sw *SmartWindow) Draw() {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	log.Log(WARN, "Draw() window ready =", sw.ready)
	sw.window.Draw()
	if (sw.populate != nil) {
		log.Log(WARN, "Make() trying to run Custom sw.populate() here")
		sw.populate(sw)
	}
	sw.ready = true
}


func (sw *SmartWindow) Vertical() {
	if ! sw.Initialized() {return}
	if sw.Ready() {return}

	log.Log(WARN, "Draw() window ready =", sw.ready)
	sw.window.Draw()
	if (sw.populate != nil) {
		log.Log(WARN, "Make() trying to run Custom sw.populate() here")
		sw.populate(sw)
	}
	sw.ready = true
}

