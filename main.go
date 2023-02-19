// This is a control panel for DNS
package main

import 	(
	"runtime"
	// "net"
	"git.wit.org/wit/gui"
	arg "github.com/alexflint/go-arg"
)

func main() {
	arg.MustParse(&args)

	// initialize the maps to track IP addresses and network interfaces
	me.ipmap = make(map[string]*IPtype)
	me.ifmap = make(map[int]*IFtype)

	go checkNetworkChanges()

	log()
	log(true, "this is true")
	log(false, "this is false")
	sleep(.4)
	sleep(.3)
	sleep(.2)
	sleep("done scanning net")
	// exit("done scanning net")

	// Example_listLink()
	// exit()

	log("Toolkit = ", args.Toolkit)
	// gui.InitPlugins([]string{"andlabs"})
	gui.Main(initGUI)
}

/*
	Poll for changes to the networking settings
*/
func checkNetworkChanges() {
	for {
		sleep(2)
		if (runtime.GOOS == "linux") {
			scanInterfaces()
		} else {
			log("Windows and MacOS don't work yet")
		}
	}
}
