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

	fixes	map[string]*anError

	something *gadgets.OneLiner
}

type anError struct {
	kind string // what kind of error is it?
	ip string
	status string

	kindLabel *gui.Node
	ipLabel *gui.Node
	statusLabel *gui.Node
	button *gui.Node
}

func NewErrorBox(p *gui.Node, name string, ip string) *errorBox {
	var eb *errorBox
	eb = new(errorBox)
	eb.parent = p
	eb.group = p.NewGroup(name)
	eb.grid = eb.group.NewGrid("stuff", 4, 1)

	eb.grid.NewLabel("Type")
	eb.grid.NewLabel("IP")
	eb.grid.NewLabel("Status")
	eb.grid.NewLabel("")

	eb.fixes = make(map[string]*anError)
	return eb
}


func (eb *errorBox) add(kind string, ip string) bool {
	tmp := kind + " " + ip
	if eb.fixes[tmp] != nil {
		log.Log(WARN, "Error is already here", kind, ip)
		return false
	}

	anErr := new(anError)
	anErr.kind = kind

	anErr.kindLabel = eb.grid.NewLabel(kind)
	anErr.ipLabel = eb.grid.NewLabel(ip)
	anErr.statusLabel = eb.grid.NewLabel("")
	anErr.button = eb.grid.NewButton(kind, func() {
		log.Log(WARN, "got", kind, "here. IP =", ip)
		eb.fix(kind, ip)
	})
	eb.fixes[tmp] = anErr
	return false
}

func (eb *errorBox) fix(name string, ip string) bool {
	log.Log(WARN, "should try to fix", name, "here. IP =", ip)
	return false
}

func (eb *errorBox) update() bool {
	return false
}

func (eb *errorBox) toggle() {
}
