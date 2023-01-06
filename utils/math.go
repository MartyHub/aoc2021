package utils

func Abs[N int | int64](i N) N {
	if i < 0 {
		return -i
	}

	return i
}

func Max[N int | int64](values ...N) N {
	l := len(values)

	if l == 0 {
		panic("No values")
	}

	result := values[0]

	for i := 1; i < l; i++ {
		if values[i] > result {
			result = values[i]
		}
	}

	return result
}

func Min[N int | int64](values ...N) N {
	l := len(values)

	if l == 0 {
		panic("No values")
	}

	result := values[0]

	for i := 1; i < l; i++ {
		if values[i] < result {
			result = values[i]
		}
	}

	return result
}
