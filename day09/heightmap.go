package main

import (
	"aoc2021/utils"
	"image"
	"math"
)

type heightmap [][]int

func parseHeightmap(fileName string) heightmap {
	return utils.ParseFile(fileName, utils.ParseIntegerLine)
}

func (h heightmap) maxX() int {
	return len(h[0]) - 1
}

func (h heightmap) maxY() int {
	return len(h) - 1
}

func (h heightmap) height(x, y int) int {
	if x < 0 || y < 0 || x > h.maxX() || y > h.maxY() {
		return math.MaxInt
	}

	return h[y][x]
}

func (h heightmap) lowPoint(x, y int) bool {
	height := h.height(x, y)

	for _, p := range utils.CardinalNeighbors(x, y) {
		if h.height(p.X, p.Y) <= height {
			return false
		}
	}

	return true
}

func (h heightmap) lowPoints() []image.Point {
	result := make([]image.Point, 0)

	for y := 0; y <= h.maxY(); y++ {
		for x := 0; x <= h.maxX(); x++ {
			if h.lowPoint(x, y) {
				result = append(result, image.Pt(x, y))
			}
		}
	}

	return result
}

func (h heightmap) riskLevel() int {
	result := 0

	for _, p := range h.lowPoints() {
		result += 1 + h.height(p.X, p.Y)
	}

	return result
}

func (h heightmap) basinSize(lowPoint image.Point) int {
	queue := []image.Point{lowPoint}
	visited := map[image.Point]bool{}

	for len(queue) > 0 {
		p := queue[0]

		queue = queue[1:]

		if visited[p] {
			continue
		}

		visited[p] = true

		ph := h.height(p.X, p.Y)

		for _, n := range utils.CardinalNeighbors(p.X, p.Y) {
			nh := h.height(n.X, n.Y)

			if nh != math.MaxInt && nh != 9 && nh > ph {
				queue = append(queue, n)
			}
		}
	}

	return len(visited)
}
