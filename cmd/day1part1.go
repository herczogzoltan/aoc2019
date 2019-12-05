package cmd

import (
	"bufio"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day1part1Cmd)
}

var (
	day1part1Cmd = &cobra.Command{
		Use: "day1part1",
		Run: func(cmd *cobra.Command, args []string) {
			day1part1()
		},
	}
)

func day1part1() {
	file := LoadSourceFile(InputSource)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Error("An error occured while converting the line of the file from string to int")
		}
		total += calculateFuel(mass)
	}

	fmt.Println("Total: ", total)
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}
