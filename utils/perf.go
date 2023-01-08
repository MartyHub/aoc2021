package utils

import (
	"fmt"
	"time"
)

type stat struct {
	max, min, sum time.Duration
	count         int
}

func (s stat) avg() time.Duration {
	return s.sum / time.Duration(s.count)
}

func (s stat) String() string {
	return fmt.Sprintf("%d calls in %v (min = %v, avg = %v, max = %v)", s.count, s.sum, s.min, s.avg(), s.max)
}

var stats = make(map[string]*stat)

func add(name string, elapsed time.Duration) {
	s, found := stats[name]

	if !found {
		s = &stat{}
		stats[name] = s
	}

	s.count++
	s.sum += elapsed

	if elapsed > s.max {
		s.max = elapsed
	}

	if elapsed < s.min || s.min == 0 {
		s.min = elapsed
	}
}

func Record(name string) func() {
	start := time.Now()

	return func() {
		add(name, time.Since(start))
	}
}

func PrintStats() {
	for name, s := range stats {
		fmt.Printf("%s: %v\n", name, s)
	}
}
