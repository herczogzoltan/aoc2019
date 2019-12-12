package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type day3 Day

type wire struct {
	coordinates []coordinate
}

type coordinate struct {
	x int
	y int
}

var (
	day3Command = &cobra.Command{
		Use: "day3",
		Run: func(cmd *cobra.Command, args []string) {
			day3 := day3{
				filePath: InputSource,
			}
			switch Part {
			case 1:
				fmt.Printf("First part result: %v", day3.firstPart())
			case 2:
				//fmt.Printf("Second part result: %v", day3.secondPart())
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(day3Command)
}

func (d day3) firstPart() int {
	file := LoadSourceFile(InputSource)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstWire := getWirePath(scanner.Text())
	scanner.Scan()
	secondWire := getWirePath(scanner.Text())
	intersectPoints := findIntersectPoints(firstWire, secondWire)

	closest := 0
	for i, intersect := range intersectPoints {
		newValue := calculateManhattanDistance(intersect)
		if i == 0 {
			closest = newValue
			continue
		}

		if newValue < closest {
			closest = newValue
		}
	}

	return closest
}

func getWirePath(line string) wire {
	pathString := strings.Split(line, ",")
	var wire wire
	x := 0
	y := 0
	for _, route := range pathString {
		direction := route[0]
		distance, err := strconv.Atoi(route[1:len(route)])
		Check(err)

		switch direction {
		case 'U':
			for i := 0; i < distance; i++ {
				y++
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		case 'D':
			for i := 0; i < distance; i++ {
				y--
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		case 'R':
			for i := 0; i < distance; i++ {
				x++
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		case 'L':
			for i := 0; i < distance; i++ {
				x--
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		}
	}

	return wire
}

func findIntersectPoints(a wire, b wire) []coordinate {
	var coordinates []coordinate

	for _, Acoords := range a.coordinates {
		for _, Bcoords := range b.coordinates {

			if Acoords.x == Bcoords.x {
				if Acoords.y == Bcoords.y {
					coordinates = append(coordinates, Bcoords)
				}
			}
		}
	}

	return coordinates
}

func calculateManhattanDistance(coordinate coordinate) int {
	return absoluteValue(coordinate.x) + absoluteValue(coordinate.y)
}

func absoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d day3) secondPart() {

}
