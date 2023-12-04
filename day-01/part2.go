package main

import (
	"bufio"
	"strings"
)

func part2(scanner *bufio.Scanner) {
	sum := 0

	numberWords := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = replaceWordsWithNumbers(line, numberWords)
		num := extractNumber(line)
		sum += num
	}

	println("Part 2: ", sum)
}

func replaceWordsWithNumbers(line string, numberWords map[string]string) string {
	for word, number := range numberWords {
		line = strings.ReplaceAll(line, word, number)
	}
	return line
}
