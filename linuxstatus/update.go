package linuxstatus

import (
	"fmt"
	"time"
	"errors"

	"go.wit.com/log"
)

func (ls *LinuxStatus) Update() {
	if ! ls.Ready() {
		log.Log(WARN, "can't update yet. ready is false")
		log.Error(errors.New("Update() is not ready yet"))
		return
	}
	log.Log(INFO, "Update() START")
	duration := timeFunction(func () {
		linuxLoop()
	})
	ls.SetSpeed(duration)
	log.Log(INFO, "Update() END")
}

func (ls *LinuxStatus) SetSpeed(duration time.Duration) {
	s := fmt.Sprint(duration)
	if ls.speedActual == nil {
		log.Log(WARN, "can't actually warn")
		return
	}
	ls.speedActual.Set(s)

	if (duration > 500 * time.Millisecond ) {
		ls.speed.Set("SLOW")
	} else if (duration > 100 * time.Millisecond ) {
		ls.speed.Set("OK")
	} else {
		ls.speed.Set("FAST")
	}
}
