package main

import (
	"bufio"
	"fmt"
)

func part1(scanner *bufio.Scanner) {
	var prevLine, line, nextLine string
	sum := 0

	scanner.Scan()
	nextLine = scanner.Text()

	for scanner.Scan() {
		prevLine = line
		line = nextLine
		nextLine = scanner.Text()
		occurrences := getNumberOccurrences(line)

		for _, occ := range occurrences {
			if isPartNumber(occ, prevLine, line, nextLine) {
				sum += occ.number
			}
		}

	}

	prevLine = line
	line = nextLine
	nextLine = ""

	occurrences := getNumberOccurrences(line)

	for _, occ := range occurrences {
		if isPartNumber(occ, prevLine, line, nextLine) {
			sum += occ.number
		}
	}

	fmt.Println("Part 1: ", sum)
}
