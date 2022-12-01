package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func sumSlice(sl []int) int {
	t := 0
	for _, v := range sl {
		t += v
	}
	return t
}

func main() {
	// Read input file
	var inputFile string
	var noOfOutputs int
	flag.StringVar(&inputFile, "file", "input.txt", "Path to Inputfile")
	flag.IntVar(&noOfOutputs, "elves", 1, "Number of top Elves to show")

	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize Slices
	cals := make([]int, 1)
	calSum := make([]int, 0)

	for scanner.Scan() {
		// An empty line as delimiter for Elves
		if scanner.Text() == "" {
			// Sum existing and add to Final Slice
			cals = append(cals, sumSlice(calSum))
			// Reset counting slice
			calSum = make([]int, 0)
		} else {
			// Add calories to counting slice
			c := scanner.Text()
			cal, _ := strconv.Atoi(c)
			calSum = append(calSum, cal)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort our Final Slice
	sort.Slice(cals, func(i, j int) bool {
		return cals[i] < cals[j]
	})

	// Return the highest Calorie Count
	fmt.Println("Top Elves")
	fmt.Println(cals[len(cals)-noOfOutputs:])
	fmt.Println("Sum of Top Elves")
	fmt.Println(sumSlice(cals[len(cals)-noOfOutputs:]))
}
