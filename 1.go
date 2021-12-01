package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// first
	measurements := readNumbersFromFile()

	var counter int
	for i := 0; i < len(measurements)-1; i++ {
		if measurements[i] < measurements[i+1] {
			counter++
		}
	}

	fmt.Println(counter)

	// second
	// input is the same
	var counterMovingWindow int

	for i := 0; i < len(measurements)-3; i++ {
		// a + b + c < b + c + d iff a < d
		if measurements[i] < measurements[i+3] {
			counterMovingWindow++
		}
	}

	fmt.Println(counterMovingWindow)
}

func readNumbersFromFile() []int {
	file, openErr := os.Open("1_input.txt")

	if openErr != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var numbers []int

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, x)
	}

	closeErr := file.Close()
	if closeErr != nil {
		return []int{}
	}

	return numbers
}
