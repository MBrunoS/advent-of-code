package main

import (
	"fmt"
)

func part1(lines []string) {
	sum := 0

	for i := 0; i < len(lines); i++ {
		occurrences := getNumberOccurrences(lines[i])

		for _, occ := range occurrences {
			if isPartNumber(occ, lines, i) {
				sum += occ.number
			}
		}

	}

	fmt.Println("Part 1: ", sum)
}
