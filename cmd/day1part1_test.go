package cmd

import "testing"

func TestCalculateFuel(t *testing.T) {
	result := calculateFuel(9)

	if result != 1 {
		t.Errorf("Calculation is incorrect")
	}

}
