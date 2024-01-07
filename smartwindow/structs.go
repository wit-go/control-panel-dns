package smartwindow

import (
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

type SmartWindow struct {
	ready	bool	// track if the window is ready
	hidden	bool	// track if the window is hidden from the toolkits
	changed	bool	// track if something changed in the window

	title	string	// what the user sees as the name
	name	string	// the programatic name aka: "CALANDAR"

	parent	*gui.Node // where to place the window if you try to draw it
	window	*gadgets.BasicWindow // the underlying BasicWindow
	box	*gui.Node // the box inside the window // get this from BasicWindow() ?

	populate func(*SmartWindow) // the function to generate the widgets
}
