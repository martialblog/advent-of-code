package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func areaContains(a []string, b []string) bool {
	aStart, _ := strconv.Atoi(a[0])
	aStop, _ := strconv.Atoi(a[1])
	bStart, _ := strconv.Atoi(b[0])
	bStop, _ := strconv.Atoi(b[1])

	if bStart >= aStart && bStop <= aStop {
		return true
	}
	return false
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

	// Initialize Final Count
	var finalCount int

	// Part 1
	for scanner.Scan() {
		r := scanner.Text()
		// 2-4,6-8
		areas := strings.Split(r, ",")
		areaOne := strings.Split(areas[0], "-")
		areaTwo := strings.Split(areas[1], "-")

		oneContainsTwo := areaContains(areaOne, areaTwo)

		if oneContainsTwo {
			finalCount += 1
			continue
		}
		twoContainsOne := areaContains(areaTwo, areaOne)
		if twoContainsOne {
			finalCount += 1
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Final Count
	// Part 1
	fmt.Println(finalCount)
}
