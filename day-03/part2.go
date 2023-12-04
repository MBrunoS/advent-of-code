package main

import "fmt"

func part2(lines []string) {
	sum := 0

	for i := 0; i < len(lines); i++ {
		indexes := getGearsIndexes(lines[i])

		for _, index := range indexes {
			adjPartNums := getAdjacentPartNumbers(index, lines, i)
			if len(adjPartNums) == 2 {
				sum += adjPartNums[0] * adjPartNums[1]
			}
		}

	}

	fmt.Println("Part 2: ", sum)
}

func getGearsIndexes(line string) []int {
	var indexes []int

	for i, c := range line {
		if c == '*' {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func getAdjacentPartNumbers(index int, lines []string, lineIndex int) []int {
	var prevLine, nextLine string
	adjacentPartNumbers := []int{}

	// previous line

	if lineIndex > 0 {
		prevLine = lines[lineIndex-1]
	} else {
		prevLine = ""
	}

	numbersOcc := getNumberOccurrences(prevLine)

	for _, occ := range numbersOcc {
		if isPartNumber(occ, lines, lineIndex-1) {
			if index >= occ.startIndex-1 && index <= occ.endIndex+1 {
				adjacentPartNumbers = append(adjacentPartNumbers, occ.number)
			}
		}
	}

	// current line

	line := lines[lineIndex]
	numbersOcc = getNumberOccurrences(line)

	for _, occ := range numbersOcc {
		if isPartNumber(occ, lines, lineIndex) {
			if index >= occ.startIndex-1 && index <= occ.endIndex+1 {
				adjacentPartNumbers = append(adjacentPartNumbers, occ.number)
			}
		}
	}

	// next line

	if lineIndex < len(lines)-1 {
		nextLine = lines[lineIndex+1]
	} else {
		nextLine = ""
	}

	numbersOcc = getNumberOccurrences(nextLine)

	for _, occ := range numbersOcc {
		if isPartNumber(occ, lines, lineIndex+1) {
			if index >= occ.startIndex-1 && index <= occ.endIndex+1 {
				adjacentPartNumbers = append(adjacentPartNumbers, occ.number)
			}
		}
	}

	return adjacentPartNumbers
}
