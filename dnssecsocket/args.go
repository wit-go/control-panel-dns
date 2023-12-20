package dnssecsocket

//
// By using the package "github.com/alexflint/go-arg",
// these can be configured from the command line
//

import (
	// arg "github.com/alexflint/go-arg"
	// "log"
	// "os"
)

type Args struct {
	VerboseDnssec bool `arg:"--verbose-dnssec" help:"debug dnssec lookups"`
	Foo string `arg:"env:USER"`
}

var args struct {
	Args
	Verbose bool
}

func Parse (b bool) {
	args.Verbose = b
	args.VerboseDnssec = b
}

// I attempted to pass the *arg.Parser down
// to see if I could find the value somewhere but I couldn't find it
/*
var conf arg.Config

func Parse (p *arg.Parser) {
	// conf.Program = "control-panel-dns"
	// conf.IgnoreEnv = false
	// arg.NewParser(conf, &args)
	log.Println("fuckit", p, args.VerboseDnssec)
	for i, v := range p.SubcommandNames() {
		log.Println("dnssec.Parse", i, v)
	}
	p.Jcarr()
}
*/
