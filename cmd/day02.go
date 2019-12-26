package cmd

import (
	"fmt"

	"github.com/herczogzoltan/aoc2019/src/core"
	"github.com/herczogzoltan/aoc2019/src/day02"
)

const desiredOutput = 19690720

type Day02 struct{}

func init() {
	RegisterDay("day02", func() core.Day { return &Day02{} })
}

func (d Day02) Part1(input []byte) string {
	parsed := day02.Parser(input)
	// because 1202 program alarm: https://www.hq.nasa.gov/alsj/a11/a11.landing.html#1023832
	parsed[1] = 12
	parsed[2] = 02
	return fmt.Sprintf("%d", day02.IntCode(parsed))
}

func (d Day02) Part2(input []byte) string {
	noun := 0
	for noun <= 99 {
		verb := 0
		for verb <= 99 {
			parsed := day02.Parser(input)
			parsed[1] = noun
			parsed[2] = verb
			result := day02.IntCode(parsed)

			if result == desiredOutput {
				return fmt.Sprintf("%d", 100*noun+verb)
			}

			verb++
		}
		noun++
	}

	return fmt.Sprintf("%d", 0)
}
