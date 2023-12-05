package main

import (
	"log"
	"os"
)

func main() {
	lines, err := readLinesFrom("day-03-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if (len(os.Args) > 1) && (os.Args[1] == "2") {
		part2(lines)
		return
	}

	part1(lines)
}
