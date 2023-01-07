package main

import (
	"aoc2021/utils"
	"image"
	"math"
)

type cave [][]int

func parseCave(fileName string) cave {
	return utils.ParseFile(fileName, utils.ParseIntegerLine)
}

func (c cave) width() int {
	return len(c[0])
}

func (c cave) height() int {
	return len(c)
}

func (c cave) value(x, y int) int {
	if x < 0 || y < 0 || x >= c.width() || y >= c.height() {
		return math.MinInt
	}

	return c[y][x]
}

func (c cave) inc() {
	for y := 0; y < c.height(); y++ {
		for x := 0; x < c.width(); x++ {
			c[y][x]++
		}
	}
}

func (c cave) clean() {
	for y := 0; y < c.height(); y++ {
		for x := 0; x < c.width(); x++ {
			if c[y][x] > 9 {
				c[y][x] = 0
			}
		}
	}
}

func (c cave) steps(iterations int) int {
	result := 0

	for i := 0; i < iterations; i++ {
		result += c.step()
	}

	return result
}

func (c cave) step() int {
	c.inc()

	queue := c.flash()
	result := len(queue)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, n := range utils.ValidNeighbors(p.X, p.Y, c.width(), c.height()) {
			c[n.Y][n.X]++

			if c[n.Y][n.X] == 10 {
				queue = append(queue, n)
				result++
			}
		}
	}

	c.clean()

	return result
}

func (c cave) flash() []image.Point {
	result := make([]image.Point, 0)

	for y := 0; y < c.height(); y++ {
		for x := 0; x < c.width(); x++ {
			if c.value(x, y) == 10 {
				result = append(result, image.Pt(x, y))
			}
		}
	}

	return result
}

func (c cave) zeros() bool {
	for y := 0; y < c.height(); y++ {
		for x := 0; x < c.width(); x++ {
			if c[y][x] != 0 {
				return false
			}
		}
	}

	return true
}
