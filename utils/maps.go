package utils

import "math"

func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, len(m))
	i := 0

	for k := range m {
		result[i] = k
		i++
	}

	return result
}

func MaxValue[T comparable](m map[T]int) int {
	result := math.MinInt

	for _, v := range m {
		if v > result {
			result = v
		}
	}

	return result
}

func MinValue[T comparable](m map[T]int) int {
	result := math.MaxInt

	for _, v := range m {
		if v < result {
			result = v
		}
	}

	return result
}
