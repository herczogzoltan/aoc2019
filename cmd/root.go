package cmd

import (
	"os"

	"fmt"
	"io/ioutil"

	"github.com/herczogzoltan/aoc2019/src/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	inputSource string
	part        int

	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 ",
		Long: `Advent of Code 2019 solutions written in
	Go-lang as a practice and the joy of my own.`,
	}
)

type CommandFunc func() core.Day

// RegisterDay will be used to register a day from outside.
func RegisterDay(name string, f CommandFunc) {
	rootCmd.AddCommand(&cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {
			input, err := ioutil.ReadFile(inputSource)
			if err != nil {
				input = []byte(inputSource)
			}

			for input[len(input)-1] == 0x0a {
				input = input[:len(input)-1]
			}

			day := f()
			switch part {
			case 1:
				fmt.Println("First part result: ", day.Part1(input))
			case 2:
				fmt.Println("Second part result: ", day.Part2(input))
			}
		},
	})
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputSource, "source", "s", "", "File path which contains the input for the assignment")
	rootCmd.PersistentFlags().IntVarP(&part, "part", "p", 1, "Which part do you want to run for the specific day?")
	rootCmd.MarkFlagRequired("source")
	rootCmd.MarkFlagRequired("part")
}

// Execute the root command.
func Execute() error {
	return rootCmd.Execute()
}

func LoadSourceFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("An error occured when tried to open the file")
	}
	return file
}

// Check checks if error is not nil, and then logs it.
func Check(e error) {
	if e != nil {
		log.Error(e)
	}
}
