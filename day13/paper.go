package main

import (
	"aoc2021/utils"
	"fmt"
	"image"
	"log"
	"strings"
)

type instruction struct {
	axis  rune
	value int
}

func parseInstruction(line string) instruction {
	result := instruction{}

	if _, err := fmt.Sscanf(line, "fold along %c=%d", &result.axis, &result.value); err != nil {
		log.Fatalf("Failed to parse instruction %s: %v", line, err)
	}

	return result
}

type paper struct {
	points map[image.Point]bool
}

func newPaper() paper {
	return paper{points: make(map[image.Point]bool)}
}

func parse(fileName string) (paper, []instruction) {
	lr := utils.NewLineReader(fileName)
	paper := newPaper()

	for lr.HasNext() {
		line := lr.Text()

		if line == "" {
			break
		}

		if point, err := utils.ParsePointLine(line); err != nil {
			log.Fatalf("Failed to parse point %s: %v", line, err)
		} else {
			paper.points[point] = true
		}
	}

	instructions := make([]instruction, 0)

	for lr.HasNext() {
		instructions = append(instructions, parseInstruction(lr.Text()))
	}

	return paper, instructions
}

func (p paper) fold(instruction instruction) paper {
	if instruction.axis == 'x' {
		return p.foldAlongX(instruction.value)
	} else {
		return p.foldAlongY(instruction.value)
	}
}

func (p paper) foldAlongX(x int) paper {
	result := newPaper()

	for point := range p.points {
		if point.X >= x {
			point.X = 2*x - point.X
		}

		result.points[point] = true
	}

	return result
}

func (p paper) foldAlongY(y int) paper {
	result := newPaper()

	for point := range p.points {
		if point.Y >= y {
			point.Y = 2*y - point.Y
		}

		result.points[point] = true
	}

	return result
}

func (p paper) maxX() int {
	result := 0

	for point := range p.points {
		if point.X > result {
			result = point.X
		}
	}

	return result
}

func (p paper) maxY() int {
	result := 0

	for point := range p.points {
		if point.Y > result {
			result = point.Y
		}
	}

	return result
}

func (p paper) draw() string {
	sb := strings.Builder{}

	for y := 0; y <= p.maxY(); y++ {
		if y > 0 {
			sb.WriteRune('\n')
		}

		for x := 0; x <= p.maxX(); x++ {
			if p.points[image.Point{X: x, Y: y}] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune(' ')
			}
		}
	}

	return sb.String()
}
