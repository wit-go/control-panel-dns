package main

import (
	"bytes"
	"fmt"
	"strings"
	"os/exec"

	"git.wit.org/wit/shell"
)

func run(s string) string {
	cmdArgs := strings.Fields(s)
	// Define the command you want to run
	// cmd := exec.Command(cmdArgs)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)

	// Create a buffer to capture the output
	var out bytes.Buffer

	// Set the output of the command to the buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return ""
	}

	tmp := shell.Chomp(out.String())
	// Output the results
	fmt.Println("Command Output:", tmp)

	return tmp
}

