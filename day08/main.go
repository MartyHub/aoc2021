package main

import (
	"aoc2021/utils"
	"log"
	"strings"
)

func part1(lines []string) {
	result := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")

		if len(parts) != 2 {
			log.Fatalf("Invalid line: %s", line)
		}

		result += newDecoder().addAll(parseDigits(parts[1]))
	}

	utils.Check(1, 375, result)
}

func part2(lines []string) {
	result := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")

		if len(parts) != 2 {
			log.Fatalf("Invalid line: %s", line)
		}

		d := newDecoder()

		d.addAll(parseDigits(parts[0]))
		d.addAll(parseDigits(parts[1]))
		d.compute()

		result += d.decode(parts[1])
	}

	utils.Check(2, 1019355, result)
}

func main() {
	lines := utils.NewLineReader("input.txt").Lines()

	part1(lines)
	part2(lines)
}
