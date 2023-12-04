package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day-03-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if (len(os.Args) > 1) && (os.Args[1] == "1") {
		part1(scanner)
		return
	}

	if (len(os.Args) > 1) && (os.Args[1] == "2") {
		part2(scanner)
		return
	}

	part1(scanner)
}
