package day03

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	DirectionUp    string = "U"
	DirectionDown  string = "D"
	DirectionLeft  string = "L"
	DirectionRight string = "R"
)

type Wire struct {
	coordinates []coordinate
}

type coordinate struct {
	x int
	y int
}

func CalculateManhattanDistance(coordinate coordinate) int {
	return AbsoluteValue(coordinate.x) + AbsoluteValue(coordinate.y)
}

func FindClosest(coordinates []coordinate) int {
	closest := 0
	for i, intersect := range coordinates {
		newValue := CalculateManhattanDistance(intersect)
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

func FindIntersectPoints(a Wire, b Wire) []coordinate {
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

func GetWirePath(line string) Wire {
	pathString := strings.Split(line, ",")
	var wire Wire
	x := 0
	y := 0
	for _, route := range pathString {
		direction := string(route[0])
		distance, err := strconv.Atoi(route[1:len(route)])
		if err != nil {
			log.Error(err)
		}

		switch direction {
		case DirectionUp:
			for i := 0; i < distance; i++ {
				y++
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		case DirectionDown:
			for i := 0; i < distance; i++ {
				y--
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		case DirectionRight:
			for i := 0; i < distance; i++ {
				x++
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		case DirectionLeft:
			for i := 0; i < distance; i++ {
				x--
				wire.coordinates = append(wire.coordinates, coordinate{x: x, y: y})
			}
		}
	}

	return wire
}
