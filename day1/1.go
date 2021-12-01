package main

import (
	u "aoc-2021/utils"
	"fmt"
)

func main() {
	measurements := u.ReadNumbersFromFile("1_input.txt")

	// first
	fmt.Println(countMovingWindow(measurements, 1))

	// second
	fmt.Println(countMovingWindow(measurements, 3))
}

func countMovingWindow(measurements []int, windowLength int) int {
	var counter int
	for i := 0; i < len(measurements)-windowLength; i++ {
		// a + b + c < b + c + d iff a < d
		if measurements[i] < measurements[i+windowLength] {
			counter++
		}
	}

	return counter
}
