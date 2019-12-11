package cmd

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day1)
}

var (
	day1 = &cobra.Command{
		Use: "day1",
		Run: func(cmd *cobra.Command, args []string) {
			switch Part {
			case 1:
				fmt.Println("First part result: ", firstPart())
			case 2:
				fmt.Println("Second part result: ", secondPart())
			}

		},
	}
)

func firstPart() int {
	file := LoadSourceFile(InputSource)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		Check(err)
		total += calculateFuel(mass)
	}

	return total
}

func calculateFuel(mass int) int {
	return int(math.Trunc(((float64(mass) / 3) - 2)))
}

func secondPart() int {
	file := LoadSourceFile(InputSource)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		moduleMass, err := strconv.Atoi(scanner.Text())
		Check(err)
		total += calcSecond(moduleMass)
	}

	return total
}

func calcSecond(mass int) int {
	inMassFuel := calculateFuel(mass)
	fuelForFuel := calculateFuel(inMassFuel)

	result := inMassFuel + fuelForFuel
	for {
		fuel := calculateFuel(fuelForFuel)

		if fuel <= 0 {
			break
		}
		fuelForFuel = fuel
		result += fuel
	}

	return result

}
