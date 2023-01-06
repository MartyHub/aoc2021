package main

import (
	"aoc2021/utils"
	"log"
	"strconv"
	"strings"
)

func toBinary(ones []int, mid int, oneRune, zeroRune rune) string {
	sb := strings.Builder{}

	for _, one := range ones {
		if one > mid {
			sb.WriteRune(oneRune)
		} else {
			sb.WriteRune(zeroRune)
		}
	}

	return sb.String()
}

func toDecimal(s string) int64 {
	result, err := strconv.ParseInt(s, 2, 64)

	if err != nil {
		log.Fatalf("Failed to convert %v to decimal", s)
	}

	return result
}

func countOne(values [][]rune, index int) int {
	result := 0

	for _, value := range values {
		if value[index] == '1' {
			result += 1
		}
	}

	return result
}

func compare(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func keep(oneCount, zeroCount int, keepMax bool) rune {
	result := '1'

	switch compare(oneCount, zeroCount) {
	case -1: // oneCount < zeroCount
		if keepMax {
			result = '0'
		}
	case 1: // oneCount > zeroCount
		if !keepMax {
			result = '0'
		}
	default: // oneCount == zeroCount
		if !keepMax {
			result = '0'
		}
	}

	return result
}

func filter(values [][]rune, index int, keepMax bool) [][]rune {
	oneCount := countOne(values, index)
	zeroCount := len(values) - oneCount
	keep := keep(oneCount, zeroCount, keepMax)
	result := make([][]rune, 0)

	for _, value := range values {
		if value[index] == keep {
			result = append(result, value)
		}
	}

	return result
}

func part2(values [][]rune) {
	co2Values := values
	oxygenValues := values

	for i := 0; i < len(values[0]); i++ {
		if len(co2Values) > 1 {
			co2Values = filter(co2Values, i, false)
		}

		if len(oxygenValues) > 1 {
			oxygenValues = filter(oxygenValues, i, true)
		}

		if len(co2Values) == 1 && len(oxygenValues) == 1 {
			break
		}
	}

	if len(co2Values) != 1 || len(oxygenValues) != 1 {
		log.Fatalf("Invalid CO2 or oxygen: %v, %v", co2Values, oxygenValues)
	}

	co2 := toDecimal(string(co2Values[0]))
	oxygen := toDecimal(string(oxygenValues[0]))

	log.Printf("CO2: %v", co2)
	log.Printf("Oxygen: %v", oxygen)

	log.Printf("Life Support Rating: %v", utils.PrettyFormat(co2*oxygen))
}

func main() {
	lr := utils.NewLineReader("input.txt")

	var ones []int
	values := make([][]rune, 0)

	for lr.HasNext() {
		text := lr.Text()
		values = append(values, []rune(text))

		if len(ones) == 0 {
			ones = make([]int, len(text))
		}

		for i, c := range text {
			if c == '1' {
				ones[i] += 1
			}
		}
	}

	mid := lr.Line / 2

	bEpsilonRate := toBinary(ones, mid, '0', '1')
	bGammaRate := toBinary(ones, mid, '1', '0')

	epsilonRate := toDecimal(bEpsilonRate)
	gammaRate := toDecimal(bGammaRate)

	log.Printf("Epsilon Rate: %v = %v", bEpsilonRate, utils.PrettyFormat(epsilonRate)) // 2 795
	log.Printf("Gamma Rate:   %v = %v", bGammaRate, utils.PrettyFormat(gammaRate))     // 1 300
	log.Printf("Power Consuption: %v", utils.PrettyFormat(epsilonRate*gammaRate))      // 3 633 500

	part2(values)
}
