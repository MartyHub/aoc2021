package main

import (
	"aoc2021/utils"
	"log"
	"strconv"
)

func main() {
	lr := utils.NewLineReader("input.txt")

	increases := 0
	measures := make([]int, 0)
	previousDepth := 0

	for lr.HasNext() {
		depth, err := strconv.Atoi(lr.Text())

		if err != nil {
			log.Fatalf("Failed to parse line %v: %v", lr.Line, err)
		}

		measures = append(measures, depth)

		if lr.Line > 1 && depth > previousDepth {
			increases += 1
		}

		previousDepth = depth
	}

	log.Printf("Depth increases: %v", utils.PrettyFormat(increases)) // 1 696

	slide := 0

	for i := 3; i < len(measures); i++ {
		previous := measures[i-3] + measures[i-2] + measures[i-1]
		current := measures[i] + measures[i-1] + measures[i-2]

		if current > previous {
			slide += 1
		}
	}

	log.Printf("Slide increases: %v", utils.PrettyFormat(slide)) // 1 737
}
