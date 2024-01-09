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
	ready	bool
	hidden	bool

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
	actionLabel *gui.Node
	ipLabel *gui.Node
	statusLabel *gui.Node
	button *gui.Node

	problem *Problem
}

func NewErrorBox(p *gui.Node, name string, ip string) *errorBox {
	var eb *errorBox
	eb = new(errorBox)
	eb.parent = p
	eb.group = p.NewGroup(name)
	eb.grid = eb.group.NewGrid("stuff", 5, 1)

	eb.grid.NewLabel("Type")
	eb.grid.NewLabel("Action")
	eb.grid.NewLabel("IP")
	eb.grid.NewLabel("Status")
	eb.grid.NewLabel("")

	eb.fixes = make(map[string]*anError)
	eb.ready = true
	return eb
}

func (eb *errorBox) Show() {
	if eb == nil {return}
	eb.hidden = false
	eb.group.Show()
}

func (eb *errorBox) Hide() {
	if eb == nil {return}
	eb.hidden = true
	eb.group.Hide()
}

func (eb *errorBox) Toggle() {
	if eb == nil {return}
	if eb.hidden {
		eb.Show()
	} else {
		eb.Hide()
	}
}

func (eb *errorBox) Ready() bool {
	if eb == nil {return false}
	return eb.ready
}

func (eb *errorBox) addIPerror(kind ProblemType, action ActionType, ip string) bool {
	if ! eb.Ready() {return false}
	tmp := kind.String() + " " + ip
	if eb.fixes[tmp] != nil {
		log.Log(WARN, "Error is already here", kind, ip)
		return false
	}

	anErr := new(anError)
	anErr.aaaa = ip

	anErr.kindLabel = eb.grid.NewLabel(kind.String())
	anErr.actionLabel = eb.grid.NewLabel(action.String())
	anErr.ipLabel = eb.grid.NewLabel(ip)
	anErr.statusLabel = eb.grid.NewLabel("")
	anErr.button = eb.grid.NewButton("Try to Fix", func() {
		log.Log(WARN, "got", kind, "here. IP =", ip)
		eb.fix(tmp)
	})
	anErr.problem = new(Problem)
	anErr.problem.kind = kind
	anErr.problem.aaaa = ip
	eb.fixes[tmp] = anErr
	return false
}

// get all your problems!
func (eb *errorBox) Scan() []anError {
	for s, thing := range eb.fixes {
		log.Log(CHANGE, "Scan()", s, thing)
	}
	return nil
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
