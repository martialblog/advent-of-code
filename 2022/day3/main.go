package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const az = "abcdefghijklmnopqrstuvwxyz"
const AZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var alphaMap = map[string]int{}

func init() {
	for idx, ch := range az {
		alphaMap[string(ch)] = idx + 1
	}
	for idx, ch := range AZ {
		alphaMap[string(ch)] = idx + 26 + 1
	}
}

func findCharsinBoth(a string, b string) []string {
	var s []string

	for _, ch := range b {
		if strings.Contains(a, string(ch)) {
			s = append(s, string(ch))
		}
	}
	return s
}

func main() {

	// Read the input file
	var inputFile string

	flag.StringVar(&inputFile, "file", "input.txt", "Path to Inputfile")

	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize Slices
	var finalPrio int
	var prio int

	for scanner.Scan() {
		prio = 0
		r := scanner.Text()
		// Length of the content
		l := len([]rune(r))
		compA := r[0 : l/2]
		compB := r[l/2 : l]

		charsInBoth := findCharsinBoth(compA, compB)

		prio += alphaMap[charsInBoth[0]]

		finalPrio += prio
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Final Score
	fmt.Println(finalPrio)
}
