/* 
	Show your IPv6 addresses
*/

package main

import (
	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

type ipv6box struct {
	name	string // the problem name

	parent	*gui.Node
	group	*gui.Node
	grid	*gui.Node

	l	*gui.Node
	b	*gui.Node

	fixes	map[string]*anError

	something *gadgets.OneLiner
}

type anIPv6 struct {
	kind string // what kind of error is it?
	aaaa string
	status string

	kindLabel *gui.Node
	ipLabel *gui.Node
	statusLabel *gui.Node
	button *gui.Node
}

func NewIpv6box(p *gui.Node, name string, ip string) *ipv6box {
	var ib *ipv6box
	ib = new(ipv6box)
	ib.parent = p
	ib.group = p.NewGroup(name)
	ib.grid = ib.group.NewGrid("stuff", 4, 1)

	ib.grid.NewLabel("Type")
	ib.grid.NewLabel("IP")
	ib.grid.NewLabel("Status")
	ib.grid.NewLabel("")

	ib.fixes = make(map[string]*anError)
	return ib
}


func (ib *ipv6box) add(kind string, ip string) bool {
	tmp := kind + " " + ip
	if ib.fixes[tmp] != nil {
		log.Log(WARN, "Error is already here", kind, ip)
		return false
	}

	anErr := new(anError)
	anErr.kind = kind
	anErr.aaaa = ip

	anErr.kindLabel = ib.grid.NewLabel(kind)
	anErr.ipLabel = ib.grid.NewLabel(ip)
	anErr.statusLabel = ib.grid.NewLabel("")
	anErr.button = ib.grid.NewButton(kind, func() {
		log.Log(WARN, "got", kind, "here. IP =", ip)
		ib.fix(tmp)
	})
	ib.fixes[tmp] = anErr
	return false
}

func (ib *ipv6box) fix(key string) bool {
	if ib.fixes[key] == nil {
		log.Log(WARN, "Unknown error. could not find key =", key)
		log.Log(WARN, "TODO: probably remove this error. key =", key)
		return true
	}
	myErr :=  ib.fixes[key]
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

func (ib *ipv6box) update() bool {
	return false
}

func (ib *ipv6box) toggle() {
}
