package digitalocean

import 	(
	"log"

	"go.wit.com/gui"
)

func MakeWindow(n *gui.Node) *gui.Node {
	log.Println("digitalocean MakeWindow() START")

	win := n.NewWindow("DigitalOcean Control Panel")

	// box := g1.NewGroup("data")
	group := win.NewGroup("data")
	log.Println("digitalocean MakeWindow() END", group)
	return win
}
