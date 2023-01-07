package utils

import "sort"

func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}

func CopyAndAppend[T any](s []T, v T) []T {
	l := len(s)
	result := make([]T, l+1)

	copy(result, s)

	result[l] = v

	return result
}

func Mul(s []int) int {
	result := 1

	for _, v := range s {
		result *= v
	}

	return result
}

func TopN(s []int, n int) []int {
	l := len(s)

	if l <= n {
		return s
	}

	sort.Ints(s)

	return s[l-n:]
}
