package main

import (
	"aoc2021/utils"
	"image"
)

func parse(fileName string) utils.Board[int] {
	return utils.ParseFile(fileName, utils.ParseIntegerLine)
}

func part1(input utils.Board[int]) {
	g := utils.BoardGraph(input)
	start := image.Point{}
	shortestPaths := utils.Dijkstra(utils.Graph[image.Point](g), start)
	end := image.Point{X: input.Width() - 1, Y: input.Height() - 1}

	utils.Check(1, 741, shortestPaths.Links[end])
}

func expand(board [][]int, mul int) utils.Board[int] {
	result := make([][]int, len(board)*mul)
	height := len(board)
	width := len(board[0])

	for y := range board {
		result[y] = make([]int, width*mul)

		for x := range board[y] {
			result[y][x] = board[y][x]

			for i := 1; i < mul; i++ {
				v := board[y][x] + i

				if v > 9 {
					v -= 9
				}

				result[y][x+width*i] = v
			}
		}
	}

	for y := 0; y < height; y++ {
		for i := 1; i < mul; i++ {
			result[y+height*i] = make([]int, width*mul)

			for x := range result[y] {
				v := result[y][x] + i

				if v > 9 {
					v -= 9
				}

				result[y+height*i][x] = v
			}
		}
	}

	return result
}

func part2(input utils.Board[int]) {
	b := expand(input, 5)
	g := utils.BoardGraph(b)
	start := image.Point{}
	shortestPaths := utils.Dijkstra(utils.Graph[image.Point](g), start)
	end := image.Point{X: b.Width() - 1, Y: b.Height() - 1}

	utils.Check(2, 2976, shortestPaths.Links[end])
}

func main() {
	input := parse("input.txt")

	part1(input)
	part2(input)

	utils.PrintStats()
}
