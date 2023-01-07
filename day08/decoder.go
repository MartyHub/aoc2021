package main

import (
	"aoc2021/utils"
	"log"
	"math/bits"
	"strconv"
	"strings"
)

var lengthMappings = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	// 5: 2, 3, 5
	// 6: 0, 6, 9
	7: 8,
}

type decoder struct {
	mappings map[digit]int
	unknowns []digit
}

func newDecoder() *decoder {
	return &decoder{
		mappings: make(map[digit]int),
	}
}

func (d *decoder) addAll(digits []digit) int {
	result := 0

	for _, digit := range digits {
		if d.add(digit) {
			result++
		} else {
			d.unknowns = append(d.unknowns, digit)
		}
	}

	return result
}

func (d *decoder) add(n digit) bool {
	if mapping, lengthMappingFound := lengthMappings[n.length]; lengthMappingFound {
		d.addMapping(n, mapping)

		return true
	}

	return false
}

func (d *decoder) addMapping(n digit, value int) {
	mapping, found := d.mappings[n]

	if found {
		if mapping != value {
			log.Fatalf("Duplicate mapping for %v", n)
		}
	} else {
		d.mappings[n] = value
	}
}

func (d *decoder) get(n int) (digit, bool) {
	for digit, value := range d.mappings {
		if value == n {
			return digit, true
		}
	}

	return digit{}, false
}

func (d *decoder) compute() {
	d.compute9()
	d.compute5()
	d.compute6()
	d.compute0()
	d.compute3()
	d.compute2()
}

func (d *decoder) compute0() bool {
	d6, f6 := d.get(6)
	d8, f8 := d.get(8)
	d9, f9 := d.get(9)

	if f6 && f8 && f9 {
		for _, digit := range d.unknowns {
			if d6.code != digit.code && d9.code != digit.code && bits.OnesCount8(digit.code^d8.code) == 1 {
				d.addMapping(digit, 0)
			}
		}

		return true
	}

	return false
}

func (d *decoder) compute2() bool {
	for _, digit := range d.unknowns {
		if _, found := d.mappings[digit]; !found {
			d.addMapping(digit, 2)
		}
	}

	return true
}

func (d *decoder) compute3() bool {
	d5, f5 := d.get(5)
	d9, f9 := d.get(9)

	if f5 && f9 {
		for _, digit := range d.unknowns {
			if d5.code != digit.code && bits.OnesCount8(digit.code^d9.code) == 1 {
				d.addMapping(digit, 3)
			}
		}

		return true
	}

	return false
}

func (d *decoder) compute5() bool {
	d1, f1 := d.get(1)
	d9, f9 := d.get(9)

	if f1 && f9 {
		code := d1.code ^ d9.code

		for _, digit := range d.unknowns {
			if bits.OnesCount8(digit.code^code) == 1 {
				d.addMapping(digit, 5)
			}
		}

		return true
	}

	return false
}

func (d *decoder) compute6() bool {
	d5, f5 := d.get(5)
	d9, f9 := d.get(9)

	if f5 && f9 {
		for _, digit := range d.unknowns {
			if d9.code != digit.code && bits.OnesCount8(digit.code^d5.code) == 1 {
				d.addMapping(digit, 6)
			}
		}

		return true
	}

	return false
}

func (d *decoder) compute9() bool {
	d4, f4 := d.get(4)
	d7, f7 := d.get(7)

	if f4 && f7 {
		code := d4.code | d7.code

		for _, digit := range d.unknowns {
			if bits.OnesCount8(digit.code^code) == 1 {
				d.addMapping(digit, 9)
			}
		}

		return true
	}

	return false
}

func (d *decoder) decode(s string) int {
	sb := strings.Builder{}

	for _, digit := range parseDigits(s) {
		sb.WriteRune(rune(d.mappings[digit] + '0'))
	}

	return utils.Must(strconv.Atoi(sb.String()))
}
