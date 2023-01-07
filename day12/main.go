package main

import (
	"aoc2021/utils"
	"strings"
)

func parse(fileName string) graph {
	result := newGraph()

	for lr := utils.NewLineReader(fileName); lr.HasNext(); {
		caves := strings.Split(lr.Text(), "-")

		result.add(result.cave(caves[0]).add(caves[1]))
		result.add(result.cave(caves[1]).add(caves[0]))
	}

	return result
}

func part1(input graph) {
	result := 0
	queue := []utils.Path{utils.NewPath("start")}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		key := path.String()

		if visited[key] {
			continue
		}

		visited[key] = true

		for _, link := range input.cave(path.Last()).links {
			if link == end {
				result++
			} else if input.cave(link).big() || !path.Contains(link) {
				queue = append(queue, path.Add(link))
			}
		}
	}

	utils.Check(1, 3576, result)
}

func part2(input graph) {
	result := 0
	queue := []utils.Path{utils.NewPath("start")}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		key := path.String()

		if visited[key] {
			continue
		}

		visited[key] = true

		for _, link := range input.cave(path.Last()).links {
			if link == end {
				result++
			} else if link != start && (input.cave(link).big() ||
				!path.Contains(link) ||
				input.onlySingleSmallCave(path)) {
				queue = append(queue, path.Add(link))
			}
		}
	}

	utils.Check(2, 84271, result)
}

func main() {
	input := parse("input.txt")

	part1(input)
	part2(input)
}
