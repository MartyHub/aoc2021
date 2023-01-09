package main

import (
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	data string
}

func newParser(input string) *parser {
	sb := strings.Builder{}

	for _, c := range input {
		i, err := strconv.ParseUint(string(c), 16, 4)

		if err != nil {
			panic(err)
		}

		sb.WriteString(fmt.Sprintf("%04b", i))
	}

	return &parser{sb.String()}
}

func (p *parser) parse() packet {
	result := packet{}

	result.version = p.nextInt(3)
	result.typeId = p.nextInt(3)

	if result.typeId == 4 {
		result.value = p.nextLiteral()
	} else {
		if p.skip(1) == "0" {
			bitsLength := p.nextInt(15)
			result.children = make([]packet, 0)

			for l := len(p.data); ; {
				result.children = append(result.children, p.parse())

				if len(p.data) == l-bitsLength {
					break
				}
			}
		} else {
			packetsCount := p.nextInt(11)
			result.children = make([]packet, packetsCount)

			for i := 0; i < packetsCount; i++ {
				result.children[i] = p.parse()
			}
		}
	}

	return result
}

func (p *parser) skip(n int) string {
	result := p.data[:n]

	p.data = p.data[n:]

	return result
}

func (p *parser) parseInt(s string) int {
	result, err := strconv.ParseUint(s, 2, len(s))

	if err != nil {
		panic(err)
	}

	return int(result)
}

func (p *parser) nextInt(bitSize int) int {
	return p.parseInt(p.skip(bitSize))
}

func (p *parser) nextLiteral() int {
	sb := strings.Builder{}

	for p.skip(1) == "1" {
		sb.WriteString(p.skip(4))
	}

	sb.WriteString(p.skip(4))

	return p.parseInt(sb.String())
}
