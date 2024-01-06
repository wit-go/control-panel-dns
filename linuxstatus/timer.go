package linuxstatus

import (
	"time"
	"sort"
	"strings"
)

// timeFunction takes a function as an argument and returns the execution time.
func timeFunction(f func()) time.Duration {
	startTime := time.Now() // Record the start time
	f()                     // Execute the function
	return time.Since(startTime) // Calculate the elapsed time
}

// sortLines takes a string, splits it on newlines, sorts the lines,
// and rejoins them with newlines.
func sortLines(input string) string {
	lines := strings.Split(input, "\n")

	// Trim leading and trailing whitespace from each line
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	sort.Strings(lines)
	tmp := strings.Join(lines, "\n")
	tmp = strings.TrimLeft(tmp, "\n")
	tmp = strings.TrimRight(tmp, "\n")
	return tmp
}
