package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day-02-input.txt")
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

func getIdAndSets(line string) (int, []string) {
	game := strings.Split(line, ":")[0]
	id_str := strings.Split(game, " ")[1]
	id, err := strconv.Atoi(id_str)

	if err != nil {
		log.Fatal(err)
	}

	set_str := strings.Split(line, ": ")[1]
	sets := strings.Split(set_str, "; ")

	return id, sets
}

func getSetAmounts(set string) (int, int, int) {
	colors := strings.Split(set, ", ")

	red := 0
	green := 0
	blue := 0

	for _, color := range colors {
		amount, err := strconv.Atoi(strings.Split(color, " ")[0])
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(color, "red") {
			red = amount
		} else if strings.Contains(color, "green") {
			green = amount
		} else if strings.Contains(color, "blue") {
			blue = amount
		}
	}
	return red, green, blue
}
