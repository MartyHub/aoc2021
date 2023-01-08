package utils

import (
	"fmt"
	"strings"
)

type Path[S comparable] struct {
	Steps []S
	Costs []int
}

func NewPath[S comparable](step S) Path[S] {
	return Path[S]{
		Steps: []S{step},
		Costs: []int{0},
	}
}

func (p Path[S]) Add(step S, cost int) Path[S] {
	return Path[S]{
		Steps: CopyAndAppend(p.Steps, step),
		Costs: CopyAndAppend(p.Costs, cost),
	}
}

func (p Path[S]) Contains(step S) bool {
	return Contains(p.Steps, step)
}

func (p Path[S]) Cost() int {
	return Sum(p.Costs)
}

func (p Path[S]) First() S {
	return p.Steps[0]
}

func (p Path[S]) Last() S {
	return p.Steps[len(p.Steps)-1]
}

func (p Path[S]) Length() int {
	return len(p.Steps)
}

func (p Path[S]) String() string {
	sb := strings.Builder{}
	cost := 0

	for i, step := range p.Steps {
		if i > 0 {
			sb.WriteString(" -> ")
		}

		sb.WriteString(fmt.Sprintf("%v (%d)", step, p.Costs[i]))
		cost += p.Costs[i]
	}

	sb.WriteString(fmt.Sprintf(" = %d", cost))

	return sb.String()
}
