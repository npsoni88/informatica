package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// To - do
// Find processor type and number of cores
// Find path
// convert memory from kb to MB
// Find volume and mount points
// Will decide
// test comment

func getOS() {
	var osystem string
	file, _ := os.Open("/etc/os-release")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME") {
			osystem = scanner.Text()
		}
	}

	osR := strings.Split(osystem, "=")
	fmt.Printf("Your operating system is - %v \n", osR[1])
}

func getMem() {
	var memTotal string
	file, _ := os.Open("/proc/meminfo")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "MemTotal") {
			memTotal = scanner.Text()
		}
	}

	memTotalR := strings.Split(memTotal, ":")

	fmt.Printf("Total Memory on this system is - %v\n", memTotalR[1])
}

func main() {
	var something int
	fmt.Println(`

###################################################################
#Script Name  - Informatica
#Description  - Displays most common information about your system
#Author       - Nishant (npsoni.com) (@npsoni88)
#Made with Go with love!
###################################################################
`)
	getOS()
	getMem()
}
