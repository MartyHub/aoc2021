package utils

import "image"

func CardinalNeighbors(x, y int) []image.Point {
	return []image.Point{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
}

func ValidNeighbors(x, y int, width, height int) []image.Point {
	result := make([]image.Point, 0, 8)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			if x+dx < 0 || y+dy < 0 || x+dx >= width || y+dy >= height {
				continue
			}

			result = append(result, image.Point{X: x + dx, Y: y + dy})
		}
	}

	return result
}
