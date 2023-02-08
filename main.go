// This creates a simple hello world window
package main

import 	(
	"log"
	"git.wit.org/wit/gui"
	arg "github.com/alexflint/go-arg"
)

func main() {
	arg.MustParse(&args)
	// fmt.Println(args.Foo, args.Bar, args.User)
	log.Println("Toolkit = ", args.Toolkit)

	// gui.InitPlugins([]string{"andlabs"})
	gui.Main(initGUI)
}
