package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("day-04-input.txt")
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

func getMatchingNumbers(line string) int {
	numbers := strings.Split(line, ": ")[1]
	win_nums_str := strings.Split(numbers, " | ")[0]
	have_nums_str := strings.Split(numbers, " | ")[1]

	win_nums := regexp.MustCompile(`\d+`).FindAllString(win_nums_str, -1)
	have_nums := regexp.MustCompile(`\d+`).FindAllString(have_nums_str, -1)
	matching_nums := 0

	for _, num := range win_nums {
		if slices.Contains(have_nums, num) {
			matching_nums++
		}
	}

	return matching_nums
}
