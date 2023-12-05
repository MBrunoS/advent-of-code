package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func part2(scanner *bufio.Scanner) {
	cards_copies := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		card_id, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(line))
		matching_nums := getMatchingNumbers(line)

		cards_copies[card_id]++

		for i := 1; i <= matching_nums; i++ {
			copy_id := card_id + i
			cards_copies[copy_id] += cards_copies[card_id]
		}
	}

	sum := sumCards(cards_copies)
	fmt.Println("Part 2: ", sum)
}

func sumCards(cards map[int]int) int {
	sum := 0

	for _, value := range cards {
		sum += value
	}

	return sum
}
