package utils

import (
	"strconv"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func PrettyFormat[N int | int64](i N) string {
	runes := []rune(strconv.FormatInt(int64(i), 10))
	start := len(runes) - 1
	result := ""

	for i := start; i >= 0; i-- {
		if i != start && (start-i)%3 == 0 {
			result = " " + result
		}

		result = string(runes[i]) + result
	}

	return result
}
