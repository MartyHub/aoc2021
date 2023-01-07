package utils

import "log"

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
