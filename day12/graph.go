package main

import "aoc2021/utils"

type graph struct {
	caves map[string]cave
}

func newGraph() graph {
	return graph{caves: make(map[string]cave)}
}

func (g graph) add(c cave) {
	g.caves[c.name] = c
}

func (g graph) cave(name string) cave {
	result, found := g.caves[name]

	if !found {
		result = cave{name: name}

		g.caves[name] = result
	}

	return result
}

func (g graph) onlySingleSmallCave(p utils.Path) bool {
	count := make(map[string]int)

	for _, name := range p {
		if !g.cave(name).big() && count[name] == 1 {
			return false
		}

		count[name]++
	}

	return true
}
