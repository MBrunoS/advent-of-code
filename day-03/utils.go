package main

import (
	"strconv"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

type NumberOcurrence struct {
	number     int
	startIndex int
	endIndex   int
}

func getNumberOccurrences(line string) []NumberOcurrence {
	var numbers []NumberOcurrence

	for i := 0; i < len(line); i++ {
		c := line[i]

		if isDigit(c) {
			var num string
			startIndex := i

			for isDigit(c) {
				num += string(c)
				if i == len(line)-1 {
					break
				}
				i++
				c = line[i]
			}

			n, _ := strconv.Atoi(num)
			numbers = append(numbers, NumberOcurrence{n, startIndex, i - 1})
		}
	}

	return numbers
}

func isPartNumber(occ NumberOcurrence, prevLine, line, nextLine string) bool {
	var startIndex, endIndex int

	if occ.startIndex == 0 {
		startIndex = occ.startIndex
	} else {
		startIndex = occ.startIndex - 1
	}

	if occ.endIndex == len(line)-1 {
		endIndex = occ.endIndex
	} else {
		endIndex = occ.endIndex + 1
	}

	// check previous line
	if prevLine != "" {
		for i := startIndex; i <= endIndex; i++ {
			c := prevLine[i]

			if isSymbol(c) {
				return true
			}
		}
	}

	// check current line
	if isSymbol(line[startIndex]) || isSymbol(line[endIndex]) {
		return true
	}

	// check next line
	if nextLine != "" {
		for i := startIndex; i <= endIndex; i++ {
			c := nextLine[i]

			if isSymbol(c) {
				return true
			}
		}
	}

	return false
}
