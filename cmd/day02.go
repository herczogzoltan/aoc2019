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
	return fmt.Sprintf("%d", day02.IntCode(day02.Parser(input), 12, 02))
}

func (d Day02) Part2(input []byte) string {
	noun := 0
	for noun <= 99 {
		verb := 0
		for verb <= 99 {

			result := day02.IntCode(day02.Parser(input), noun, verb)

			if result == desiredOutput {
				return fmt.Sprintf("%d", 100*noun+verb)
			}

			verb++
		}
		noun++
	}

	return fmt.Sprintf("%d", 0)
}
