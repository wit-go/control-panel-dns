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
	value string
	fixed bool
	duration *time.Duration
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
