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

		output := parts[1]
		tokens := strings.Split(output, " ")

		for _, token := range tokens {
			l := len(token)

			if l == 2 || l == 3 || l == 4 || l == 7 {
				result++
			}
		}
	}

	utils.Check(1, 375, result)
}

func main() {
	lines := utils.NewLineReader("input.txt").Lines()

	part1(lines)
}
