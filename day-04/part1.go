package main

import (
	"bufio"
	"fmt"

	"github.com/mbrunos/advent-of-code/day-04/utils"
)

func part1(scanner *bufio.Scanner) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		matching_nums := getMatchingNumbers(line)

		var card_value int

		if matching_nums == 0 {
			card_value = 0
		} else {
			card_value = utils.IntPow(2, matching_nums-1)
		}

		sum += card_value
	}

	fmt.Println("Part 1: ", sum)
}
