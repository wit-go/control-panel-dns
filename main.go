// This creates a simple hello world window
package main

import 	(
	"log"
	"net"
	"github.com/fsnotify/fsnotify"
	"git.wit.org/wit/gui"
	arg "github.com/alexflint/go-arg"
)

func main() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)
	log.Println("Toolkit = ", args.Toolkit)

	// gui.InitPlugins([]string{"andlabs"})

	scanInterfaces()
	watchNetworkInterfaces()
	go inotifyNetworkInterfaceChanges()
	gui.Main(initGUI)
}

func watchNetworkInterfaces() {
	// Get list of network interfaces
	interfaces, _ := net.Interfaces()

	// Set up a notification channel
	notification := make(chan net.Interface)

	// Start goroutine to watch for changes
	go func() {
		for {
			// Check for changes in each interface
			for _, i := range interfaces {
				if status := i.Flags & net.FlagUp; status != 0 {
					notification <- i
					log.Println("something on i =", i)
				}
			}
		}
	}()
}

func inotifyNetworkInterfaceChanges() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Watch for network interface changes
	err = watcher.Add("/sys/class/net")
	if err != nil {
		return err
	}
	for {
		select {
		case event := <-watcher.Events:
			log.Println("inotifyNetworkInterfaceChanges() event =", event)
			if event.Op&fsnotify.Create == fsnotify.Create {
					// Do something on network interface creation
			}
		case err := <-watcher.Errors:
			return err
		}
	}
}

func scanInterfaces() {
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		log.Println(i)
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			log.Println("\taddr =", addr)
			switch v := addr.(type) {
			case *net.IPNet:
				log.Println("\t\taddr.(type) = *net.IPNet")
			default:
				log.Println("\t\taddr.(type) =", v)
			}
		}
	}
}
