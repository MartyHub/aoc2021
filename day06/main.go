package main

import (
	"aoc2021/utils"
	"log"
)

func parse(fileName string) map[int]int {
	result := make(map[int]int)

	for lr := utils.NewLineReader(fileName); lr.HasNext(); {
		for _, i := range lr.Integers() {
			result[i]++
		}
	}

	return result
}

func simulate(data map[int]int, days int) int {
	currentData := data

	for i := 0; i < days; i++ {
		newData := make(map[int]int)

		for k, v := range currentData {
			if k > 0 {
				newData[k-1] += v
			} else {
				newData[6] += v
				newData[8] += v
			}
		}

		currentData = newData
	}

	result := 0

	for _, count := range currentData {
		result += count
	}

	return result
}
func part1(data map[int]int) {
	utils.Check(1, 394994, simulate(data, 80))
}

func part2(data map[int]int) {
	utils.Check(2, 1765974267455, simulate(data, 256))
}

func main() {
	data := parse("input.txt")

	log.Printf("Data: %v", data)

	part1(data)
	part2(data)
}
