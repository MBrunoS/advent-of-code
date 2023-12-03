package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func replaceWordsWithNumbers(line string, numberWords map[string]string) string {
	for word, number := range numberWords {
		line = strings.ReplaceAll(line, word, number)
	}
	return line
}

func extractNumber(line string) int {
	var digits [2]rune

	for _, c := range line {
		if c >= '0' && c <= '9' {
			if digits[0] == 0 { // first digit found
				digits[0] = c
			} else { // second+ digit found
				digits[1] = c
			}
		}
	}

	if digits[1] == 0 { // only one digit found
		digits[1] = digits[0]
	}

	return int(digits[0]-'0')*10 + int(digits[1]-'0') // convert runes to int
}
