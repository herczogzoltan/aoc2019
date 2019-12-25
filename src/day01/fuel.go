package day01

func CalculateFuel(mass int64) int64 {
	return ((mass / 3) - 2)
}

func CalcSecond(mass int64) int64 {
	inMassFuel := CalculateFuel(mass)
	fuelForFuel := CalculateFuel(inMassFuel)

	totalFuel := inMassFuel + fuelForFuel
	for {
		fuel := CalculateFuel(fuelForFuel)

		if fuel <= 0 {
			break
		}
		fuelForFuel = fuel
		totalFuel += fuel
	}

	return totalFuel
}
