package main

import (
	"aoc2021/utils"
)

func part1(p polymer, rules map[string]rune) {
	for i := 0; i < 10; i++ {
		p = p.apply(rules)
	}

	utils.Check(1, 2975, p.score())
}

func part2(p polymer, rules map[string]rune) {
	for i := 0; i < 40; i++ {
		p = p.apply(rules)
	}

	utils.Check(2, 3015383850689, p.score())
}

func main() {
	polymer, rules := parseManual("input.txt")

	part1(polymer, rules)
	part2(polymer, rules)
}
