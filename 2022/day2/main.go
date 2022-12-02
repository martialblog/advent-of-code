package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var resultsMap = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

func chooseAnswer(results []string) string {
	// X means you need to lose,
	// Y means you need to end the round in a draw
	// Z means you need to win.
	resultOther := resultsMap[results[0]]

	switch results[1] {
	case "X":
		switch resultOther {
		case "Rock":
			return "C"
		case "Paper":
			return "A"
		case "Scissors":
			return "B"
		}
	case "Y":
		switch resultOther {
		case "Rock":
			return "A"
		case "Paper":
			return "B"
		case "Scissors":
			return "C"
		}
	case "Z":
		switch resultOther {
		case "Rock":
			return "B"
		case "Paper":
			return "C"
		case "Scissors":
			return "A"
		}
	}
	return ""
}

func calculateScore(results []string) int {
	// The score for a single round is the score for the shape you selected
	// (1 for Rock, 2 for Paper, and 3 for Scissors)
	// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	var score int

	scoreMap := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	// First Col:  A for Rock, B for Paper, and C for Scissors.
	// The score for a single round is the score for the shape you selected
	// (1 for Rock, 2 for Paper, and 3 for Scissors)

	// Map results to actual values
	// Just for readability
	resultOther := resultsMap[results[0]]
	resultMe := resultsMap[results[1]]

	// Score for shape
	score += scoreMap[resultMe]

	// Draw == 3 Points
	if resultMe == resultOther {
		score += 3
		return score
	}

	// Paper beats Rock
	// Rock beats Scissors
	// Scissors beats Paper

	if resultMe == "Rock" {
		if resultOther == "Scissors" {
			// Won
			score += 6
		}
		if resultOther == "Paper" {
			/// Lost
			score += 0
		}
	}

	if resultMe == "Paper" {
		if resultOther == "Rock" {
			// Won
			score += 6
		}
		if resultOther == "Scissors" {
			// Lost
			score += 0
		}
	}

	if resultMe == "Scissors" {
		if resultOther == "Paper" {
			// Won
			score += 6
		}
		if resultOther == "Rock" {
			// Lost
			score += 0
		}
	}

	return score
}

func main() {
	// Second Col: X for Rock, Y for Paper, and Z for Scissors.
	// Total Score: Sum of socres for each round

	// The score for a single round is the score for the shape you selected
	// (1 for Rock, 2 for Paper, and 3 for Scissors)
	// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	// Read the input file
	var inputFile string
	var strategy bool

	flag.BoolVar(&strategy, "strategy", false, "Use second col to choose answer")
	flag.StringVar(&inputFile, "file", "input.txt", "Path to Inputfile")

	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize Slices
	var finalScore int

	for scanner.Scan() {
		r := scanner.Text()
		results := strings.Split(r, " ")

		// Use the second col to select answer
		if strategy {
			results[1] = chooseAnswer(results)
		}

		score := calculateScore(results)
		finalScore += score
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Final Score
	fmt.Println(finalScore)
}
