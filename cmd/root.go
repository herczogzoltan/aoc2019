package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// InputSource is the filepath to read input values from.
	InputSource string

	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 ",
		Long: `Advent of Code 2019 solutions written in
	Go-lang as a practice and the joy of my own.`,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&InputSource, "source", "src", "", "File path which contains the input for the assignment")
	rootCmd.MarkFlagRequired("source")

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
