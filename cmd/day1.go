package cmd

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

type day1 Day

var (
	day1Command = &cobra.Command{
		Use: "day1",
		Run: func(cmd *cobra.Command, args []string) {
			day1 := day1{
				filePath: InputSource,
			}
			switch Part {
			case 1:
				fmt.Println("First part result: ", day1.firstPart())
			case 2:
				fmt.Println("Second part result: ", day1.secondPart())
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(day1Command)
}

func (d day1) firstPart() int {
	file := LoadSourceFile(d.filePath)
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

func (d day1) secondPart() int {
	file := LoadSourceFile(d.filePath)
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
