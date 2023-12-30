/* 
	The Digital Ocean Struct
*/

package digitalocean

import (
	"go.wit.com/gui"
	"go.wit.com/gui/gadgets"
)

type DigitalOcean struct {
	ready		bool
	hidden		bool

	token		string // You're Digital Ocean API key

	parent	*gui.Node // should be the root of the 'gui' package binary tree
	window	*gui.Node // our window for displaying digital ocean droplets
	group	*gui.Node // our window for displaying digital ocean droplets
	grid	*gui.Node // our window for displaying digital ocean droplets

	// Primary Directives
	status		*gadgets.OneLiner
	summary		*gadgets.OneLiner
	statusIPv4	*gadgets.OneLiner
	statusIPv6	*gadgets.OneLiner
}
