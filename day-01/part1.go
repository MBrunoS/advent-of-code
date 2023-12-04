package main

import "bufio"

func part1(scanner *bufio.Scanner) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		num := extractNumber(line)
		sum += num
	}

	println("Part 1: ", sum)
}
