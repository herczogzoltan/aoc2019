package cmd

import (
	"bufio"
	"fmt"
	"os"
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
				fmt.Println("First part result: ", gravityAssist(12, 2))
			case 2:
				fmt.Printf("Second part result: %v", instructionPointer())
			}
		},
	}

	desiredOutput = 19690720
)

func init() {
	rootCmd.AddCommand(day2)
}

func gravityAssist(noun int, verb int) int {
	file := LoadSourceFile(InputSource)
	defer file.Close()
	addresses := getRowAsNumberArray(file)

	i := 0
	// because 1202 program alarm: https://www.hq.nasa.gov/alsj/a11/a11.landing.html#1023832
	addresses[1] = noun
	addresses[2] = verb
	for i < len(addresses) {
		opCode := addresses[i]

		switch opCode {
		case 1:
			addresses[addresses[i+3]] = addresses[addresses[i+1]] + addresses[addresses[i+2]]
		case 2:
			addresses[addresses[i+3]] = addresses[addresses[i+1]] * addresses[addresses[i+2]]
		case 99:
			break
		}

		i += 4
	}

	return addresses[0]
}

func instructionPointer() int {
	noun := 0
	for noun <= 99 {
		verb := 0
		for verb <= 99 {
			result := gravityAssist(noun, verb)

			fmt.Printf("Verb: %v, Noun: %v Result: %v\n", verb, noun, result)
			if result == desiredOutput {
				return 100*noun + verb
			}

			verb++
		}
		noun++
	}

	return 0
}

func getRowAsNumberArray(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	nums := strings.Split(firstLine, ",")

	var addresses []int
	for _, value := range nums {
		num, err := strconv.Atoi(value)
		Check(err)
		addresses = append(addresses, num)
	}
	return addresses
}
