package main

import "unicode"

const (
	start = "start"
	end   = "end"
)

type cave struct {
	name  string
	links []string
}

func (c cave) add(link string) cave {
	return cave{name: c.name, links: append(c.links, link)}
}

func (c cave) big() bool {
	for _, r := range c.name {
		if unicode.IsUpper(r) {
			return true
		}
	}

	return false
}
