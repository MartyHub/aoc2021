package main

import (
	"aoc2021/utils"
	"sort"
)

func parse(fileName string) []string {
	return utils.NewLineReader(fileName).Lines()
}

func part1(input []string) {
	result := 0

	for _, line := range input {
		score1, _ := check(line)

		result += score1
	}

	utils.Check(1, 345441, result)
}

func part2(input []string) {
	scores := make([]int, 0)

	for _, line := range input {
		if score1, score2 := check(line); score1 == 0 {
			scores = append(scores, score2)
		}
	}

	sort.Ints(scores)

	utils.Check(2, 3235371166, scores[(len(scores)-1)/2])
}

func main() {
	input := parse("input.txt")

	part1(input)
	part2(input)
}
