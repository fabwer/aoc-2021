package main

import (
	"aoc-2021/utils"
	"fmt"
	"strconv"
)

const (
	zero uint8 = iota
	one
)

func main() {
	lines := utils.ReadLinesFromFile("3_input")

	// part 1
	fmt.Println(calcPowerConsuption(lines))

	// part 2
	fmt.Println(calcLifeSupportRating(lines))
}

func calcPowerConsuption(lines []string) int64 {
	gamma := ""
	epsilon := ""

	for position := 0; position < len(lines[0]); position++ {
		ones, zeros := countAtPosition(position, lines)
		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return g * e
}

func calcLifeSupportRating(lines []string) int64 {
	o := findSignificantLine(lines, one)
	co2 := findSignificantLine(lines, zero)

	return o * co2
}

func countAtPosition(position int, lines []string) (ones int, zeros int) {
	for _, line := range lines {
		if line[position] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return
}

func calcCommonBit(factor uint8, ones int, zeros int) byte {
	if (factor == 1 && ones >= zeros) || (factor == 0 && ones < zeros) {
		return '1'
	}
	return '0'
}

func findSignificantLine(lines []string, factor uint8) int64 {
	filtered := lines
	for i := 0; len(filtered) > 1; i++ {
		ones, zeros := countAtPosition(i, filtered)
		commonBitByFactor := calcCommonBit(factor, ones, zeros)
		var newFiltered []string
		for _, line := range filtered {
			if line[i] == commonBitByFactor {
				newFiltered = append(newFiltered, line)
			}
		}
		filtered = newFiltered
	}

	significantLine, _ := strconv.ParseInt(filtered[0], 2, 64)
	return significantLine
}
