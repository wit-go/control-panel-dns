// this defines the kinds of problems that can be detected
package main

import 	(
	"time"
)

type ProblemType int
type ActionType int

type Problem struct {
	kind ProblemType
	action ActionType

	id int
	Name string
	desc string
	value string
	aaaa string
	fixed bool
	duration time.Duration // how long until you should check to see if it's fixed
	born time.Time // when first reported

	begun bool // weather or not fixing it has begun
	begunTime time.Time // when the attempt to fix the problem happened
	begunResult bool // weather or not the attempt worked
}

var IPcreate = Problem {
	kind: RR,
	action: CREATE,
	desc: "This RR entry in the zonefile needs to be removed",
}

/*
var hostname Problem = (
	kind: ProblemType.OS,
	action: ActionType.CREATE,
	Name: "Your /etc/hostname file is incorrect",
	fixed: false,
)
*/

const (
	OS ProblemType = iota
	ETC
	RESOLVE
	RR
	PING
	LOOKUP
)

const (
	USER ActionType = iota
	CREATE
	DELETE
)

func (s Problem) String() string {
	return s.Name
}

func (s ProblemType) String() string {
	switch s {
	case OS:
		return "OS"
	case RR:
		return "RR"
	default:
		return "FIXMEP"
	}
	return "FIXMEP"
}

func (s ActionType) String() string {
	switch s {
	case USER:
		return "USER"
	case CREATE:
		return "CREATE"
	case DELETE:
		return "DELETE"
	default:
		return "FIXMEA"
	}
	return "FIXMEA"
}
