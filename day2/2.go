package main

import (
	"aoc-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLinesFromFile("2_input")

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

		if command == "f" {
			posX += number
		} else if command == "u" {
			posZ -= number
		} else if command == "d" {
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

		if command == "f" {
			posX += number
			posZ += number * aim
		} else if command == "u" {
			aim -= number
		} else if command == "d" {
			aim += number
		}
	}

	return posX * posZ
}

func parseCommandAndNumber(line string) (string, int) {
	args := strings.Fields(line)
	command := args[0]
	number, _ := strconv.Atoi(args[1])

	return command[0:1], number
}
