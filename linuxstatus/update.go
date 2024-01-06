package linuxstatus

import (
	"errors"
	"fmt"
	"time"

	"go.wit.com/log"
)

func (ls *LinuxStatus) Update() {
	log.Info("linuxStatus() Update() START")
	if ls == nil {
		log.Error(errors.New("linuxStatus() Update() ls == nil"))
		return
	}
	duration := timeFunction(func () {
		linuxLoop()
	})
	s := fmt.Sprint(duration)
	ls.speedActual.Set(s)

	if (duration > 500 * time.Millisecond ) {
		// ls.speed, "SLOW")
	} else if (duration > 100 * time.Millisecond ) {
		// ls.speed, "OK")
	} else {
		// ls.speed, "FAST")
	}
	log.Info("linuxStatus() Update() END")
}
