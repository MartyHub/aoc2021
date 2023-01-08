package main

import (
	"aoc2021/utils"
	"fmt"
)

func parseRule(s string) (string, rune) {
	var (
		source string
		target rune
	)

	if _, err := fmt.Sscanf(s, "%s -> %c", &source, &target); err != nil {
		panic(err)
	}

	return source, target
}

type polymer struct {
	elements map[string]int
	first    rune
}

func newPolymer(first rune) polymer {
	return polymer{
		elements: make(map[string]int),
		first:    first,
	}
}

func parsePolymer(s string) polymer {
	result := newPolymer(rune(s[0]))

	for i := 0; i < len(s)-1; i++ {
		result.elements[s[i:i+2]]++
	}

	return result
}

func (p polymer) apply(rules map[string]rune) polymer {
	result := newPolymer(p.first)

	for e, n := range p.elements {
		if target, ok := rules[e]; ok {
			result.elements[string([]rune{rune(e[0]), target})] += n
			result.elements[string([]rune{target, rune(e[1])})] += n
		} else {
			result.elements[e] += n
		}
	}

	return result
}

func (p polymer) count() map[rune]int {
	result := make(map[rune]int)

	result[p.first]++

	for e, n := range p.elements {
		result[rune(e[1])] += n
	}

	return result
}

func (p polymer) score() int {
	counts := p.count()

	return utils.MaxValue(counts) - utils.MinValue(counts)
}

func parseManual(fileName string) (polymer, map[string]rune) {
	lr := utils.NewLineReader(fileName)

	if !lr.HasNext() {
		panic("no polymer")
	}

	polymer := parsePolymer(lr.Text())
	rules := make(map[string]rune)

	for lr.HasNext() {
		line := lr.Text()

		if line != "" {
			source, target := parseRule(line)

			rules[source] = target
		}
	}

	return polymer, rules
}
