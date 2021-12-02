package main

import (
	u "aoc-2021/utils"
	"fmt"
)

func main() {
	measurements := u.ConvertStringsToInts(u.ReadLinesFromFile("1_input"))

	// part 1
	fmt.Println(countMovingWindow(measurements, 1))

	// part 2
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
