// This is a simple example
package main

import 	(
	"os"
	"log"
	"bufio"
	"strings"

	"go.wit.com/control-panel-dns/cloudflare"
)

func saveConfig() {
	log.Println("TODO")
}

func readConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("searchPaths() error. exiting here?")
	}
	filename := homeDir + "/" + configfile
	log.Println("filename =", filename)

	readFileLineByLine(filename)
	// os.Exit(0)
}

// readFileLineByLine opens a file and reads through each line.
func readFileLineByLine(filename string) error {
	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Println("readFileLineByLine() =", filename)

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(file)

	// Read through each line using scanner.
	for scanner.Scan() {
		var newc *cloudflare.ConfigT
		newc = new(cloudflare.ConfigT)

		line := scanner.Text()
		parts := strings.Fields(line)

		if (len(parts) < 4) {
			log.Println("readFileLineByLine() SKIP =", parts)
			continue
		}

		newc.Domain = parts[0]
		newc.ZoneID = parts[1]
		newc.Auth = parts[2]
		newc.Email = parts[3]

		cloudflare.Config[parts[0]] = newc
		log.Println("readFileLineByLine() =", newc.Domain, newc.ZoneID, newc.Auth, newc.Email)
	}

	// Check for errors during Scan.
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
