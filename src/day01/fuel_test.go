package day01

import "testing"

func TestCalculateFuel(t *testing.T) {
	c := []struct {
		i int64
		o int64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, v := range c {
		f := CalculateFuel(v.i)
		if f != v.o {
			t.Errorf("Calculation of fuel for %d was incorrect, got: %d want: %d\n", v.i, f, v.o)
		}
	}
}

func TestCalcSecond(t *testing.T) {
	c := []struct {
		i int64
		o int64
	}{
		{1969, 966},
		{100756, 50346},
	}

	for _, v := range c {
		f := CalcSecond(v.i)
		if f != v.o {
			t.Errorf("Calculation of second for %d was incorrect, got: %d want: %d\n", v.i, f, v.o)
		}
	}
}
