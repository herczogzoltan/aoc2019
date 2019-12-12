package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type day3Mock struct {
	mock.Mock
}

func (d *day3Mock) firstPart() int {
	firstWire := getWirePath("R8,U5,L5,D3")
	secondWire := getWirePath("U7,R6,D4,L4")
	return findClosest(findIntersectPoints(firstWire, secondWire))
}

func TestDay3FirstPart(t *testing.T) {
	day3 := new(day3Mock)
	assert.Equal(t, 6, day3.firstPart(), "they should be equal")
}
