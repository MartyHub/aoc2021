package main

import (
	"aoc2021/utils"
	"log"
	"strconv"
	"strings"
)

var bingos = make([]Bingo, 0)

func parseNumbers(s string) []int {
	values := strings.Split(s, ",")
	result := make([]int, 0)

	for _, v := range values {
		n, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, n)
	}

	return result
}

func parseGridLine(s string) [5]int {
	values := strings.Split(s, " ")
	result := [5]int{}
	i := 0

	for _, v := range values {
		if v != "" {
			n, err := strconv.Atoi(v)

			if err != nil {
				log.Fatal(err)
			}

			result[i] = n
			i += 1
		}
	}

	return result
}

func rowBingo(row [5]int) bool {
	for _, v := range row {
		if v == 0 {
			return false
		}
	}

	return true
}

func columnBingo(marks [][5]int, column int) bool {
	for i := 0; i < 5; i++ {
		if marks[i][column] == 0 {
			return false
		}
	}

	return true
}

func exists(grid, index int, row bool) bool {
	for _, b := range bingos {
		if grid == b.grid && index == b.index && row == b.row {
			return true
		}
	}

	return false
}

func gridBingo(grid int, marks [][5]int) bool {
	for i, row := range marks {
		if rowBingo(row) && !exists(grid, i, true) {
			bingos = append(bingos, Bingo{
				grid:  grid,
				index: i,
				row:   true,
			})

			return true
		}
	}

	for j := 0; j < 5; j++ {
		if columnBingo(marks, j) && !exists(grid, j, false) {
			bingos = append(bingos, Bingo{
				grid:  grid,
				index: j,
				row:   false,
			})

			return true
		}
	}

	return false
}

func bingo(marks [][5]int) int {
	for i := 0; i < len(marks); i += 5 {
		if gridBingo(i/5, marks[i:i+5]) {
			return i
		}
	}

	return -1
}

func score(grid, marks [][5]int, n int) int {
	result := 0

	for i, row := range marks {
		for j, v := range row {
			if v == 0 {
				result += grid[i][j]
			}
		}
	}

	return result * n
}

func main() {
	lr := utils.NewLineReader("input.txt")

	var numbers []int
	grids := make([][5]int, 0)

	for lr.HasNext() {
		text := lr.Text()

		if lr.Line == 1 {
			numbers = parseNumbers(text)
		} else if text != "" {
			grids = append(grids, parseGridLine(text))
		}
	}

	log.Printf("Numbers: %v", numbers)
	log.Printf("%v grid(s)", len(grids)/5)

	marks := make([][5]int, len(grids))
	bingoGrids := make(map[int]bool)

	for _, n := range numbers {
		for i, row := range grids {
			for j, v := range row {
				if n == v {
					marks[i][j] = 1
				}
			}
		}

		if grid := bingo(marks); grid != -1 {
			if !bingoGrids[grid] {
				bingoGrids[grid] = true
				log.Printf("First Bingo for grid %v with number %v, score = %v", grid, n, score(grids[grid:grid+5], marks[grid:grid+5], n))
			}
		}
	}
}
