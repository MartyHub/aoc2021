package main

import (
	"aoc2021/utils"
	"log"
)

func main() {
	lr := utils.NewLineReader("input.txt")

	s := Submarine{}
	s2 := Submarine{}

	for lr.HasNext() {
		c := Parse(lr.Text())

		s = s.move(c)
		s2 = s2.move2(c)
	}

	log.Printf("Submarine: %+v => %v", s, utils.PrettyFormat(s.depth*s.position))     // 1 250 395
	log.Printf("Submarine2: %+v => %v", s2, utils.PrettyFormat(s2.depth*s2.position)) // 1 451 210 346
}
