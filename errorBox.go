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
	aaaa string
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
	anErr.aaaa = ip

	anErr.kindLabel = eb.grid.NewLabel(kind)
	anErr.ipLabel = eb.grid.NewLabel(ip)
	anErr.statusLabel = eb.grid.NewLabel("")
	anErr.button = eb.grid.NewButton(kind, func() {
		log.Log(WARN, "got", kind, "here. IP =", ip)
		eb.fix(tmp)
	})
	eb.fixes[tmp] = anErr
	return false
}

func (eb *errorBox) fix(key string) bool {
	if eb.fixes[key] == nil {
		log.Log(WARN, "Unknown error. could not find key =", key)
		log.Log(WARN, "TODO: probably remove this error. key =", key)
		return true
	}
	myErr :=  eb.fixes[key]
	log.Log(WARN, "should try to fix", myErr.kind, "here. IP =", myErr.aaaa)
	if myErr.kind == "DELETE" {
		if deleteFromDNS(myErr.aaaa) {
			log.Log(INFO, "Delete AAAA", myErr.aaaa, "Worked")
		} else {
			log.Log(INFO, "Delete AAAA", myErr.aaaa, "Failed")
		}
		return true
	}
	if myErr.kind == "CREATE" {
		if addToDNS(myErr.aaaa) {
			log.Log(INFO, "Delete AAAA", myErr.aaaa, "Worked")
		} else {
			log.Log(INFO, "Delete AAAA", myErr.aaaa, "Failed")
		}
		return true
	}
	return false
}

func (eb *errorBox) update() bool {
	return false
}

func (eb *errorBox) toggle() {
}
