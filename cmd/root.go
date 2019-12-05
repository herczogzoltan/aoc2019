package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// InputSource is the filepath to read input values from.
	InputSource string
	// Part of the daily assignment. Either 1 or 2.
	Part int

	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 ",
		Long: `Advent of Code 2019 solutions written in
	Go-lang as a practice and the joy of my own.`,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&InputSource, "source", "s", "", "File path which contains the input for the assignment")
	rootCmd.PersistentFlags().IntVarP(&Part, "part", "p", 1, "Which part do you want to run for the specific day?")
	rootCmd.MarkFlagRequired("source")
	rootCmd.MarkFlagRequired("part")

}

// Execute the root command.
func Execute() error {
	return rootCmd.Execute()
}

// LoadSourceFile loads a file to process by another command.
func LoadSourceFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("An error occured when tried to open the file")
	}
	return file
}
