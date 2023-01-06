package utils

import "image"

type Line struct {
	From, To image.Point
}

func (l Line) DeltaX() int {
	dx := l.To.X - l.From.X

	if dx < 0 {
		return -1
	} else if dx > 0 {
		return 1
	} else {
		return 0
	}
}

func (l Line) DeltaY() int {
	dy := l.To.Y - l.From.Y

	if dy < 0 {
		return -1
	} else if dy > 0 {
		return 1
	} else {
		return 0
	}
}

func (l Line) Horizontal() bool {
	return l.From.Y == l.To.Y
}

func (l Line) Vertical() bool {
	return l.From.X == l.To.X
}
