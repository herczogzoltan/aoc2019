package cmd

import (
	"fmt"

	"github.com/herczogzoltan/aoc2019/src/core"
	"github.com/herczogzoltan/aoc2019/src/day01"
)

type Day01 struct{}

func init() {
	RegisterDay("day01", func() core.Day { return &Day01{} })
}

func (d *Day01) Part1(input []byte) string {
	var sum int64
	for _, v := range day01.ParseInput(input) {
		sum += day01.CalculateFuel(v)
	}
	return fmt.Sprintf("%d", sum)
}

func (d *Day01) Part2(input []byte) string {
	var totalMass int64
	for _, v := range day01.ParseInput(input) {
		totalMass += day01.CalcSecond(v)
	}

	return fmt.Sprintf("%d", totalMass)
}
