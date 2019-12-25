package cmd

import (
	"fmt"

	"github.com/herczogzoltan/aoc2019/src/core"
	"github.com/herczogzoltan/aoc2019/src/day04"
)

type Day04 struct{}

func init() {
	RegisterDay("day04", func() core.Day { return &Day04{} })
}

func (d *Day04) Part1(input []byte) string {
	return fmt.Sprintf("%d", day04.FindPasswords(day04.ParseInput(input)))
}

func (d *Day04) Part2(input []byte) string {
	return fmt.Sprintf("Part 2 solution!")
}
