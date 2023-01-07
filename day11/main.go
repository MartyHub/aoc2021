package main

import "aoc2021/utils"

func part1(input cave) {
	utils.Check(1, 1625, input.steps(100))
}

func part2(input cave) {
	result := 0

	for {
		input.step()
		result++

		if input.zeros() {
			break
		}
	}

	utils.Check(2, 244, result)
}

func main() {
	part1(parseCave("input.txt"))
	part2(parseCave("input.txt"))
}
