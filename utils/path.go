package utils

import "strings"

type Path []string

func NewPath(s string) Path {
	return Path{s}
}

func (p Path) Add(s string) Path {
	return CopyAndAppend(p, s)
}

func (p Path) Contains(s string) bool {
	return Contains(p, s)
}

func (p Path) First() string {
	return p[0]
}

func (p Path) Last() string {
	return p[len(p)-1]
}

func (p Path) Length() int {
	return len(p)
}

func (p Path) String() string {
	return strings.Join(p, " -> ")
}
