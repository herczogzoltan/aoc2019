package cmd

import (
	"fmt"

	"github.com/herczogzoltan/aoc2019/src/core"
	"github.com/herczogzoltan/aoc2019/src/day02"
)

type Day05 struct{}

func init() {
	RegisterDay("day05", func() core.Day { return &Day05{} })
}

func (d *Day05) Part1(input []byte) string {
	return fmt.Sprintf("%d", day02.IntCode(day02.Parser(input)))
}

func (d *Day05) Part2(input []byte) string {
	return fmt.Sprintf("%d", day02.IntCode(day02.Parser(input)))
}
