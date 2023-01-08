package utils

import (
	"math"
)

func Dijkstra[T comparable](g Graph[T], from T) Node[T] {
	queue := g.Ids()
	result := dijkstraInit(from, queue)

	for len(queue) > 0 {
		i := result.min(queue)
		id := queue[i]
		queue = Remove(queue, i)

		for link, value := range g.Node(id).Links {
			cost := result.Links[id] + value

			if cost < result.Links[link] {
				result.Links[link] = cost

				if index, found := FindFirst(queue, link); found && index != 0 {
					queue = Prepend(Remove(queue, index), link)
				}
			}
		}
	}

	return result
}

func dijkstraInit[T comparable](from T, to []T) Node[T] {
	result := NewNode(from)

	for _, id := range to {
		result.Links[id] = math.MaxInt
	}

	result.Links[from] = 0

	return result
}

//func (d dijkstra[T]) LongestPath(from T) Path[T] {
//	queue := []Path[T]{NewPath(from)}
//	l := d.Graph.Length()
//	var result Path[T]
//
//	for len(queue) > 0 {
//		path := queue[0]
//		queue = queue[1:]
//
//		if path.Length() == l {
//			if result.Length() == 0 || path.Cost() > result.Cost() {
//				result = path
//			}
//		} else {
//			for link, value := range d.Graph.Node(path.Last()).Links {
//				if !path.Contains(link) {
//					queue = append(queue, path.Add(link, value))
//				}
//			}
//		}
//	}
//
//	return result
//}

//func (d dijkstra[T]) ShortestPath(from T) Path[T] {
//	queue := []Path[T]{NewPath(from)}
//	l := d.Graph.Length()
//	var result Path[T]
//
//	for len(queue) > 0 {
//		path := queue[0]
//		queue = queue[1:]
//
//		if path.Length() == l {
//			if result.Length() == 0 || path.Cost() < result.Cost() {
//				result = path
//			}
//		} else {
//			for link, value := range d.Graph.Node(path.Last()).Links {
//				if !path.Contains(link) {
//					queue = append(queue, path.Add(link, value))
//				}
//			}
//		}
//	}
//
//	return result
//}
