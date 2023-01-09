package main

import (
	"aoc2021/utils"
	"log"
)

func parse(fileName string) *parser {
	return newParser(utils.NewLineReader(fileName).Lines()[0])
}

func part1(p *parser) {
	result := p.parse().sumVersion()

	utils.Check(1, 971, result)
}

func part2(p *parser) {
	utils.Check(2, 831996589851, p.parse().compute())
}

func main() {
	log.Println(newParser("D2FE28").parse())
	log.Println(newParser("38006F45291200").parse())
	log.Println(newParser("EE00D40C823060").parse())

	utils.Check(0, 16, newParser("8A004A801A8002F478").parse().sumVersion())
	utils.Check(0, 12, newParser("620080001611562C8802118E34").parse().sumVersion())
	utils.Check(0, 23, newParser("C0015000016115A2E0802F182340").parse().sumVersion())
	utils.Check(0, 31, newParser("A0016C880162017C3686B18A3D4780").parse().sumVersion())

	part1(parse("input.txt"))

	utils.Check(0, 3, newParser("C200B40A82").parse().compute())
	utils.Check(0, 54, newParser("04005AC33890").parse().compute())
	utils.Check(0, 7, newParser("880086C3E88112").parse().compute())
	utils.Check(0, 9, newParser("CE00C43D881120").parse().compute())
	utils.Check(0, 1, newParser("D8005AC2A8F0").parse().compute())
	utils.Check(0, 0, newParser("F600BC2D8F").parse().compute())
	utils.Check(0, 0, newParser("9C005AC2F8F0").parse().compute())
	utils.Check(0, 1, newParser("9C0141080250320F1802104A08").parse().compute())

	part2(parse("input.txt"))
}
