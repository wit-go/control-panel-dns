// This creates a 'smart window'
// it should work even when it is hidden
// from the gui toolkit plugins
package smartwindow

import 	(
	"go.wit.com/log"
)

func (sw *SmartWindow) Ready() bool {
	log.Log(INFO, "Ready() START")
	if sw == nil {return false}
	if sw.window == nil {return false}
	log.Log(INFO, "Ready() END sw.ready =", sw.ready)
	return sw.ready
}

func (sw *SmartWindow) Initialized() bool {
	log.Log(INFO, "checking Initialized()")
	if sw == nil {return false}
	return true
}
