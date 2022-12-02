package main

import (
	"fmt"
	"testing"
)

// resultsMap := map[string]string{
//	"A": "Rock",
//	"B": "Paper",
//	"C": "Scissors",
//	"X": "Rock",
//	"Y": "Paper",
//	"Z": "Scissors",
// }

func TestCalculateScore(t *testing.T) {
	a := calculateScore([]string{"A", "Y"})
	if a != 8 {
		fmt.Println("Actual:", a, "Expected", 8)
		t.Fail()
	}
	a = calculateScore([]string{"B", "X"})
	if a != 1 {
		fmt.Println("Actual:", a, "Expected", 1)
		t.Fail()
	}
	a = calculateScore([]string{"C", "Z"})
	if a != 6 {
		fmt.Println("Actual:", a, "Expected", 6)
		t.Fail()
	}
}
