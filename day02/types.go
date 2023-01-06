package main

import (
	"log"
	"strconv"
	"strings"
)

type Action string

const (
	Forward Action = "forward"
	Down    Action = "down"
	Up      Action = "up"
)

type Command struct {
	action Action
	amount int
}

func Parse(s string) Command {
	substrings := strings.Split(s, " ")

	if len(substrings) != 2 {
		log.Fatalf("Invalid command: %v", s)
	}

	amount, err := strconv.Atoi(substrings[1])

	if err != nil {
		log.Fatalf("Invalid amount for command: %v", s)
	}

	return Command{
		action: Action(substrings[0]),
		amount: amount,
	}
}

type Submarine struct {
	aim      int
	depth    int
	position int
}

func (s Submarine) move(c Command) Submarine {
	switch c.action {
	case Forward:
		s.position += c.amount
	case Down:
		s.depth += c.amount
	case Up:
		s.depth -= c.amount
	}

	return s
}

func (s Submarine) move2(c Command) Submarine {
	switch c.action {
	case Forward:
		s.position += c.amount
		s.depth += s.aim * c.amount
	case Down:
		s.aim += c.amount
	case Up:
		s.aim -= c.amount
	}

	return s
}
