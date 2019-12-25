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
	coordinates map[int]coordinate
}

type coordinate struct {
	x int
	y int
}

func CalculateManhattanDistance(coordinate coordinate) int {
	return AbsoluteValue(coordinate.x) + AbsoluteValue(coordinate.y)
}

func FindClosest(wire Wire) int {
	closest := 0
	for _, intersect := range wire.coordinates {
		newValue := CalculateManhattanDistance(intersect)
		if closest == 0 {
			closest = newValue
			continue
		}

		if newValue < closest {
			closest = newValue
		}
	}

	return closest
}

func FindFastest(wire Wire) int {
	fastest := 0
	for i, _ := range wire.coordinates {
		newValue := i
		if fastest == 0 {
			fastest = newValue
			continue
		}

		if newValue < fastest {
			fastest = newValue
		}
	}

	return fastest
}

func FindIntersectPoints(a Wire, b Wire) Wire {
	var coordinates Wire
	coordinates.coordinates = make(map[int]coordinate)
	for i, Acoords := range a.coordinates {
		for j, Bcoords := range b.coordinates {

			if Acoords.x == Bcoords.x {
				if Acoords.y == Bcoords.y {
					coordinates.coordinates[i+j] = Bcoords
				}
			}
		}
	}

	return coordinates
}

func GetWirePath(line string) Wire {
	pathString := strings.Split(line, ",")
	var wire Wire
	wire.coordinates = make(map[int]coordinate)

	x := 0
	y := 0
	j := 0
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
				j++
				wire.coordinates[j] = coordinate{x: x, y: y}
			}
		case DirectionDown:
			for i := 0; i < distance; i++ {
				y--
				j++
				wire.coordinates[j] = coordinate{x: x, y: y}
			}
		case DirectionRight:
			for i := 0; i < distance; i++ {
				x++
				j++
				wire.coordinates[j] = coordinate{x: x, y: y}
			}
		case DirectionLeft:
			for i := 0; i < distance; i++ {
				x--
				j++
				wire.coordinates[j] = coordinate{x: x, y: y}
			}
		}

	}

	return wire
}
