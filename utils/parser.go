package utils

import (
	"fmt"
	"image"
	"log"
)

type LineParser[T any] func(string) (T, error)

func ParseFile[T any](fileName string, lp LineParser[T]) []T {
	result := make([]T, 0)

	for lr := NewLineReader(fileName); lr.HasNext(); {
		if elem, err := lp(lr.Text()); err != nil {
			log.Fatalf("Failed to parse line %d (%s): %v", lr.Line, lr.Text(), err)
		} else {
			result = append(result, elem)
		}
	}

	if len(result) == 0 {
		log.Fatalf("No lines parsed in %s", fileName)
	}

	return result
}

func ParseIntegerLine(s string) ([]int, error) {
	result := make([]int, len(s))

	for i, c := range s {
		result[i] = int(c - '0')
	}

	return result, nil
}

func ParsePointLine(s string) (image.Point, error) {
	result := image.Point{}
	_, err := fmt.Sscanf(s, "%d,%d", &result.X, &result.Y)

	return result, err
}
