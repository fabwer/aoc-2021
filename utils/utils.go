package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
)

func ReadLinesFromFile(filename string) []string {
	file := openFile(filename)
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		x := scanner.Text()
		lines = append(lines, x)
	}

	return lines
}

func ConvertStringsToInts(strings []string) []int {
	var ints []int
	for _, s := range strings {
		x, _ := strconv.Atoi(s)
		ints = append(ints, x)
	}
	return ints
}

func openFile(filename string) *os.File {
	absPath, _ := filepath.Abs("input/" + filename)

	file, openErr := os.Open(absPath)
	if openErr != nil {
		panic(openErr)
	}

	return file
}

func closeFile(file *os.File) {
	closeErr := file.Close()
	if closeErr != nil {
		panic(closeErr)
	}
}
