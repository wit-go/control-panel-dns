/* 
	The Digital Ocean Struct
*/

package digitalocean

import (
	"github.com/digitalocean/godo"

	"go.wit.com/gui"
	"go.wit.com/gui/gadgets"
)

type DigitalOcean struct {
	ready		bool
	hidden		bool
	err		error

	token		string // You're Digital Ocean API key
	droplets	[]godo.Droplet

	parent	*gui.Node // should be the root of the 'gui' package binary tree
	window	*gui.Node // our window for displaying digital ocean droplets
	group	*gui.Node // our window for displaying digital ocean droplets
	grid	*gui.Node // our window for displaying digital ocean droplets

	dGrid	*gui.Node // the grid for the droplets

	// Primary Directives
	status		*gadgets.OneLiner
	summary		*gadgets.OneLiner
	statusIPv4	*gadgets.OneLiner
	statusIPv6	*gadgets.OneLiner
}

type ipButton struct {
	ip	*gui.Node
	c	*gui.Node
}

type Droplet struct {
	ready		bool
	hidden		bool
	err		error

	poll		*godo.Droplet // store what the digital ocean API returned

	name		*gui.Node

	// a box and grid of the IPv4 addresses
	box4		*gui.Node
	grid4		*gui.Node
	ipv4		[]ipButton

	// a box and grid of the IPv6 addresses
	box6		*gui.Node
	grid6		*gui.Node
	ipv6		[]ipButton
}
