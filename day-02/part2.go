package main

import (
	"bufio"
	"fmt"
)

func part2(scanner *bufio.Scanner) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		_, sets := getIdAndSets(line)
		red, green, blue := getMinimumValidSet(sets)

		sum += red * green * blue
	}

	fmt.Println("Part 2: ", sum)
}

func getMinimumValidSet(sets []string) (int, int, int) {
	red_max, green_max, blue_max := 0, 0, 0

	for _, set := range sets {
		red, green, blue := getSetAmounts(set)

		if red > red_max {
			red_max = red
		}

		if green > green_max {
			green_max = green
		}

		if blue > blue_max {
			blue_max = blue
		}
	}

	return red_max, green_max, blue_max
}
