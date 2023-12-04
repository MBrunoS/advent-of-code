package main

import (
	"bufio"
	"fmt"
)

const (
	reds_max_amount   = 12
	greens_max_amount = 13
	blues_max_amount  = 14
)

func part1(scanner *bufio.Scanner) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		id, sets := getIdAndSets(line)

		if isValidGame(sets) {
			sum += id
		}
	}

	fmt.Println("Part 1: ", sum)
}

func isValidGame(sets []string) bool {
	for _, set := range sets {
		if !isValidSet(set) {
			return false
		}
	}
	return true
}

func isValidSet(set string) bool {
	red, green, blue := getSetAmounts(set)

	if red > reds_max_amount || green > greens_max_amount || blue > blues_max_amount {
		return false
	}

	return true
}
