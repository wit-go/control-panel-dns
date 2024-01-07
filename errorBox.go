/* 
	Show a box for a configuration error
*/

package main

import (
	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

type errorBox struct {
	name	string // the problem name

	parent	*gui.Node
	group	*gui.Node
	grid	*gui.Node

	l	*gui.Node
	b	*gui.Node

	something *gadgets.OneLiner
}

func NewErrorBox(p *gui.Node, name string) *errorBox {
	var eb *errorBox
	eb = new(errorBox)
	eb.parent = p
	// eb.group = p.NewGroup("eg")
	// eb.grid = eb.group.NewGrid("labels", 2, 1)

	eb.l = p.NewLabel("click to fix")
	eb.b = p.NewButton("fix", func() {
		log.Log(WARN, "should try to fix here")
	})
	eb.something = gadgets.NewOneLiner(eb.grid, "something")

	return eb
}

func (eb *errorBox) update() bool {
	return false
}

func (eb *errorBox) toggle() {
}
