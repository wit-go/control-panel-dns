package main

/*
	this enables command line options from other packages like 'gui' and 'log'
*/

import 	(
	arg "github.com/alexflint/go-arg"
	"go.wit.com/gui"
	"go.wit.com/log"
)


func init() {
	arg.MustParse()
	log.Bool(true, "INIT() args.ArgDebug =", gui.ArgDebug())
}
