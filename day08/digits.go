package main

import (
	"strings"
)

func toByte(s string) uint8 {
	result := uint8(0)

	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'

		result += 1 << c
	}

	return result
}

type digit struct {
	code   uint8
	length int
}

func newDigit(s string) digit {
	return digit{
		code:   toByte(s),
		length: len(s),
	}
}

func parseDigits(s string) []digit {
	tokens := strings.Split(s, " ")
	result := make([]digit, len(tokens))

	for i, token := range tokens {
		result[i] = newDigit(token)
	}

	return result
}

func (d digit) String() string {
	sb := strings.Builder{}

	for i := 0; i < 7; i++ {
		if d.code&(1<<i) != 0 {
			sb.WriteByte('a' + byte(i))
		}
	}

	return sb.String()
}
