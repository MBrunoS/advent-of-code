package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day-01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if (len(os.Args) > 1) && (os.Args[1] == "2") {
		part2(scanner)
		return
	}

	part1(scanner)
}

func extractNumber(line string) int {
	var first_digit, sec_digit string

	for _, c := range line {
		if c >= '0' && c <= '9' {
			if first_digit == "" {
				first_digit = string(c)
			} else {
				sec_digit = string(c)
			}
		}
	}

	if sec_digit == "" {
		sec_digit = first_digit
	}

	num, _ := strconv.Atoi(first_digit + sec_digit)
	return num
}
