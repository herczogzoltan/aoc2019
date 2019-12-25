package cmd

import (
	"fmt"

	"github.com/herczogzoltan/aoc2019/src/core"
	"github.com/herczogzoltan/aoc2019/src/day03"
)

type Day03 struct{}

func init() {
	RegisterDay("day03", func() core.Day { return &Day03{} })
}

func (d Day03) Part1(input []byte) string {

	var wires []day03.Wire
	lines := day03.ParseInput(input)

	for _, line := range lines {
		wires = append(wires, day03.GetWirePath(line))
	}

	return fmt.Sprintf("%d", day03.FindClosest(day03.FindIntersectPoints(wires[0], wires[1])))
}

func (d Day03) Part2(input []byte) string {

	var wires []day03.Wire
	lines := day03.ParseInput(input)

	for _, line := range lines {
		wires = append(wires, day03.GetWirePath(line))
	}

	return fmt.Sprintf("%d", day03.FindFastest(day03.FindIntersectPoints(wires[0], wires[1])))
}
