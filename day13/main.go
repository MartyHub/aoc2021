package main

import (
	"aoc2021/utils"
	"fmt"
)

func part1(input paper, instructions []instruction) {
	paper := input.fold(instructions[0])

	utils.Check(1, 671, len(paper.points))
}

func part2(input paper, instructions []instruction) {
	for _, instruction := range instructions {
		input = input.fold(instruction)
	}

	fmt.Println(input.draw())

	utils.Check(2, "PCPHARKL", "PCPHARKL")
}

func main() {
	paper, instructions := parse("input.txt")

	part1(paper, instructions)
	part2(paper, instructions)
}
