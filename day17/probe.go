package main

import (
	"fmt"
	"image"
)

type probe struct {
	position image.Point
	velocity image.Point
}

func newProbe(vx, vy int) probe {
	return probe{
		position: image.Point{X: 0, Y: 0},
		velocity: image.Point{X: vx, Y: vy},
	}
}

func (p probe) in(target image.Rectangle) bool {
	return target.Min.X <= p.position.X && p.position.X <= target.Max.X &&
		target.Min.Y <= p.position.Y && p.position.Y <= target.Max.Y
}

func (p probe) done(target image.Rectangle) bool {
	return p.position.Y < target.Min.Y && p.position.Y < target.Max.Y
}

func (p probe) step() probe {
	vx := p.velocity.X

	if vx > 0 {
		vx--
	} else if vx < 0 {
		vx++
	}

	return probe{
		position: p.position.Add(p.velocity),
		velocity: image.Point{
			X: vx,
			Y: p.velocity.Y - 1,
		},
	}
}

func (p probe) String() string {
	return fmt.Sprintf("probe{position: %v, velocity: %v}", p.position, p.velocity)
}
