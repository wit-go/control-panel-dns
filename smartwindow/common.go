// This creates a 'smart window'
// it should work even when it is hidden
// from the gui toolkit plugins
package smartwindow

import 	(
	"go.wit.com/log"
)

func (sw *SmartWindow) Ready() bool {
	log.Log(WARN, "Ready() maybe not ready? sw =", sw)
	if sw == nil {return false}
	if sw == nil {return false}
	if sw.window == nil {return false}
	return sw.ready
}

func (sw *SmartWindow) Initialized() bool {
	log.Log(WARN, "checking Initialized()")
	if sw == nil {return false}
	if sw == nil {return false}
	if sw.parent == nil {return false}
	return true
}
