package main

import (
	"aoc2021/utils"
	"fmt"
)

type packet struct {
	version  int
	typeId   int
	children []packet
	value    int
}

func (p packet) operator() bool {
	return p.typeId != 4
}

func (p packet) sumVersion() int {
	result := p.version

	for _, c := range p.children {
		result += c.sumVersion()
	}

	return result
}

func (p packet) compute() int {
	if !p.operator() {
		return p.value
	}

	values := make([]int, len(p.children))

	for i, c := range p.children {
		values[i] = c.compute()
	}

	switch p.typeId {
	case 0:
		return utils.Sum(values)
	case 1:
		return utils.Mul(values)
	case 2:
		return utils.Min(values...)
	case 3:
		return utils.Max(values...)
	case 5:
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	}

	panic(fmt.Sprintf("Don't know how to handle type %d", p.typeId))
}

func (p packet) String() string {
	if p.operator() {
		return fmt.Sprintf("Operator %d v%d: %v", p.typeId, p.version, p.children)
	} else {
		return fmt.Sprintf("Literal v%d = %d", p.version, p.value)
	}
}
