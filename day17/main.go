package main

import (
	"aoc2021/utils"
	"fmt"
	"image"
	"log"
)

func parse(fileName string) image.Rectangle {
	lr := utils.NewLineReader(fileName)

	if !lr.HasNext() {
		log.Fatalf("Empty file %s", fileName)
	}

	result := image.Rectangle{}

	if _, err := fmt.Sscanf(lr.Text(), "target area: x=%d..%d, y=%d..%d",
		&result.Min.X,
		&result.Max.X,
		&result.Min.Y,
		&result.Max.Y); err != nil {
		log.Fatalf("Failed to parse %s: %v", lr.Text(), err)
	}

	return result.Canon()
}

func part1(input image.Rectangle) {
	result := 0

	for vx := 1; vx <= 1_000; vx++ {
		for vy := 1; vy <= 1_000; vy++ {
			in := false
			maxY := 0

			for p := newProbe(vx, vy); !p.done(input); p = p.step() {
				in = in || p.in(input)

				if p.position.Y > maxY {
					maxY = p.position.Y
				}
			}

			if in && maxY > result {
				result = maxY
			}
		}
	}

	utils.Check(1, 25200, result)
}

func part2(input image.Rectangle) {
	result := 0

	for vx := 1; vx <= 1_000; vx++ {
		for vy := -1_000; vy <= 1_000; vy++ {
			for p := newProbe(vx, vy); !p.done(input); p = p.step() {
				if p.in(input) {
					result++

					break
				}
			}
		}
	}

	utils.Check(2, 3012, result)
}

func main() {
	input := parse("input.txt")

	log.Printf("Input: %v", input)

	part1(input)
	part2(input)
}
