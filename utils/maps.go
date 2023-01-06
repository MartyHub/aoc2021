package utils

import "math"

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
