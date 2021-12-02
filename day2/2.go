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

	// part 2
	fmt.Println(calcPositionWithAim(lines))
}

func calcPosition(lines []string) int {
	posX := 0
	posZ := 0

	for _, line := range lines {
		command, number := parseCommandAndNumber(line)

		if command == "forward" {
			posX += number
		} else if command == "up" {
			posZ -= number
		} else if command == "down" {
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
		command, number := parseCommandAndNumber(line)

		if command == "forward" {
			posX += number
			posZ += number * aim
		} else if command == "up" {
			aim -= number
		} else if command == "down" {
			aim += number
		}
	}

	return posX * posZ
}

func parseCommandAndNumber(line string) (string, int) {
	args := strings.Fields(line)
	command := args[0]
	number, _ := strconv.Atoi(args[1])

	return command, number
}
