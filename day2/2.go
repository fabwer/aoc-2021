package main

import (
	"aoc-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadStringsFromFile("2_input.txt")

	// part 1
	fmt.Println(calcPosition(lines))

	// part 1
	fmt.Println(calcPositionWithAim(lines))
}

func calcPosition(lines []string) int {
	posX := 0
	posZ := 0

	for _, line := range lines {
		number, _ := strconv.Atoi(line[len(line)-1:])
		if strings.Contains(line, "forward") {
			posX += number
		} else if strings.Contains(line, "up") {
			posZ -= number
		} else if strings.Contains(line, "down") {
			posZ += number
		}
	}

	return posX * posZ
}

func calcPositionWithAim(lines []string) int {
	posX := 0
	posZ := 0
	aim := 0

	for _, line := range lines {
		number, _ := strconv.Atoi(line[len(line)-1:])
		if strings.Contains(line, "forward") {
			posX += number
			posZ += number * aim
		} else if strings.Contains(line, "up") {
			aim -= number
		} else if strings.Contains(line, "down") {
			aim += number
		}
	}

	return posX * posZ
}
