// This creates a simple hello world window
package linuxstatus

import 	(
)

func New() *LinuxStatus {
	me = &LinuxStatus {
		hidden: true,
		ready: false,
	}
	
	me.init = true
	return me

	// me.window = gadgets.NewBasicWindow(me.myGui, "Linux OS Details")
}
