package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	day2 = &cobra.Command{
		Use: "day2",
		Run: func(cmd *cobra.Command, args []string) {
			switch Part {
			case 1:
				fmt.Println("First part result: ", gravityAssist())
				// case 2:
				// 	fmt.Println("Second part result: ", secondPart())
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(day2)
}

func gravityAssist() int {
	file := LoadSourceFile(InputSource)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	nums := strings.Split(firstLine, ",")

	var numbers []int
	for _, value := range nums {
		num, err := strconv.Atoi(value)
		Check(err)
		numbers = append(numbers, num)
	}

	i := 0
	// because 1202 program alarm: https://www.hq.nasa.gov/alsj/a11/a11.landing.html#1023832
	numbers[1] = 12
	numbers[2] = 2
	for i != 99 && i < len(numbers) {
		opCode := numbers[i]

		switch opCode {
		case 1:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
		case 2:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
		case 99:
			break
		}

		i += 4
	}

	return numbers[0]
}
