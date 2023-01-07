package main

import (
	"aoc2021/utils"
)

func part1(input heightmap) {
	utils.Check(1, 478, input.riskLevel())
}

func part2(input heightmap) {
	lowPoints := input.lowPoints()
	basinSizes := make([]int, len(lowPoints))

	for i, p := range lowPoints {
		basinSizes[i] = input.basinSize(p)
	}

	utils.Check(2, 1327014, utils.Mul(utils.TopN(basinSizes, 3)))
}

func main() {
	input := parseHeightmap("input.txt")

	part1(input)
	part2(input)
}
