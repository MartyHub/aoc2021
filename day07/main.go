package main

import (
	"aoc2021/utils"
	"fmt"
)

func parse(fileName string) []int {
	for lr := utils.NewLineReader(fileName); lr.HasNext(); {
		return lr.Integers()
	}

	panic(fmt.Sprintf("No data in file %s", fileName))
}

func positions(data []int) map[int]int {
	result := make(map[int]int)

	for i := 0; i < utils.Max(data...); i++ {
		result[i] = 0
	}

	return result
}

func fuel1(data []int, position int) int {
	result := 0

	for _, i := range data {
		result += utils.Abs(i - position)
	}

	return result
}

func part1(data []int) {
	fuelByPosition := positions(data)

	for p := range fuelByPosition {
		fuelByPosition[p] = fuel1(data, p)
	}

	utils.Check(1, 341558, utils.MinValue(fuelByPosition))
}

func move(amount int) int {
	result := 0

	for ; amount > 0; amount-- {
		result += amount
	}

	return result
}

func fuel2(data []int, position int) int {
	result := 0

	for _, i := range data {
		result += move(utils.Abs(i - position))
	}

	return result
}

func part2(data []int) {
	fuelByPosition := positions(data)

	for p := range fuelByPosition {
		fuelByPosition[p] = fuel2(data, p)
	}

	utils.Check(2, 93214037, utils.MinValue(fuelByPosition))
}

func main() {
	data := parse("input.txt")

	part1(data)
	part2(data)
}
