package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadNumbersFromFile(filename string) []int {
	absPath, _ := filepath.Abs("input/" + filename)

	file, openErr := os.Open(absPath)
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

func ReadStringsFromFile(filename string) []string {
	absPath, _ := filepath.Abs("input/" + filename)

	file, openErr := os.Open(absPath)
	if openErr != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var strings []string
	for scanner.Scan() {
		x := scanner.Text()
		strings = append(strings, x)
	}

	closeErr := file.Close()
	if closeErr != nil {
		return []string{}
	}

	return strings
}
