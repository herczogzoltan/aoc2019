package day03

import "testing"

func TestAbsoluteValue(t *testing.T) {
	c := []struct {
		i int
		o int
	}{
		{1, 1},
		{-2, 2},
		{33583, 33583},
		{-100756, 100756},
	}

	for _, v := range c {
		f := AbsoluteValue(v.i)
		if f != v.o {
			t.Errorf("Converting to absoule value for %d was incorrect, got: %d want: %d\n", v.i, f, v.o)
		}
	}
}
