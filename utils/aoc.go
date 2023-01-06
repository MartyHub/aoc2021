package utils

import "log"

func Check[T comparable](part int, expected, result T) {
	if result == expected {
		log.Printf("[Part %d] result = %v", part, result)
	} else {
		log.Printf("[Part %d] expected %v, got %v", part, expected, result)
	}
}
