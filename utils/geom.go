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
