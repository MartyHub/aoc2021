package utils

import (
	"fmt"
	"image"
	"log"
	"math"
	"strings"
)

type Node[T comparable] struct {
	Id    T
	Links map[T]int
}

func NewNode[T comparable](id T) Node[T] {
	return Node[T]{
		Id:    id,
		Links: make(map[T]int),
	}
}

func (n Node[T]) min(ids []T) int {
	min := math.MaxInt
	result := -1

	for i, id := range ids {
		if value, found := n.Links[id]; found {
			if value < min {
				min = value
				result = i
			} else if value == math.MaxInt {
				break
			}
		}
	}

	if result == -1 {
		log.Fatalf("Failed to find min in %v", ids)
	}

	return result
}

func (n Node[T]) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Node %v:", n.Id))

	for k, v := range n.Links {
		sb.WriteString(fmt.Sprintf(" \n  - %v = %d", k, v))
	}

	return sb.String()
}

type Graph[T comparable] interface {
	Ids() []T
	Node(id T) Node[T]
}

type SimpleGraph[T comparable] struct {
	Nodes map[T]Node[T]
}

func NewSimpleGraph[T comparable]() SimpleGraph[T] {
	return SimpleGraph[T]{
		Nodes: make(map[T]Node[T]),
	}
}

func (g SimpleGraph[T]) Ids() []T {
	return Keys(g.Nodes)
}

func (g SimpleGraph[T]) Length() int {
	return len(g.Nodes)
}

func (g SimpleGraph[T]) Node(id T) Node[T] {
	result, found := g.Nodes[id]

	if !found {
		result = NewNode(id)

		g.Nodes[id] = result
	}

	return result
}

func (g SimpleGraph[T]) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Graph with %d node(s):", g.Length()))

	for _, n := range g.Nodes {
		sb.WriteRune('\n')
		sb.WriteString(n.String())
	}

	return sb.String()
}

type BoardGraph Board[int]

func (g BoardGraph) Ids() []image.Point {
	result := make([]image.Point, 0, Board[int](g).Surface())

	for y := range g {
		for x := range g[y] {
			result = append(result, image.Point{X: x, Y: y})
		}
	}

	return result
}

func (g BoardGraph) Node(id image.Point) Node[image.Point] {
	result := NewNode(id)

	for _, n := range ValidCardinalNeighbors(id.X, id.Y, Board[int](g).Width(), Board[int](g).Height()) {
		result.Links[n] = g[n.Y][n.X]
	}

	return result
}
