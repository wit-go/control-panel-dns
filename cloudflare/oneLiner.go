// This is a simple example
package cloudflare

import 	(
	"log"

	"go.wit.com/gui"
)

type OneLiner struct {
	p	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	v	*gui.Node	// value widget

	value	string
	label	string

	Custom func()
}

func (n *OneLiner) Set(value string) {
	log.Println("OneLiner.Set() =", value)
	n.v.Set(value)
	n.value = value
}

func NewOneLiner(n *gui.Node, name string) *OneLiner {
	d := OneLiner {
		p: n,
		value: "",
	}

	// various timeout settings
	d.l = n.NewLabel(name)
	d.v = n.NewLabel("")
	d.v.Custom = func() {
		d.value = d.v.S
		log.Println("OneLiner.Custom() user changed value to =", d.value)
	}

	return &d
}
