package main

/*
	this parses the command line arguements
*/

import 	(
	"log"
	"fmt"
	"reflect"
	"strconv"
	arg "github.com/alexflint/go-arg"
	"git.wit.org/wit/gui"
	// log "git.wit.org/wit/gui/log"
)


var args struct {
	Verbose bool
	VerboseNet bool  `arg:"--verbose-net" help:"debug your local OS network settings"`
	VerboseDNS bool  `arg:"--verbose-dns" help:"debug your dns settings"`
	LogFile string `help:"write all output to a file"`
	// User string `arg:"env:USER"`
	Display string `arg:"env:DISPLAY"`

	Foo string
	Bar bool
	User string `arg:"env:USER"`
	Demo bool `help:"run a demo"`
	gui.GuiArgs
	// log.LogArgs
}

func init() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar, args.User)

	if (args.Gui != "") {
		gui.GuiArg.Gui = args.Gui
	}
	log.Println(true, "INIT() args.GuiArg.Gui =", gui.GuiArg.Gui)

	Set(&me, "default")
	log.Println("init() me.artificialSleep =", me.artificialSleep)
	log.Println("init() me.artificialS =", me.artificialS)
	me.artificialSleep = 2.3 
	log.Println("init() me.artificialSleep =", me.artificialSleep)
	sleep(me.artificialSleep)
}

func Set(ptr interface{}, tag string) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		log.Println(logError, "Set() Not a pointer", ptr, "with tag =", tag)
		return fmt.Errorf("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		defaultVal := t.Field(i).Tag.Get(tag)
		name := t.Field(i).Name
		// log("Set() try name =", name, "defaultVal =", defaultVal)
		setField(v.Field(i), defaultVal, name)
	}
	return nil
}

func setField(field reflect.Value, defaultVal string, name string) error {

	if !field.CanSet() {
		// log("setField() Can't set value", field, defaultVal)
		return fmt.Errorf("Can't set value\n")
	} else {
		log.Println("setField() Can set value", name, defaultVal)
	}

	switch field.Kind() {
	case reflect.Int:
		val, _ := strconv.Atoi(defaultVal)
		field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
	case reflect.Float64:
		val, _ := strconv.ParseFloat(defaultVal, 64)
		field.Set(reflect.ValueOf(float64(val)).Convert(field.Type()))
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	case reflect.Bool:
		if defaultVal == "true" {
			field.Set(reflect.ValueOf(true))
		} else {
			field.Set(reflect.ValueOf(false))
		}
	}

	return nil
}
