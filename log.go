// I like things to be easy. Why can't the standard language be like this?

package main

import 	(
	"os"
	golog "log"
	"time"
	"reflect"
	"github.com/davecgh/go-spew/spew"
	// "net"
)

var 	LOGOFF	bool	= false	// turn this off, all logging stops
var	WARN	bool
var 	INFO	bool

type spewt struct {
	a bool
}

var SPEW spewt


/*
	sleep()		# you know what this does? sleeps for 1 second. yep. dump. easy.
	sleep(.1)	# you know what this does? yes, it sleeps for 1/10th of a second
*/
func sleep(a ...any) {
	if (a == nil) {
		time.Sleep(time.Second)
		return
	}

	log(args.Verbose, "sleep", a[0])

	switch a[0].(type) {
	case int:
		time.Sleep(time.Duration(a[0].(int)) * time.Second)
	case float64:
		time.Sleep(time.Duration(a[0].(float64) * 1000) * time.Millisecond)
	default:
		log("sleep a[0], type = ", a[0], reflect.TypeOf(a[0]))
	}
}

/*
	exit()		# yep. exits. I guess everything must be fine
	exit(3)		# I guess 3 it is then
	exit("dont like apples")	# ok. I'll make a note of that
*/
func exit(a ...any) {
	log("exit", a)
	//if (a) {
	//	os.Exit(a)
	//}
	os.Exit(0)
}

/*
	I've spent, am spending, too much time thinking about 'logging'. 'log', 'logrus', 'zap', whatever.
	I'm not twitter. i don't give a fuck about how many nanoseconds it takes to log. Anyway, this
	implementation is probably faster than all of those because you just set one bool to FALSE
	and it all stops.
	Sometimes I need to capture to stdout, sometimes stdout can't
	work because it doesn't exist for the user. This whole thing is a PITA. Then it's spread
	over 8 million references in every .go file. I'm tapping out and putting
	it in one place. here it is. Also, this makes having debug levels really fucking easy.
	You can define whatever level of logging you want from anywhere (command line) etc.

	log()		# doesn't do anything
	log(stuff)	# sends it to whatever log you define in a single place. here is the place
*/

func log(a ...any) {
	if (LOGOFF) {
		return
	}

	if (a == nil) {
		return
	}
	var blah bool
	if (reflect.TypeOf(a[0]) == reflect.TypeOf(blah)) {
		// golog.Println("\t a[0] = bool")
		if (a[0] == false) {
			return
		}
		a = a[1:]
	}

	if (reflect.TypeOf(a[0]) == reflect.TypeOf(SPEW)) {
		a = a[1:]
		spew.Dump(a)
		return
	}

	golog.Println(a...)
	// golog.Println("\t a[0] =", a[0])
	// for argNum, arg := range a {
	// 	golog.Println("\t", argNum, arg)
	// }
}
