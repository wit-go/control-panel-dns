/* 
	Show a box for a configuration error
*/

package main

import (
	"time"

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
//	kind ProblemType // what kind of error is it?
//	action ActionType
//	aaaa string
//	status string

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
		log.Log(WARN, "kind =", kind)
		log.Log(WARN, "action =", action)
		log.Log(WARN, "ip =", ip)
		return false
	}

	anErr := new(anError)

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
	anErr.problem.action = action
	anErr.problem.aaaa = ip
	anErr.problem.born = time.Now()
	anErr.problem.duration = 30 * time.Second
	anErr.problem.begun = false
	anErr.problem.begunResult = false
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
	log.Log(WARN, "should try to fix", myErr.problem.kind, "here. IP =", myErr.problem.aaaa)
	if ! me.autofix.GetBool() {
		log.Log(WARN, "not autofixing. autofix == false")
		log.Log(WARN, "problem.kind =", myErr.problem.kind)
		log.Log(WARN, "problem.action =", myErr.problem.action)
		log.Log(WARN, "problem.aaaa =", myErr.problem.aaaa)
		log.Log(WARN, "problem.duration =", myErr.problem.duration)
		log.Log(WARN, "problem.begun =", myErr.problem.begun)
		log.Log(WARN, "problem.begunResult =", myErr.problem.begunResult)
		// if  myErr.problem.begunTime != nil {
			log.Log(WARN, "problem.begunTime =", myErr.problem.begunTime)
		// }
		return false
	}
	if myErr.problem.begun {
		log.Log(WARN, "problem has already begun. need to check the status of the problem here")
		log.Log(WARN, "problem.begun =", myErr.problem.begun)
		log.Log(WARN, "problem.begunResult =", myErr.problem.begunResult)
		log.Log(WARN, "problem.duration =", myErr.problem.duration)
		delay := time.Since(myErr.problem.begunTime)
		log.Log(WARN, "problem duration time =", delay)
		if delay >= myErr.problem.duration {
			log.Log(WARN, "duration eclipsed. check the status of the error here")
		}
		return false
	}
	if myErr.problem.kind == RR {
		if myErr.problem.action == DELETE {
			myErr.problem.begun = true
			myErr.problem.begunTime = time.Now()
			if deleteFromDNS(myErr.problem.aaaa) {
				log.Log(INFO, "Delete AAAA", myErr.problem.aaaa, "Worked")
				myErr.problem.begunResult = true
			} else {
				log.Log(INFO, "Delete AAAA", myErr.problem.aaaa, "Failed")
				myErr.problem.begunResult = false
			}
			return true
		}
		if myErr.problem.action == CREATE {
			myErr.problem.begun = true
			myErr.problem.begunTime = time.Now()
			if addToDNS(myErr.problem.aaaa) {
				log.Log(INFO, "Delete AAAA", myErr.problem.aaaa, "Worked")
				myErr.problem.begunResult = true
			} else {
				log.Log(INFO, "Delete AAAA", myErr.problem.aaaa, "Failed")
				myErr.problem.begunResult = false
			}
			return true
		}
	}
	return false
}

func (eb *errorBox) update() bool {
	return false
}

func (eb *errorBox) toggle() {
}
