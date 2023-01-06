package main

import (
	"aoc2021/utils"
	"fmt"
	"image"
	"log"
)

func parse(fileName string) []utils.Line {
	result := make([]utils.Line, 0)

	for lr := utils.NewLineReader(fileName); lr.HasNext(); {
		line := utils.Line{}

		if _, err := fmt.Sscanf(lr.Text(), "%d,%d -> %d,%d", &line.From.X, &line.From.Y, &line.To.X, &line.To.Y); err != nil {
			log.Fatalf("Failed to parse line %d (%s): %v", lr.Line, lr.Text(), err)
		}

		result = append(result, line)
	}

	return result
}

func part1(lines []utils.Line) {
	points := make(map[image.Point]int)

	for _, line := range lines {
		if line.Horizontal() {
			for x := line.From.X; x != line.To.X; x += line.DeltaX() {
				points[image.Point{X: x, Y: line.From.Y}]++
			}

			points[line.To]++
		} else if line.Vertical() {
			for y := line.From.Y; y != line.To.Y; y += line.DeltaY() {
				points[image.Point{X: line.From.X, Y: y}]++
			}

			points[line.To]++
		}
	}

	result := 0

	for _, v := range points {
		if v > 1 {
			result++
		}
	}

	log.Printf("Part 1: %d", result)
}

func part2(lines []utils.Line) {
	points := make(map[image.Point]int)

	for _, line := range lines {
		x := line.From.X
		y := line.From.Y

		for {
			points[image.Point{X: x, Y: y}]++

			x += line.DeltaX()
			y += line.DeltaY()

			if x == line.To.X && y == line.To.Y {
				break
			}
		}

		points[line.To]++
	}

	result := 0

	for _, v := range points {
		if v > 1 {
			result++
		}
	}

	log.Printf("Part 2: %d", result)
}

func main() {
	lines := parse("input.txt")

	part1(lines)
	part2(lines)
}
