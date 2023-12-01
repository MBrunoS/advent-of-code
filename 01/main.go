package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		var digits [2]rune
		line := scanner.Text()

		for word, digit := range numberWords {
			line = strings.ReplaceAll(line, word, digit)
		}

		for _, c := range line {
			if c >= '0' && c <= '9' {
				if digits[0] == 0 {
					digits[0] = c
				} else {
					digits[1] = c
				}
			}
		}

		if digits[1] == 0 {
			digits[1] = digits[0]
		}

		num := int(digits[0]-'0')*10 + int(digits[1]-'0')
		sum += num
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
