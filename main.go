package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// To - do
// Find processor type and number of cores
// Find path
// Find memory details
// Find volume and mount points
// Will decide later

func getOS() {
	var osystem string
	file, err := os.Open("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME") {
			osystem = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	osR := strings.Split(osystem, "=")
	fmt.Printf("Your operating system is - %v \n", osR[1])
}

func main() {
	fmt.Println(`

###################################################################
#Script Name  - Informatica
#Description  - Displays most common information about your system
#Author       - Nishant (npsoni.com) (@npsoni88)
#Made with Go with love!
###################################################################
`)
	getOS()
}
